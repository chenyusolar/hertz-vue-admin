package initialize

import (
	"github.com/cloudwego/hertz/pkg/route"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func InstallPlugin(PrivateGroup *route.RouterGroup, PublicRouter *route.RouterGroup, engine *server.Hertz) {
	if global.GVA_DB == nil {
		global.GVA_LOG.Info("项目暂未初始化，无法安装插件，初始化后重启项目即可完成插件安装")
		return
	}
	bizPluginV1(PrivateGroup, PublicRouter)
	bizPluginV2(engine)
}
