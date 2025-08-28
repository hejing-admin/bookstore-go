package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// Success 200 - 成功响应
func Success(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: message,
		Data:    data,
	})
}

// BadRequest 400 - 请求参数错误
func BadRequest(c *gin.Context, message string, err error) {
	c.JSON(http.StatusBadRequest, Response{
		Code:    -1,
		Message: message,
		Error:   parseErrMessage(err),
	})
}

// Unauthorized 401 - 未认证
func Unauthorized(c *gin.Context, message string, err error) {
	c.JSON(http.StatusUnauthorized, Response{
		Code:    -1,
		Message: message,
		Error:   parseErrMessage(err),
	})
}

// Forbidden 403 - 无权限
func Forbidden(c *gin.Context, message string, err error) {
	c.JSON(http.StatusForbidden, Response{
		Code:    -1,
		Message: message,
		Error:   parseErrMessage(err),
	})
}

// NotFound 404 - 资源未找到
func NotFound(c *gin.Context, message string, err error) {
	c.JSON(http.StatusNotFound, Response{
		Code:    -1,
		Message: message,
		Error:   parseErrMessage(err),
	})
}

// Conflict 409 - 资源冲突
func Conflict(c *gin.Context, message string, err error) {
	c.JSON(http.StatusConflict, Response{
		Code:    -1,
		Message: message,
		Error:   parseErrMessage(err),
	})
}

// InternalError 500 - 服务器内部错误
func InternalError(c *gin.Context, message string, err error) {
	c.JSON(http.StatusInternalServerError, Response{
		Code:    -1,
		Message: message,
		Error:   parseErrMessage(err),
	})
}

func parseErrMessage(err error) string {
	var errStr string
	if err != nil {
		errStr = err.Error()
	}
	return errStr
}
