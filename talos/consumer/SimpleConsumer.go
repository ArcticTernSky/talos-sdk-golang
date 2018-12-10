/**
 * Copyright 2018, Xiaomi.
 * All rights reserved.
 * Author: wangfan8@xiaomi.com
 */

package consumer

import (
	"fmt"

	"github.com/XiaoMi/talos-sdk-golang/talos/client/compression"
	"github.com/XiaoMi/talos-sdk-golang/talos/thrift/common"
	"github.com/XiaoMi/talos-sdk-golang/talos/thrift/message"
	"github.com/XiaoMi/talos-sdk-golang/talos/thrift/topic"
	"github.com/XiaoMi/talos-sdk-golang/talos/utils"
	"github.com/XiaoMi/talos-sdk-golang/thrift"
	"github.com/alecthomas/log4go"
)

type SimpleConsumer struct {
	topicAndPartition *topic.TopicAndPartition
	messageClient     message.MessageService
	consumerConfig    *TalosConsumerConfig
	requestId         *int64
	simpleConsumerId  string
	//scheduleInfoCache client.ScheduleInfo
}

func NewSimpleConsumer(consumerConfig *TalosConsumerConfig,
	topicAndPartition *topic.TopicAndPartition,
	messageClient message.MessageService,
	consumerIdPrefix string) *SimpleConsumer {

	if err := utils.CheckTopicAndPartition(topicAndPartition); err != nil {
		log4go.Error("topicAndPartition error: %s, initial simpleConsumer failed",
			err.Error())
		return nil
	}
	consumerId, e := utils.CheckAndGenerateClientId(consumerIdPrefix)
	if e != nil {
		log4go.Error("generate clientId failed: %s, initial simpleConsumer failed",
			e.Error())
		return nil
	}
	//infoCache := client.NewScheduleInfoCache(topicAndPartition.TopicTalosResourceName,
	//	consumerConfig.TalosClientConfig, messageClient, talosClientFactory)
	return &SimpleConsumer{
		consumerConfig:    consumerConfig,
		topicAndPartition: topicAndPartition,
		messageClient:     messageClient,
		simpleConsumerId:  consumerId,
		requestId:         thrift.Int64Ptr(1),
		//scheduleInfoCache: infoCache,
	}
}

func NewSimpleConsumerForTest(consumerConfig *TalosConsumerConfig,
	topicAndPartition *topic.TopicAndPartition,
	messageClient message.MessageService) *SimpleConsumer {

	if err := utils.CheckTopicAndPartition(topicAndPartition); err != nil {
		log4go.Error("topicAndPartition error: %s, initial simpleConsumer failed",
			err.Error())
		return nil
	}
	clientId, _ := utils.CheckAndGenerateClientId("")
	return &SimpleConsumer{
		consumerConfig:    consumerConfig,
		topicAndPartition: topicAndPartition,
		messageClient:     messageClient,
		simpleConsumerId:  clientId,
		//scheduleInfoCache: scheduleInfoCache,
	}
}

func (c *SimpleConsumer) TopicAndPartition() *topic.TopicAndPartition {
	return c.topicAndPartition
}

func (c *SimpleConsumer) TopicTalosResourceName() *topic.TopicTalosResourceName {
	return c.topicAndPartition.GetTopicTalosResourceName()
}

func (c *SimpleConsumer) SetSimpleConsumerId(id string) {
	c.simpleConsumerId = id
}

func (c *SimpleConsumer) SimpleConsumerId() string {
	return c.simpleConsumerId
}

func (c *SimpleConsumer) PartitionId() int32 {
	return c.topicAndPartition.GetPartitionId()
}

