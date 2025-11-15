package hertz_compat

import (

	"github.com/cloudwego/hertz/pkg/app"
)

// Context 是Hertz RequestContext的包装器
type Context struct {
	Ctx context.Context
	C   *app.RequestContext
}

// RouterGroup 是Hertz路由组的包装器
type RouterGroup = route.RouterGroup

// HandlerFunc 是Hertz处理函数类型
type HandlerFunc func(ctx context.Context, c *app.RequestContext)

// IRoutes 接口
type IRoutes = route.IRoutes

// Engine 是Hertz服务器的包装器（虽然很少直接使用）
// 我们用server.Hertz来代替
