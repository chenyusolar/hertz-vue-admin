package response

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 7
	SUCCESS = 0
)

func Result(code int, data interface{}, msg string, c *app.RequestContext) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(ctx context.Context, c *app.RequestContext) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *app.RequestContext) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *app.RequestContext) {
	Result(SUCCESS, data, "成功", c)
}

func OkWithDetailed(data interface{}, message string, c *app.RequestContext) {
	Result(SUCCESS, data, message, c)
}

func Fail(ctx context.Context, c *app.RequestContext) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *app.RequestContext) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func NoAuth(message string, c *app.RequestContext) {
	c.JSON(http.StatusUnauthorized, Response{
		7,
		nil,
		message,
	})
}

func FailWithDetailed(data interface{}, message string, c *app.RequestContext) {
	Result(ERROR, data, message, c)
}
