package example

import (
	"github.com/cloudwego/hertz/pkg/route"
)

type AttachmentCategoryRouter struct{}

func (r *AttachmentCategoryRouter) InitAttachmentCategoryRouterRouter(Router *route.RouterGroup) {
	router := Router.Group("attachmentCategory")
	{
		router.GET("getCategoryList", attachmentCategoryApi.GetCategoryList) // 分类列表
		router.POST("addCategory", attachmentCategoryApi.AddCategory)        // 添加/编辑分类
		router.POST("deleteCategory", attachmentCategoryApi.DeleteCategory)  // 删除分类
	}
}
