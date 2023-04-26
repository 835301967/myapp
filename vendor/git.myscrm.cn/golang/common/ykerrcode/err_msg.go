// 通用错误码组件-错误消息相关接口
package ykerrcode

import (
	"log"
)

var errMsgDict = make(map[int]string)

// Deprecated: use RegisterErrMsg instead.
func RegisterErrMsgDict(dict map[int]string) {
	for code, errMsg := range dict {
		if errMsgDict[code] != "" {
			log.Fatalf("错误码初始化错误,重复定义的code:%d", code)
		}
		if code < 0 || (1 <= code && code <= 9999999) || code > 29999999 {
			log.Fatalf("错误码初始化错误,不符合规范的code:%d", code)
		}
		errMsgDict[code] = errMsg
	}
}

func RegisterErrMsg(dict map[interface{}]string) {
	for code, errMsg := range dict {
		var codec int
		switch code.(type) {
		case int:
			codec = code.(int)
		case (*ErrorCode):
			codec = code.(*ErrorCode).BusinessCode
		}
		if errMsgDict[codec] != "" {
			log.Fatalf("错误码初始化错误,重复定义的code:%d", code)
		}
		if codec < 0 || (1 <= codec && codec <= 9999999) || codec > 29999999 {
			log.Fatalf("错误码初始化错误,不符合规范的code:%d", code)
		}
		errMsgDict[codec] = errMsg
	}
}

func GetErrMsg(code int) string {
	return errMsgDict[code]
}
