package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/cloudwego/hertz/pkg/route"
)

type CasbinRouter struct{}

func (s *CasbinRouter) InitCasbinRouter(Router *route.RouterGroup) {
	casbinRouter := Router.Group("casbin").Use(middleware.OperationRecord())
	casbinRouterWithoutRecord := Router.Group("casbin")
	{
		casbinRouter.POST("updateCasbin", casbinApi.UpdateCasbin)
	}
	{
		casbinRouterWithoutRecord.POST("getPolicyPathByAuthorityId", casbinApi.GetPolicyPathByAuthorityId)
	}
}
