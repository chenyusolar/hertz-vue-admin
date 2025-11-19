package initialize

import (
	"{{.Module}}/global"
	"{{.Module}}/middleware"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/route"
)

func Router(engine *server.Hertz) {
	public := engine.Group(global.GVA_CONFIG.System.RouterPrefix).Group("")
	public.Use()
	private := engine.Group(global.GVA_CONFIG.System.RouterPrefix).Group("")
	private.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
}
