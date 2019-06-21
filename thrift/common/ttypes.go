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

var GoUnusedProtection__ int

//HTTP状态码列表，用于传输层，签名错误等
type HttpStatusCode int64

const (
	HttpStatusCode_BAD_REQUEST       HttpStatusCode = 400
	HttpStatusCode_INVALID_AUTH      HttpStatusCode = 401
	HttpStatusCode_CLOCK_TOO_SKEWED  HttpStatusCode = 412
	HttpStatusCode_REQUEST_TOO_LARGE HttpStatusCode = 413
	HttpStatusCode_INTERNAL_ERROR    HttpStatusCode = 500
)

func (p HttpStatusCode) String() string {
	switch p {
	case HttpStatusCode_BAD_REQUEST:
		return "HttpStatusCode_BAD_REQUEST"
	case HttpStatusCode_INVALID_AUTH:
		return "HttpStatusCode_INVALID_AUTH"
	case HttpStatusCode_CLOCK_TOO_SKEWED:
		return "HttpStatusCode_CLOCK_TOO_SKEWED"
	case HttpStatusCode_REQUEST_TOO_LARGE:
		return "HttpStatusCode_REQUEST_TOO_LARGE"
	case HttpStatusCode_INTERNAL_ERROR:
		return "HttpStatusCode_INTERNAL_ERROR"
	}
	return "<UNSET>"
}

func HttpStatusCodeFromString(s string) (HttpStatusCode, error) {
	switch s {
	case "HttpStatusCode_BAD_REQUEST":
		return HttpStatusCode_BAD_REQUEST, nil
	case "HttpStatusCode_INVALID_AUTH":
		return HttpStatusCode_INVALID_AUTH, nil
	case "HttpStatusCode_CLOCK_TOO_SKEWED":
		return HttpStatusCode_CLOCK_TOO_SKEWED, nil
	case "HttpStatusCode_REQUEST_TOO_LARGE":
		return HttpStatusCode_REQUEST_TOO_LARGE, nil
	case "HttpStatusCode_INTERNAL_ERROR":
		return HttpStatusCode_INTERNAL_ERROR, nil
	}
	return HttpStatusCode(0), fmt.Errorf("not a valid HttpStatusCode string")
}

func HttpStatusCodePtr(v HttpStatusCode) *HttpStatusCode { return &v }

type ErrorCode int64

const (
	ErrorCode_NO_ERROR                      ErrorCode = 0
	ErrorCode_UNKNOWN                       ErrorCode = 1
	ErrorCode_UNEXPECTED_ERROR              ErrorCode = 2
	ErrorCode_UNEXPECTED_MESSAGE_OFFSET     ErrorCode = 3
	ErrorCode_TOPIC_EXIST                   ErrorCode = 4
	ErrorCode_TOPIC_NOT_EXIST               ErrorCode = 5
	ErrorCode_QUOTA_NOT_EXIST               ErrorCode = 6
	ErrorCode_LOCK_EXIST                    ErrorCode = 7
	ErrorCode_LOCK_NOT_EXIST                ErrorCode = 8
	ErrorCode_LOCK_VALUE_NOT_EXPECTED       ErrorCode = 9
	ErrorCode_MESSAGE_MISSING               ErrorCode = 10
	ErrorCode_MESSAGE_INCOMPLETE            ErrorCode = 11
	ErrorCode_MESSAGE_INDEX_UNDESIRED_ERROR ErrorCode = 12
	ErrorCode_MESSAGE_INDEX_NOT_EXIST       ErrorCode = 13
	ErrorCode_MESSAGE_OFFSET_OUT_OF_RANGE   ErrorCode = 14
	ErrorCode_INVALID_NAME_ERROR            ErrorCode = 15
	ErrorCode_INVALID_TOPIC_PARAMS          ErrorCode = 16
	ErrorCode_OPERATION_FAILED              ErrorCode = 17
	ErrorCode_HDFS_OPERATION_FAILED         ErrorCode = 18
	ErrorCode_HBASE_OPERATION_FAILED        ErrorCode = 19
	ErrorCode_ZOOKEEPER_OPERATION_FAILED    ErrorCode = 20
	ErrorCode_PARTITION_NOT_SERVING         ErrorCode = 21
	ErrorCode_PARTITION_NOT_EXIST           ErrorCode = 22
	ErrorCode_ZK_NODE_EXIST                 ErrorCode = 23
	ErrorCode_ZK_NODE_NOT_EXIST             ErrorCode = 24
	ErrorCode_REST_SERVER_INIT_ERROR        ErrorCode = 25
	ErrorCode_INTERNAL_SERVER_ERROR         ErrorCode = 26
	ErrorCode_EXCESSIVE_PENDING_MESSAGE     ErrorCode = 27
	ErrorCode_PERMISSION_DENIED_ERROR       ErrorCode = 28
	ErrorCode_HDFS_FILE_NOT_EXIST           ErrorCode = 29
	ErrorCode_INVALID_AUTH_INFO             ErrorCode = 30
	ErrorCode_SUB_RESOURCE_NAME_EXIST       ErrorCode = 31
	ErrorCode_SUB_RESOURCE_NAME_NOT_EXIST   ErrorCode = 32
	ErrorCode_REQUEST_PROCESS_TIMEOUT       ErrorCode = 33
	ErrorCode_QUOTA_EXCEEDED                ErrorCode = 34
	ErrorCode_THROTTLE_REJECT_ERROR         ErrorCode = 35
	ErrorCode_QUOTA_AUTO_APPROVE_FAILED     ErrorCode = 36
	ErrorCode_HBASE_OPERATION_BLOCKED       ErrorCode = 37
	ErrorCode_CLOCK_TOO_SKEWED              ErrorCode = 38
	ErrorCode_REQUEST_TOO_LARGE             ErrorCode = 39
	ErrorCode_BAD_REQUEST                   ErrorCode = 40
)