func (c *SimpleConsumer) FetchMessage(startOffset int64,
	maxFetchedNumber int64) ([]*message.MessageAndOffset, error) {

	if err := utils.CheckStartOffsetValidity(startOffset); err != nil {
		return nil, err
	}
	if err := utils.CheckParameterRange(GALAXY_TALOS_CONSUMER_MAX_FETCH_RECORDS,
		maxFetchedNumber, GALAXY_TALOS_CONSUMER_MAX_FETCH_RECORDS_MINIMUM,
		GALAXY_TALOS_CONSUMER_MAX_FETCH_RECORDS_MAXIMUM); err != nil {
		return nil, err
	}
	requestSequenceId, err := utils.GenerateRequestSequenceId(c.simpleConsumerId,
		c.requestId)
	if err != nil {
		return nil, err
	}
	timestamp := utils.CurrentTimeMills() + c.consumerConfig.ClientTimeout()

	// limit the default max fetch bytes 2M
	getMessageRequest := &message.GetMessageRequest{
		TopicAndPartition:   c.topicAndPartition,
		MessageOffset:       startOffset,
		SequenceId:          requestSequenceId,
		MaxGetMessageNumber: int32(maxFetchedNumber),
		MaxGetMessageBytes:  GALAXY_TALOS_CONSUMER_MAX_FETCH_BYTES_DEFAULT,
		TimeoutTimestamp:    &timestamp,
	}

	getMessageResponse, e := c.messageClient.GetMessage(getMessageRequest)
	if e != nil {
		if getMessageRequest.GetMessageOffset() != -1 {
			getTopicOffsetRequest := &message.GetTopicOffsetRequest{
				TopicTalosResourceName: c.TopicTalosResourceName(),
			}
			getTopicOffsetResponse, e2 := c.messageClient.GetTopicOffset(getTopicOffsetRequest)
			if e2 == nil && getTopicOffsetResponse.GetOffsetInfoList()[c.PartitionId()].GetStartOffset() != startOffset {
				errCode := common.ErrorCode_MESSAGE_OFFSET_OUT_OF_RANGE
				log4go.Warn("getMessage error: %v, %s", errCode, e.Error())
				return nil, utils.NewTalosRuntimeError(errCode, e)
			}
			log4go.Warn("getTopicOffset error: %s", e2.Error())
			return nil, utils.NewTalosRuntimeError(common.ErrorCode_UNEXPECTED_ERROR, e2)
		} else {
			errCode := common.ErrorCode_UNEXPECTED_ERROR
			log4go.Warn("getMessage error: %v, %s", errCode, e.Error())
			return nil, utils.NewTalosRuntimeError(errCode, e)
		}
	}

	messageAndOffsetList := make([]*message.MessageAndOffset, 0)
	messageAndOffsetList, e = compression.Decompress(
		getMessageResponse.GetMessageBlocks(),
		getMessageResponse.GetUnHandledMessageNumber())
	if e != nil {
		errCode := common.ErrorCode_UNEXPECTED_ERROR
		log4go.Error("decompress messageBlock error: %s", e.Error())
		return nil, utils.NewTalosRuntimeError(errCode, e)
	}
	if len(messageAndOffsetList) <= 0 {
		return messageAndOffsetList, nil
	}

	actualStartOffset := messageAndOffsetList[0].GetMessageOffset()
	if messageAndOffsetList[0].GetMessageOffset() == startOffset ||
		startOffset == int64(message.MessageOffset_START_OFFSET) ||
		startOffset == int64(message.MessageOffset_LATEST_OFFSET) {
		return messageAndOffsetList, nil
	} else {
		start := startOffset - actualStartOffset
		if start <= 0 {
			errCode := common.ErrorCode_UNEXPECTED_ERROR
			err := fmt.Errorf("Unexpected subList start index: %d ", start)
			return nil, utils.NewTalosRuntimeError(errCode, err)
		}
		return messageAndOffsetList[start:], nil
	}
}

//func (c *SimpleConsumer) Shutdown() {
//	//onec called, all request of this topic in the process cannot auto location
//	c.scheduleInfoCache.Shutdown(c.TopicTalosResourceName())
//}
