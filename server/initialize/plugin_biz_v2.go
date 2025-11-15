package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/announcement"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/plugin/v2"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func PluginInitV2(group *server.Hertz, plugins ...plugin.Plugin) {
	for i := 0; i < len(plugins); i++ {
		plugins[i].Register(group)
	}
}
func bizPluginV2(engine *server.Hertz) {
	PluginInitV2(engine, announcement.Plugin)
}