func (p ErrorCode) String() string {
	switch p {
	case ErrorCode_NO_ERROR:
		return "ErrorCode_NO_ERROR"
	case ErrorCode_UNKNOWN:
		return "ErrorCode_UNKNOWN"
	case ErrorCode_UNEXPECTED_ERROR:
		return "ErrorCode_UNEXPECTED_ERROR"
	case ErrorCode_UNEXPECTED_MESSAGE_OFFSET:
		return "ErrorCode_UNEXPECTED_MESSAGE_OFFSET"
	case ErrorCode_TOPIC_EXIST:
		return "ErrorCode_TOPIC_EXIST"
	case ErrorCode_TOPIC_NOT_EXIST:
		return "ErrorCode_TOPIC_NOT_EXIST"
	case ErrorCode_QUOTA_NOT_EXIST:
		return "ErrorCode_QUOTA_NOT_EXIST"
	case ErrorCode_LOCK_EXIST:
		return "ErrorCode_LOCK_EXIST"
	case ErrorCode_LOCK_NOT_EXIST:
		return "ErrorCode_LOCK_NOT_EXIST"
	case ErrorCode_LOCK_VALUE_NOT_EXPECTED:
		return "ErrorCode_LOCK_VALUE_NOT_EXPECTED"
	case ErrorCode_MESSAGE_MISSING:
		return "ErrorCode_MESSAGE_MISSING"
	case ErrorCode_MESSAGE_INCOMPLETE:
		return "ErrorCode_MESSAGE_INCOMPLETE"
	case ErrorCode_MESSAGE_INDEX_UNDESIRED_ERROR:
		return "ErrorCode_MESSAGE_INDEX_UNDESIRED_ERROR"
	case ErrorCode_MESSAGE_INDEX_NOT_EXIST:
		return "ErrorCode_MESSAGE_INDEX_NOT_EXIST"
	case ErrorCode_MESSAGE_OFFSET_OUT_OF_RANGE:
		return "ErrorCode_MESSAGE_OFFSET_OUT_OF_RANGE"
	case ErrorCode_INVALID_NAME_ERROR:
		return "ErrorCode_INVALID_NAME_ERROR"
	case ErrorCode_INVALID_TOPIC_PARAMS:
		return "ErrorCode_INVALID_TOPIC_PARAMS"
	case ErrorCode_OPERATION_FAILED:
		return "ErrorCode_OPERATION_FAILED"
	case ErrorCode_HDFS_OPERATION_FAILED:
		return "ErrorCode_HDFS_OPERATION_FAILED"
	case ErrorCode_HBASE_OPERATION_FAILED:
		return "ErrorCode_HBASE_OPERATION_FAILED"
	case ErrorCode_ZOOKEEPER_OPERATION_FAILED:
		return "ErrorCode_ZOOKEEPER_OPERATION_FAILED"
	case ErrorCode_PARTITION_NOT_SERVING:
		return "ErrorCode_PARTITION_NOT_SERVING"
	case ErrorCode_PARTITION_NOT_EXIST:
		return "ErrorCode_PARTITION_NOT_EXIST"
	case ErrorCode_ZK_NODE_EXIST:
		return "ErrorCode_ZK_NODE_EXIST"
	case ErrorCode_ZK_NODE_NOT_EXIST:
		return "ErrorCode_ZK_NODE_NOT_EXIST"
	case ErrorCode_REST_SERVER_INIT_ERROR:
		return "ErrorCode_REST_SERVER_INIT_ERROR"
	case ErrorCode_INTERNAL_SERVER_ERROR:
		return "ErrorCode_INTERNAL_SERVER_ERROR"
	case ErrorCode_EXCESSIVE_PENDING_MESSAGE:
		return "ErrorCode_EXCESSIVE_PENDING_MESSAGE"
	case ErrorCode_PERMISSION_DENIED_ERROR:
		return "ErrorCode_PERMISSION_DENIED_ERROR"
	case ErrorCode_HDFS_FILE_NOT_EXIST:
		return "ErrorCode_HDFS_FILE_NOT_EXIST"
	case ErrorCode_INVALID_AUTH_INFO:
		return "ErrorCode_INVALID_AUTH_INFO"
	case ErrorCode_SUB_RESOURCE_NAME_EXIST:
		return "ErrorCode_SUB_RESOURCE_NAME_EXIST"
	case ErrorCode_SUB_RESOURCE_NAME_NOT_EXIST:
		return "ErrorCode_SUB_RESOURCE_NAME_NOT_EXIST"
	case ErrorCode_REQUEST_PROCESS_TIMEOUT:
		return "ErrorCode_REQUEST_PROCESS_TIMEOUT"
	case ErrorCode_QUOTA_EXCEEDED:
		return "ErrorCode_QUOTA_EXCEEDED"
	case ErrorCode_THROTTLE_REJECT_ERROR:
		return "ErrorCode_THROTTLE_REJECT_ERROR"
	case ErrorCode_QUOTA_AUTO_APPROVE_FAILED:
		return "ErrorCode_QUOTA_AUTO_APPROVE_FAILED"
	case ErrorCode_HBASE_OPERATION_BLOCKED:
		return "ErrorCode_HBASE_OPERATION_BLOCKED"
	case ErrorCode_CLOCK_TOO_SKEWED:
		return "ErrorCode_CLOCK_TOO_SKEWED"
	case ErrorCode_REQUEST_TOO_LARGE:
		return "ErrorCode_REQUEST_TOO_LARGE"
	case ErrorCode_BAD_REQUEST:
		return "ErrorCode_BAD_REQUEST"
	}
	return "<UNSET>"
}

