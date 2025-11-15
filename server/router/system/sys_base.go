package system

import (
	"github.com/cloudwego/hertz/pkg/route"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *route.RouterGroup) (R route.IRoutes) {
	baseRouter := Router.Group("base")
	{
		baseRouter.POST("login", baseApi.Login)
		baseRouter.POST("captcha", baseApi.Captcha)
	}
	return baseRouter
}
