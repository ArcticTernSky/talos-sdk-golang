// Autogenerated by Thrift Compiler (0.9.2)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package common

import (
	"bytes"
	"fmt"

	"github.com/XiaoMi/talos-sdk-golang/utils/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

const MAX_RETRY = 1

var ERROR_BACKOFF map[ErrorCode]int64
var ERROR_RETRY_TYPE map[ErrorCode]RetryType

func init() {
	ERROR_BACKOFF = map[ErrorCode]int64{
		5:  1000,
		6:  1000,
		17: 2000,
		21: 2000,
		25: 1000,
		2:  1000,
		10: 1000,
		11: 1000,
		12: 1000,
		13: 1000,
		14: 1000,
		15: 1000,
		26: 1000,
	}

	ERROR_RETRY_TYPE = map[ErrorCode]RetryType{
		5:  0,
		6:  0,
		17: 0,
		21: 0,
		25: 0,
		2:  1,
		10: 1,
		11: 1,
		12: 1,
		13: 1,
		14: 1,
		15: 1,
		26: 1,
	}

}