func ErrorCodeFromString(s string) (ErrorCode, error) {
	switch s {
	case "ErrorCode_NO_ERROR":
		return ErrorCode_NO_ERROR, nil
	case "ErrorCode_UNKNOWN":
		return ErrorCode_UNKNOWN, nil
	case "ErrorCode_UNEXPECTED_ERROR":
		return ErrorCode_UNEXPECTED_ERROR, nil
	case "ErrorCode_UNEXPECTED_MESSAGE_OFFSET":
		return ErrorCode_UNEXPECTED_MESSAGE_OFFSET, nil
	case "ErrorCode_TOPIC_EXIST":
		return ErrorCode_TOPIC_EXIST, nil
	case "ErrorCode_TOPIC_NOT_EXIST":
		return ErrorCode_TOPIC_NOT_EXIST, nil
	case "ErrorCode_QUOTA_NOT_EXIST":
		return ErrorCode_QUOTA_NOT_EXIST, nil
	case "ErrorCode_LOCK_EXIST":
		return ErrorCode_LOCK_EXIST, nil
	case "ErrorCode_LOCK_NOT_EXIST":
		return ErrorCode_LOCK_NOT_EXIST, nil
	case "ErrorCode_LOCK_VALUE_NOT_EXPECTED":
		return ErrorCode_LOCK_VALUE_NOT_EXPECTED, nil
	case "ErrorCode_MESSAGE_MISSING":
		return ErrorCode_MESSAGE_MISSING, nil
	case "ErrorCode_MESSAGE_INCOMPLETE":
		return ErrorCode_MESSAGE_INCOMPLETE, nil
	case "ErrorCode_MESSAGE_INDEX_UNDESIRED_ERROR":
		return ErrorCode_MESSAGE_INDEX_UNDESIRED_ERROR, nil
	case "ErrorCode_MESSAGE_INDEX_NOT_EXIST":
		return ErrorCode_MESSAGE_INDEX_NOT_EXIST, nil
	case "ErrorCode_MESSAGE_OFFSET_OUT_OF_RANGE":
		return ErrorCode_MESSAGE_OFFSET_OUT_OF_RANGE, nil
	case "ErrorCode_INVALID_NAME_ERROR":
		return ErrorCode_INVALID_NAME_ERROR, nil
	case "ErrorCode_INVALID_TOPIC_PARAMS":
		return ErrorCode_INVALID_TOPIC_PARAMS, nil
	case "ErrorCode_OPERATION_FAILED":
		return ErrorCode_OPERATION_FAILED, nil
	case "ErrorCode_HDFS_OPERATION_FAILED":
		return ErrorCode_HDFS_OPERATION_FAILED, nil
	case "ErrorCode_HBASE_OPERATION_FAILED":
		return ErrorCode_HBASE_OPERATION_FAILED, nil
	case "ErrorCode_ZOOKEEPER_OPERATION_FAILED":
		return ErrorCode_ZOOKEEPER_OPERATION_FAILED, nil
	case "ErrorCode_PARTITION_NOT_SERVING":
		return ErrorCode_PARTITION_NOT_SERVING, nil
	case "ErrorCode_PARTITION_NOT_EXIST":
		return ErrorCode_PARTITION_NOT_EXIST, nil
	case "ErrorCode_ZK_NODE_EXIST":
		return ErrorCode_ZK_NODE_EXIST, nil
	case "ErrorCode_ZK_NODE_NOT_EXIST":
		return ErrorCode_ZK_NODE_NOT_EXIST, nil
	case "ErrorCode_REST_SERVER_INIT_ERROR":
		return ErrorCode_REST_SERVER_INIT_ERROR, nil
	case "ErrorCode_INTERNAL_SERVER_ERROR":
		return ErrorCode_INTERNAL_SERVER_ERROR, nil
	case "ErrorCode_EXCESSIVE_PENDING_MESSAGE":
		return ErrorCode_EXCESSIVE_PENDING_MESSAGE, nil
	case "ErrorCode_PERMISSION_DENIED_ERROR":
		return ErrorCode_PERMISSION_DENIED_ERROR, nil
	case "ErrorCode_HDFS_FILE_NOT_EXIST":
		return ErrorCode_HDFS_FILE_NOT_EXIST, nil
	case "ErrorCode_INVALID_AUTH_INFO":
		return ErrorCode_INVALID_AUTH_INFO, nil
	case "ErrorCode_SUB_RESOURCE_NAME_EXIST":
		return ErrorCode_SUB_RESOURCE_NAME_EXIST, nil
	case "ErrorCode_SUB_RESOURCE_NAME_NOT_EXIST":
		return ErrorCode_SUB_RESOURCE_NAME_NOT_EXIST, nil
	case "ErrorCode_REQUEST_PROCESS_TIMEOUT":
		return ErrorCode_REQUEST_PROCESS_TIMEOUT, nil
	case "ErrorCode_QUOTA_EXCEEDED":
		return ErrorCode_QUOTA_EXCEEDED, nil
	case "ErrorCode_THROTTLE_REJECT_ERROR":
		return ErrorCode_THROTTLE_REJECT_ERROR, nil
	case "ErrorCode_QUOTA_AUTO_APPROVE_FAILED":
		return ErrorCode_QUOTA_AUTO_APPROVE_FAILED, nil
	case "ErrorCode_HBASE_OPERATION_BLOCKED":
		return ErrorCode_HBASE_OPERATION_BLOCKED, nil
	case "ErrorCode_CLOCK_TOO_SKEWED":
		return ErrorCode_CLOCK_TOO_SKEWED, nil
	case "ErrorCode_REQUEST_TOO_LARGE":
		return ErrorCode_REQUEST_TOO_LARGE, nil
	case "ErrorCode_BAD_REQUEST":
		return ErrorCode_BAD_REQUEST, nil
	}
	return ErrorCode(0), fmt.Errorf("not a valid ErrorCode string")
}

