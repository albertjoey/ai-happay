package errors

import (
	"fmt"
)

// 错误码定义
const (
	// 通用错误
	CodeSuccess         = 0
	CodeServerError     = 10000
	CodeInvalidParam    = 10001
	CodeUnauthorized    = 10002
	CodeForbidden       = 10003
	CodeNotFound        = 10004
	CodeBadRequest      = 10005
	CodeTooManyRequests = 10006

	// 用户相关错误 20xxx
	CodeUserNotFound      = 20001
	CodeUserAlreadyExists = 20002
	CodeUserPasswordError = 20003
	CodeUserDisabled      = 20004
	CodeUserNotLogin      = 20005

	// 内容相关错误 30xxx
	CodeContentNotFound      = 30001
	CodeContentAlreadyExists = 30002
	CodeContentStatusError   = 30003
	CodeContentPermissionDenied = 30004

	// 互动相关错误 40xxx
	CodeInteractionAlreadyExists = 40001
	CodeInteractionNotFound      = 40002

	// 租户相关错误 50xxx
	CodeTenantNotFound      = 50001
	CodeTenantAlreadyExists = 50002
	CodeTenantDisabled      = 50003
)

// 错误消息
var errorMsg = map[int]string{
	CodeSuccess:         "成功",
	CodeServerError:     "服务器内部错误",
	CodeInvalidParam:    "参数错误",
	CodeUnauthorized:    "未授权",
	CodeForbidden:       "无权限",
	CodeNotFound:        "资源不存在",
	CodeBadRequest:      "请求错误",
	CodeTooManyRequests: "请求过于频繁",

	CodeUserNotFound:      "用户不存在",
	CodeUserAlreadyExists: "用户已存在",
	CodeUserPasswordError: "密码错误",
	CodeUserDisabled:      "用户已被禁用",
	CodeUserNotLogin:      "用户未登录",

	CodeContentNotFound:      "内容不存在",
	CodeContentAlreadyExists: "内容已存在",
	CodeContentStatusError:   "内容状态错误",
	CodeContentPermissionDenied: "无权限操作此内容",

	CodeInteractionAlreadyExists: "已操作过",
	CodeInteractionNotFound:      "操作记录不存在",

	CodeTenantNotFound:      "租户不存在",
	CodeTenantAlreadyExists: "租户已存在",
	CodeTenantDisabled:      "租户已被禁用",
}

// Error 自定义错误类型
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("code: %d, message: %s", e.Code, e.Message)
}

// New 创建错误
func New(code int, message ...string) *Error {
	msg := ""
	if len(message) > 0 {
		msg = message[0]
	} else if m, ok := errorMsg[code]; ok {
		msg = m
	}

	return &Error{
		Code:    code,
		Message: msg,
	}
}

// GetErrorMsg 获取错误消息
func GetErrorMsg(code int) string {
	if msg, ok := errorMsg[code]; ok {
		return msg
	}
	return "未知错误"
}
