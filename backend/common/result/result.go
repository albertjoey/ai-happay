package result

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// Response 统一响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Success 成功响应
func Success(w http.ResponseWriter, data interface{}) {
	resp := Response{
		Code:    0,
		Message: "success",
		Data:    data,
	}
	httpx.WriteJson(w, http.StatusOK, resp)
}

// Fail 失败响应
func Fail(w http.ResponseWriter, code int, message string) {
	resp := Response{
		Code:    code,
		Message: message,
	}
	httpx.WriteJson(w, http.StatusOK, resp)
}