func ErrorCodePtr(v ErrorCode) *ErrorCode { return &v }

type RetryType int64

const (
	RetryType_SAFE   RetryType = 0
	RetryType_UNSAFE RetryType = 1
)

func (p RetryType) String() string {
	switch p {
	case RetryType_SAFE:
		return "RetryType_SAFE"
	case RetryType_UNSAFE:
		return "RetryType_UNSAFE"
	}
	return "<UNSET>"
}

func RetryTypeFromString(s string) (RetryType, error) {
	switch s {
	case "RetryType_SAFE":
		return RetryType_SAFE, nil
	case "RetryType_UNSAFE":
		return RetryType_UNSAFE, nil
	}
	return RetryType(0), fmt.Errorf("not a valid RetryType string")
}

func RetryTypePtr(v RetryType) *RetryType { return &v }

type GalaxyTalosException struct {
	ErrorCode *ErrorCode `thrift:"errorCode,1" json:"errorCode"`
	ErrMsg    *string    `thrift:"errMsg,2" json:"errMsg"`
	Details   *string    `thrift:"details,3" json:"details"`
}

func NewGalaxyTalosException() *GalaxyTalosException {
	return &GalaxyTalosException{}
}

var GalaxyTalosException_ErrorCode_DEFAULT ErrorCode

