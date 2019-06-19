/**
 * Copyright 2018, Xiaomi.
 * All rights reserved.
 * Author: wangfan8@xiaomi.com
 */

package producer

import (
	"bytes"
	"testing"

	"talos-sdk-golang/thrift/message"

	"talos-sdk-golang/producer"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/stretchr/testify/assert"
)

var msg *message.Message
var userMessage *producer.UserMessage
var data *bytes.Buffer
var partitionKey string
var sequenceNumber string

func setUp() {
	data = bytes.NewBuffer([]byte("hello"))
	partitionKey = "123456qwerty"
	sequenceNumber = "15456"
	msg = &message.Message{
		Message:      data.Bytes(),
		PartitionKey: thrift.StringPtr(partitionKey),
	}
	userMessage = producer.NewUserMessage(msg)
}

func Test_UserMessage(t *testing.T) {
	setUp()

	// test GetMessage
	assert.Equal(t, msg, userMessage.GetMessage())
	// test GetTimestamp
	assert.NotNil(t, userMessage.GetTimestamp())

	// test GetMessageSize
	assert.Equal(t, int64(data.Len()), userMessage.GetMessageSize())
	msg.SequenceNumber = &sequenceNumber
	userMessage = producer.NewUserMessage(msg)
	assert.Equal(t, int64(data.Len()+len(sequenceNumber)), userMessage.GetMessageSize())
}
