package plugin

import (
	"github.com/cloudwego/hertz/pkg/app/server"
)

// Plugin 插件模式接口化v2
type Plugin interface {
	// Register 注册路由
	Register(group *server.Hertz)
}
