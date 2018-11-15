/**
 * Copyright 2018, Xiaomi.
 * All rights reserved.
 * Author: wangfan8@xiaomi.com
 */

package utils

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync/atomic"
	"time"
  "errors"

	"github.com/XiaoMi/talos-sdk-golang/talos/thrift/common"
	"github.com/XiaoMi/talos-sdk-golang/talos/thrift/message"
	"github.com/XiaoMi/talos-sdk-golang/talos/thrift/topic"
	"github.com/XiaoMi/talos-sdk-golang/thrift"
	"github.com/XiaoMi/talos-sdk-golang/talos/client"
	log "github.com/alecthomas/log4go"
	"github.com/nu7hatch/gouuid"
)

/**
 * read properties for client / consumer / producer
 */

type Properties struct {
	props map[string]string
}

func NewProperties() *Properties {
	properties := Properties{props: make(map[string]string)}
	return &properties
}

func LoadProperties(filename string) *Properties {

	file, err := os.Open(filename)
	if err != nil {
		log.Warn("Open file faild, ", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	maps := make(map[string]string)
	for scanner.Scan() {
		reg := regexp.MustCompile("\\s*#.*")
		line := scanner.Text()
		if !reg.MatchString(line) {
			reg = regexp.MustCompile("\\s*\\w+\\s*=.*")
			if reg.MatchString(line) {
				arr := strings.SplitN(line, "=", 2)
				key := strings.Trim(arr[0], " ")
				value := strings.Trim(arr[1], " ")
				maps[key] = value
			}
		}
	}
	return &Properties{props: maps}
}

func (p *Properties) Get(key string) string {
	return p.props[key]
}

func (p *Properties) GetProperty(key string, defaultValue string) string {
	value, ok := p.props[key]
	if !ok {
		return defaultValue
	}
	return value
}

func (p *Properties) SetProperty(key, value string) *Properties {
	p.props[key] = value
	return p
}

/**
 * The format of valid resource name is: ownerId#topicName#UUID
 * Note the 'ownerId' may contains the symbol '#',
 * so return topicName parsing from the tail of resourceName.
 */

func GetTopicNameByResourceName(topicTalosResourceName string) (string, error) {
	itemList := strings.Split(topicTalosResourceName, "#")
	len := len(itemList)
	if len < 3 {
		err := fmt.Errorf("%s is not a valid topicTalosResourceName",
			topicTalosResourceName)
		return "", err
	}
	return itemList[len-2], nil
}

func CheckArgument(expr bool) error {
	if !expr {
		err := errors.New("Illegal Argument! ")
		return err
	}
	return nil
}

func CheckParameterRange(parameter string, value int64, min int64, max int64) *client.TalosRuntimeError {
	if value < min || value > max {
		errCode := common.ErrorCode_INVALID_TOPIC_PARAMS
		err := fmt.Errorf("%s should be in range [%d,%d], got: %d ",
			parameter, min, max, value)
		return client.NewTalosRuntimeError(errCode, err)
	}
	return nil
}

func CheckTopicAndPartition(topicAndPartition *topic.TopicAndPartition) error {
	if strings.Contains(topicAndPartition.GetTopicName(),
		common.TALOS_CLOUD_TOPIC_NAME_DELIMITER) {
		err := errors.New(
			"The topic name format in TopicAndPartition should not be: orgId/topicName. ")
		return err
	}
	return nil
}

func CheckNameValidity(str string) error {
	if str == "" || len(str) <= 0 {
		return nil
	}

	matchBody, _ := regexp.MatchString(common.TALOS_NAME_BODY_REGEX, str)
	matchWrongHead, _ := regexp.MatchString(common.TALOS_NAME_HEAD_REGEX, str)
	if !(matchBody && !matchWrongHead) || len(str) > 80 {
		err := fmt.Errorf("invalid string: %s. Its only with regex set:"+
			" [a-zA-Z0-9_-] and cannot start with '_' or '-',"+
			" and must be less than 80 ", str)
		return err
	}
	return nil
}

func GenerateClientId() string {
	mils := time.Now().UnixNano() / 1000000
	uuid, _ := uuid.NewV4()
	return fmt.Sprintf("%d%s", mils, uuid.String()[0:8])
}

func CheckAndGenerateClientId(prefix string) (string, error) {
	err := CheckNameValidity(prefix)
	if err != nil {
		return "", err
	}
	return prefix + GenerateClientId(), nil
}

func CheckStartOffsetValidity(startOffset int64) *client.TalosRuntimeError {
	if startOffset >= 0 || startOffset == int64(message.MessageOffset_START_OFFSET) ||
		startOffset == int64(message.MessageOffset_LATEST_OFFSET) {
		return nil
	} else {
		errCode := common.ErrorCode_UNEXPECTED_MESSAGE_OFFSET
		err := fmt.Errorf("invalid startOffset: %d. It must be greater than "+
			"or equal to 0, or equal to Message_START_OFFSET/LATEST_OFFSET", startOffset)
		return client.NewTalosRuntimeError(errCode, err)
	}
}

func GenerateRequestSequenceId(clientId string, requestId int64) (string, *client.TalosRuntimeError) {
	if err := CheckNameValidity(clientId); err != nil {
		errCode := common.ErrorCode_UNEXPECTED_ERROR
		return "", client.NewTalosRuntimeError(errCode, err)
	}
	req := atomic.AddInt64(&requestId, 1)
	return fmt.Sprintf("%s%s%d", clientId, common.TALOS_IDENTIFIER_DELIMITER, req), nil
}

func CurrentTimeMills() int64 {
	return time.Now().UnixNano() / 1000000
}

func Serialize(tStruct thrift.TStruct) ([]byte, error) {
	transport := thrift.NewTMemoryBufferLen(1024)
	protocol := thrift.NewTCompactProtocolFactory().GetProtocol(transport)
	serializer := thrift.TSerializer{Transport: transport, Protocol: protocol}
	return serializer.Write(tStruct)
}

func HashCode(value []rune) int {
	h := 0
	if len(value) > 0 {
		val := value
		for i := 0; i < len(value); i++ {
			h = 31*h + int(val[i])
		}
	}
	return h
}

func IsTopicNotExist(err *client.TalosRuntimeError) bool {
	return err.ErrorCode == common.ErrorCode_TOPIC_NOT_EXIST
}

func IsPartitionNotServing(err *client.TalosRuntimeError) bool {
	return err.ErrorCode == common.ErrorCode_PARTITION_NOT_SERVING
}

func IsOffsetOutOfRange(err *client.TalosRuntimeError) bool {
	return err.ErrorCode == common.ErrorCode_MESSAGE_OFFSET_OUT_OF_RANGE
}

func IsUnexpectedError(err *client.TalosRuntimeError) bool {
	return err.ErrorCode == common.ErrorCode_UNEXPECTED_ERROR
}