func (p *GalaxyTalosException) GetErrorCode() ErrorCode {
	if !p.IsSetErrorCode() {
		return GalaxyTalosException_ErrorCode_DEFAULT
	}
	return *p.ErrorCode
}

var GalaxyTalosException_ErrMsg_DEFAULT string

func (p *GalaxyTalosException) GetErrMsg() string {
	if !p.IsSetErrMsg() {
		return GalaxyTalosException_ErrMsg_DEFAULT
	}
	return *p.ErrMsg
}

var GalaxyTalosException_Details_DEFAULT string

func (p *GalaxyTalosException) GetDetails() string {
	if !p.IsSetDetails() {
		return GalaxyTalosException_Details_DEFAULT
	}
	return *p.Details
}
func (p *GalaxyTalosException) IsSetErrorCode() bool {
	return p.ErrorCode != nil
}

func (p *GalaxyTalosException) IsSetErrMsg() bool {
	return p.ErrMsg != nil
}

func (p *GalaxyTalosException) IsSetDetails() bool {
	return p.Details != nil
}

func (p *GalaxyTalosException) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.ReadField3(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *GalaxyTalosException) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		temp := ErrorCode(v)
		p.ErrorCode = &temp
	}
	return nil
}

func (p *GalaxyTalosException) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		p.ErrMsg = &v
	}
	return nil
}

func (p *GalaxyTalosException) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 3: %s", err)
	} else {
		p.Details = &v
	}
	return nil
}

func (p *GalaxyTalosException) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("GalaxyTalosException"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *GalaxyTalosException) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetErrorCode() {
		if err := oprot.WriteFieldBegin("errorCode", thrift.I32, 1); err != nil {
			return fmt.Errorf("%T write field begin error 1:errorCode: %s", p, err)
		}
		if err := oprot.WriteI32(int32(*p.ErrorCode)); err != nil {
			return fmt.Errorf("%T.errorCode (1) field write error: %s", p, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 1:errorCode: %s", p, err)
		}
	}
	return err
}

