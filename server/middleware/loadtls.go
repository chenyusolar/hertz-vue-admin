package middleware

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

// 用https把这个中间件在router里面use一下就好

func LoadTls() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// Hertz不直接支持secure middleware，需要自行实现TLS逻辑
		// 可以在这里检查请求协议并重定向到HTTPS
		// 示例: if !c.Request.IsTLS() { /* redirect logic */ }
		c.Next(ctx)
	}
}
