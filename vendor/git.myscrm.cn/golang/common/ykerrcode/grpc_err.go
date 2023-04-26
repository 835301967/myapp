// 通用错误码组件-grpc错误返回定义
package ykerrcode

import (
	"net/http"
	"strings"

	"git.myscrm.cn/golang/common/proto/common"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	sts "google.golang.org/grpc/status"
)

type ErrorCode struct {
	BusinessCode int
	StatusCode   codes.Code
}

func NewInternal(businessCode int) *ErrorCode {
	return NewErrorCode(businessCode, codes.Internal)
}

func NewUnknown(businessCode int) *ErrorCode {
	return NewErrorCode(businessCode, codes.Unknown)
}

func NewCanceled(businessCode int) *ErrorCode {
	return NewErrorCode(businessCode, codes.Canceled)
}

func NewDeadlineExceeded(businessCode int) *ErrorCode {
	return NewErrorCode(businessCode, codes.DeadlineExceeded)
}

func NewErrorCode(businessCode int, statusCode codes.Code) *ErrorCode {
	return &ErrorCode{BusinessCode: businessCode, StatusCode: statusCode}
}

func TogRPCError(code interface{}, msgs ...string) error {
	var msg string
	var s *status.Status
	var businessCode int
	var statusCode codes.Code
	switch code.(type) {
	case int:
		businessCode = code.(int)
		statusCode = ToRPCCode(int32(businessCode))
	case (*ErrorCode):
		codec := code.((*ErrorCode))
		businessCode = codec.BusinessCode
		statusCode = codec.StatusCode
	}

	msg = assemblyMessage(businessCode, msgs...)
	s, _ = sts.New(statusCode, msg).WithDetails(&common.Error{Code: int32(businessCode), Message: msg})
	return s.Err()
}

func ToRPCCode(code int32) codes.Code {
	var statusCode codes.Code
	switch code {
	case FAIL:
		statusCode = codes.Internal
	case INVALID_PARAMS:
		statusCode = codes.InvalidArgument
	case UNAUTH:
		statusCode = codes.Unauthenticated
	case ACCESS_DENIED:
		statusCode = codes.PermissionDenied
	case DEADLINE_EXCEEDED:
		statusCode = codes.DeadlineExceeded
	case NOT_FOUND:
		statusCode = codes.NotFound
	case LIMIT_EXCEED:
		statusCode = codes.ResourceExhausted
	case SERVICE_UNAVAILABLE:
		statusCode = codes.Unavailable
	case METHOD_NOT_ALLOWED:
		statusCode = codes.Unimplemented
	default:
		statusCode = codes.Unknown
	}

	return statusCode
}

func ToRPCCodeText(code codes.Code) string {
	switch code {
	case codes.OK:
		return codes.OK.String()
	case codes.Canceled:
		return codes.Canceled.String()
	case codes.Unknown:
		return codes.Unknown.String()
	case codes.InvalidArgument:
		return codes.InvalidArgument.String()
	case codes.DeadlineExceeded:
		return codes.DeadlineExceeded.String()
	case codes.NotFound:
		return codes.NotFound.String()
	case codes.AlreadyExists:
		return codes.AlreadyExists.String()
	case codes.PermissionDenied:
		return codes.PermissionDenied.String()
	case codes.ResourceExhausted:
		return codes.ResourceExhausted.String()
	case codes.FailedPrecondition:
		return codes.FailedPrecondition.String()
	case codes.Aborted:
		return codes.Aborted.String()
	case codes.OutOfRange:
		return codes.OutOfRange.String()
	case codes.Unimplemented:
		return codes.Unimplemented.String()
	case codes.Internal:
		return codes.Internal.String()
	case codes.Unavailable:
		return codes.Unavailable.String()
	case codes.DataLoss:
		return codes.DataLoss.String()
	case codes.Unauthenticated:
		return codes.Unauthenticated.String()
	}

	return ""
}

func ToHttpStatusCode(code codes.Code) int {
	var statusCode int
	switch code {
	case codes.Unknown:
		statusCode = http.StatusInternalServerError
	case codes.Internal:
		statusCode = http.StatusInternalServerError
	case codes.InvalidArgument:
		statusCode = http.StatusBadRequest
	case codes.Unauthenticated:
		statusCode = http.StatusUnauthorized
	case codes.PermissionDenied:
		statusCode = http.StatusUnauthorized
	case codes.DeadlineExceeded:
		statusCode = http.StatusRequestTimeout
	case codes.NotFound:
		statusCode = http.StatusNotFound
	case codes.ResourceExhausted:
		statusCode = http.StatusTooManyRequests
	case codes.Unimplemented:
		statusCode = http.StatusMethodNotAllowed
	case codes.Unavailable:
		statusCode = http.StatusServiceUnavailable
	default:
		statusCode = http.StatusOK
	}

	return statusCode
}

func assemblyMessage(code int, msgs ...string) string {
	var msg string
	if len(msgs) == 0 {
		msg = GetErrMsg(code)
	} else {
		msg = strings.Join(msgs, ",")
	}

	return msg
}