func (p *GalaxyTalosException) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetErrMsg() {
		if err := oprot.WriteFieldBegin("errMsg", thrift.STRING, 2); err != nil {
			return fmt.Errorf("%T write field begin error 2:errMsg: %s", p, err)
		}
		if err := oprot.WriteString(string(*p.ErrMsg)); err != nil {
			return fmt.Errorf("%T.errMsg (2) field write error: %s", p, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 2:errMsg: %s", p, err)
		}
	}
	return err
}

func (p *GalaxyTalosException) writeField3(oprot thrift.TProtocol) (err error) {
	if p.IsSetDetails() {
		if err := oprot.WriteFieldBegin("details", thrift.STRING, 3); err != nil {
			return fmt.Errorf("%T write field begin error 3:details: %s", p, err)
		}
		if err := oprot.WriteString(string(*p.Details)); err != nil {
			return fmt.Errorf("%T.details (3) field write error: %s", p, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 3:details: %s", p, err)
		}
	}
	return err
}

func (p *GalaxyTalosException) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GalaxyTalosException(%+v)", *p)
}

func (p *GalaxyTalosException) Error() string {
	return p.String()
}

type Version struct {
	Major    int32  `thrift:"major,1,required" json:"major"`
	Minor    int32  `thrift:"minor,2,required" json:"minor"`
	Revision int32  `thrift:"revision,3,required" json:"revision"`
	Date     string `thrift:"date,4,required" json:"date"`
	Details  string `thrift:"details,5" json:"details"`
}

func NewVersion() *Version {
	return &Version{
		Major: 1,

		Date: "19700101",
	}
}

func (p *Version) GetMajor() int32 {
	return p.Major
}

func (p *Version) GetMinor() int32 {
	return p.Minor
}

func (p *Version) GetRevision() int32 {
	return p.Revision
}

func (p *Version) GetDate() string {
	return p.Date
}

var Version_Details_DEFAULT string = ""

func (p *Version) GetDetails() string {
	return p.Details
}
func (p *Version) IsSetDetails() bool {
	return p.Details != Version_Details_DEFAULT
}

func (p *Version) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.ReadField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.ReadField4(iprot); err != nil {
				return err
			}
		case 5:
			if err := p.ReadField5(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *Version) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.Major = v
	}
	return nil
}

func (p *Version) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		p.Minor = v
	}
	return nil
}

func (p *Version) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 3: %s", err)
	} else {
		p.Revision = v
	}
	return nil
}

func (p *Version) ReadField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 4: %s", err)
	} else {
		p.Date = v
	}
	return nil
}

func (p *Version) ReadField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 5: %s", err)
	} else {
		p.Details = v
	}
	return nil
}

func (p *Version) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Version"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := p.writeField5(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *Version) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("major", thrift.I32, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:major: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.Major)); err != nil {
		return fmt.Errorf("%T.major (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:major: %s", p, err)
	}
	return err
}

func (p *Version) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("minor", thrift.I32, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:minor: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.Minor)); err != nil {
		return fmt.Errorf("%T.minor (2) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:minor: %s", p, err)
	}
	return err
}

func (p *Version) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("revision", thrift.I32, 3); err != nil {
		return fmt.Errorf("%T write field begin error 3:revision: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.Revision)); err != nil {
		return fmt.Errorf("%T.revision (3) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 3:revision: %s", p, err)
	}
	return err
}

func (p *Version) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("date", thrift.STRING, 4); err != nil {
		return fmt.Errorf("%T write field begin error 4:date: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Date)); err != nil {
		return fmt.Errorf("%T.date (4) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 4:date: %s", p, err)
	}
	return err
}

func (p *Version) writeField5(oprot thrift.TProtocol) (err error) {
	if p.IsSetDetails() {
		if err := oprot.WriteFieldBegin("details", thrift.STRING, 5); err != nil {
			return fmt.Errorf("%T write field begin error 5:details: %s", p, err)
		}
		if err := oprot.WriteString(string(p.Details)); err != nil {
			return fmt.Errorf("%T.details (5) field write error: %s", p, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 5:details: %s", p, err)
		}
	}
	return err
}

func (p *Version) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Version(%+v)", *p)
}
