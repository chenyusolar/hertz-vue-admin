package initialize

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/flipped-aurora/gin-vue-admin/server/docs"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/router"
	hertzSwagger "github.com/hertz-contrib/swagger"
	swaggerFiles "github.com/swaggo/files"
)

// 初始化总路由

func Routers() *server.Hertz {
	// 创建Hertz服务器
	h := server.Default(
		server.WithHostPorts(fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)),
		server.WithReadTimeout(10*time.Minute),
		server.WithWriteTimeout(10*time.Minute),
	)

	if !global.GVA_CONFIG.MCP.Separate {

		_ = McpRun() // sseServer暂时未使用

		// 注册mcp服务 - SSE需要特殊适配Hertz，暂时注释
		// h.GET(global.GVA_CONFIG.MCP.SSEPath, func(ctx context.Context, c *app.RequestContext) {
		// 	req, _ := adaptor.GetCompatRequest(c.GetRequest())
		// 	sseServer.SSEHandler().ServeHTTP(c.GetResponse(), req)
		// })
		//
		// h.POST(global.GVA_CONFIG.MCP.MessagePath, func(ctx context.Context, c *app.RequestContext) {
		// 	req, _ := adaptor.GetCompatRequest(c.GetRequest())
		// 	sseServer.MessageHandler().ServeHTTP(c.GetResponse(), req)
		// })
	}

	systemRouter := router.RouterGroupApp.System
	exampleRouter := router.RouterGroupApp.Example

	// 静态文件服务
	h.StaticFS(global.GVA_CONFIG.Local.StorePath, &app.FS{Root: global.GVA_CONFIG.Local.StorePath})

	// Swagger文档
	docs.SwaggerInfo.BasePath = global.GVA_CONFIG.System.RouterPrefix
	h.GET(global.GVA_CONFIG.System.RouterPrefix+"/swagger/*any", hertzSwagger.WrapHandler(swaggerFiles.Handler))
	global.GVA_LOG.Info("register swagger handler")

	// 路由组
	PublicGroup := h.Group(global.GVA_CONFIG.System.RouterPrefix)
	PrivateGroup := h.Group(global.GVA_CONFIG.System.RouterPrefix)

	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())

	{
		// 健康监测
		PublicGroup.GET("/health", func(ctx context.Context, c *app.RequestContext) {
			c.JSON(http.StatusOK, "ok")
		})
	}
	{
		systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		systemRouter.InitInitRouter(PublicGroup) // 自动初始化相关
	}

	{
		systemRouter.InitApiRouter(PrivateGroup, PublicGroup)               // 注册功能api路由
		systemRouter.InitJwtRouter(PrivateGroup)                            // jwt相关路由
		systemRouter.InitUserRouter(PrivateGroup)                           // 注册用户路由
		systemRouter.InitMenuRouter(PrivateGroup)                           // 注册menu路由
		systemRouter.InitSystemRouter(PrivateGroup)                         // system相关路由
		systemRouter.InitSysVersionRouter(PrivateGroup)                     // 发版相关路由
		systemRouter.InitCasbinRouter(PrivateGroup)                         // 权限相关路由
		systemRouter.InitAutoCodeRouter(PrivateGroup, PublicGroup)          // 创建自动化代码
		systemRouter.InitAuthorityRouter(PrivateGroup)                      // 注册角色路由
		systemRouter.InitSysDictionaryRouter(PrivateGroup)                  // 字典管理
		systemRouter.InitAutoCodeHistoryRouter(PrivateGroup)                // 自动化代码历史
		systemRouter.InitSysOperationRecordRouter(PrivateGroup)             // 操作记录
		systemRouter.InitSysDictionaryDetailRouter(PrivateGroup)            // 字典详情管理
		systemRouter.InitAuthorityBtnRouterRouter(PrivateGroup)             // 按钮权限管理
		systemRouter.InitSysExportTemplateRouter(PrivateGroup, PublicGroup) // 导出模板
		systemRouter.InitSysParamsRouter(PrivateGroup, PublicGroup)         // 参数管理
		exampleRouter.InitCustomerRouter(PrivateGroup)                      // 客户路由
		exampleRouter.InitFileUploadAndDownloadRouter(PrivateGroup)         // 文件上传下载功能路由
		exampleRouter.InitAttachmentCategoryRouterRouter(PrivateGroup)      // 文件上传下载分类

	}

	//插件路由安装
	InstallPlugin(PrivateGroup, PublicGroup, h)

	// 注册业务路由
	initBizRouter(PrivateGroup, PublicGroup)

	global.GVA_ROUTERS = h.Routes()

	global.GVA_LOG.Info("router register success")
	return h
}
