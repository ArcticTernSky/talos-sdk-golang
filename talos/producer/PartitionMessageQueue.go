/**
 * Copyright 2018, Xiaomi.
 * All rights reserved.
 * Author: wangfan8@xiaomi.com
 */

package producer

import (
	"sync"
	"time"

	"github.com/XiaoMi/talos-sdk-golang/talos/thrift/message"
	"github.com/XiaoMi/talos-sdk-golang/talos/utils"
	"github.com/alecthomas/log4go"
)

type PartitionMessageQueue struct {
	userMessageList []*UserMessage
	curMessageBytes int64
	partitionId     int32
	talosProducer   Producer
	maxBufferedTime int64
	maxPutMsgNumber int64
	maxPutMsgBytes  int64
	MqWg            sync.WaitGroup
}

// Constructor method for test
func NewPartitionMessageQueueForTest(producerConfig *TalosProducerConfig,
	partitionId int32, talosProducerMock Producer) *PartitionMessageQueue {
	return &PartitionMessageQueue{
		userMessageList: make([]*UserMessage, 0),
		curMessageBytes: 0,
		partitionId:     partitionId,
		talosProducer:   talosProducerMock,
		maxBufferedTime: producerConfig.GetMaxBufferedMsgTime(),
		maxPutMsgNumber: producerConfig.GetMaxPutMsgNumber(),
		maxPutMsgBytes:  producerConfig.GetMaxPutMsgBytes(),
		MqWg:            sync.WaitGroup{},
	}
}

func NewPartitionMessageQueue(producerConfig *TalosProducerConfig,
	partitionId int32, talosProducer *TalosProducer) *PartitionMessageQueue {

	return &PartitionMessageQueue{
		userMessageList: make([]*UserMessage, 0),
		curMessageBytes: 0,
		partitionId:     partitionId,
		talosProducer:   talosProducer,
		maxBufferedTime: producerConfig.GetMaxBufferedMsgTime(),
		maxPutMsgNumber: producerConfig.GetMaxPutMsgNumber(),
		maxPutMsgBytes:  producerConfig.GetMaxPutMsgBytes(),
	}
}

func (q *PartitionMessageQueue) AddMessage(messageList []*UserMessage) {
	// notify partitionSender to getUserMessageList
	defer q.MqWg.Done()

	incrementBytes := int64(0)
	for _, userMessage := range messageList {
		q.userMessageList = append(q.userMessageList, userMessage)
		incrementBytes += userMessage.GetMessageSize()
	}
	q.curMessageBytes += int64(incrementBytes)
	// update total buffered count when add messageList
	q.talosProducer.IncreaseBufferedCount(int64(len(messageList)),
		int64(incrementBytes))
}

/**
 * return messageList, if not shouldPut, block in this method
 */
func (q *PartitionMessageQueue) GetMessageList() []*message.Message {
	for !q.shouldPut() {
		if waitTime := q.getWaitTime(); waitTime == 0 {
			// q.Mqwg.Add(1)
			q.MqWg.Wait()
		} else {
			time.Sleep(time.Duration(waitTime) * time.Millisecond)
		}
	}
	log4go.Info("getUserMessageList wake up for partition: %d", q.partitionId)

	returnList := make([]*message.Message, 0)
	returnMsgBytes, returnMsgNumber := int64(0), int64(0)

	for len(q.userMessageList) > 0 &&
		returnMsgNumber < q.maxPutMsgNumber && returnMsgBytes < q.maxPutMsgBytes {
		// userMessageList pollLast
		userMessage := q.userMessageList[0]
		q.userMessageList = q.userMessageList[1:]
		returnList = append(returnList, userMessage.GetMessage())
		q.curMessageBytes -= userMessage.GetMessageSize()
		returnMsgBytes += userMessage.GetMessageSize()
		returnMsgNumber++
	}

	// update total buffered count when poll messageList
	q.talosProducer.DecreaseBufferedCount(returnMsgNumber, returnMsgBytes)
	log4go.Info("Ready to put message batch: %d, queue size: %d and curBytes: %d"+
		" for partition: %d", len(returnList), len(q.userMessageList),
		q.curMessageBytes, q.partitionId)

	return returnList
}

func (q *PartitionMessageQueue) shouldPut() bool {
	// when TalosProducer is not active;
	if !q.talosProducer.IsActive() {
		return true
	}

	// when we have enough bytes data or enough number data;
	if q.curMessageBytes >= q.maxPutMsgBytes ||
		int64(len(q.userMessageList)) >= q.maxPutMsgNumber {
		return true
	}

	// when there have at least one message and it has exist enough long time;
	if len(q.userMessageList) > 0 && (utils.CurrentTimeMills()-
		q.userMessageList[0].GetTimestamp()) >= q.maxBufferedTime {
		return true
	}
	return false
}

/**
 * Note: 1. wait(0) represents wait infinite until be notified
 * 2. wait minimal 1 milli secs when time <= 0
 */
func (q *PartitionMessageQueue) getWaitTime() int64 {
	if len(q.userMessageList) <= 0 {
		return 0
	}
	time := q.userMessageList[0].GetTimestamp() + q.maxBufferedTime - utils.CurrentTimeMills()
	if time > 0 {
		return time
	} else {
		return 1
	}
}
