package xerr

import (
	"errors"
	"fmt"
)

// NewErrMsg 创建错误消息
func NewErrMsg(msg string) error {
	return errors.New(msg)
}

// NewErrCode 创建带错误码的错误
func NewErrCode(code int, msg string) error {
	return fmt.Errorf("code:%d,msg:%s", code, msg)
}
