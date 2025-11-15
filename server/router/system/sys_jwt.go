package system

import (
	"github.com/cloudwego/hertz/pkg/route"
)

type JwtRouter struct{}

func (s *JwtRouter) InitJwtRouter(Router *route.RouterGroup) {
	jwtRouter := Router.Group("jwt")
	{
		jwtRouter.POST("jsonInBlacklist", jwtApi.JsonInBlacklist) // jwt加入黑名单
	}
}
