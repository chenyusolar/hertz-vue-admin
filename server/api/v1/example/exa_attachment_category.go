package example

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	common "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	"github.com/cloudwego/hertz/pkg/app"
	"go.uber.org/zap"
)

type AttachmentCategoryApi struct{}

// GetCategoryList
// @Tags      GetCategoryList
// @Summary   媒体库分类列表
// @Security  AttachmentCategory
// @Produce   application/json
// @Success   200   {object}  response.Response{data=example.ExaAttachmentCategory,msg=string}  "媒体库分类列表"
// @Router    /attachmentCategory/getCategoryList [get]
func (a *AttachmentCategoryApi) GetCategoryList(ctx context.Context, c *app.RequestContext) {
	res, err := attachmentCategoryService.GetCategoryList()
	if err != nil {
		global.GVA_LOG.Error("获取分类列表失败!", zap.Error(err))
		response.FailWithMessage("获取分类列表失败", c)
		return
	}
	response.OkWithData(res, c)
}

// AddCategory
// @Tags      AddCategory
// @Summary   添加媒体库分类
// @Security  AttachmentCategory
// @accept    application/json
// @Produce   application/json
// @Param     data  body      example.ExaAttachmentCategory  true  "媒体库分类数据"// @Success   200   {object}  response.Response{msg=string}   "添加媒体库分类"
// @Router    /attachmentCategory/addCategory [post]
func (a *AttachmentCategoryApi) AddCategory(ctx context.Context, c *app.RequestContext) {
	var req example.ExaAttachmentCategory
	if err := c.BindJSON(&req); err != nil {
		global.GVA_LOG.Error("参数错误!", zap.Error(err))
		response.FailWithMessage("参数错误", c)
		return
	}

	if err := attachmentCategoryService.AddCategory(&req); err != nil {
		global.GVA_LOG.Error("创建/更新失败!", zap.Error(err))
		response.FailWithMessage("创建/更新失败："+err.Error(), c)
		return
	}
	response.OkWithMessage("创建/更新成功", c)
}

// DeleteCategory
// @Tags      DeleteCategory
// @Summary   删除分类
// @Security  AttachmentCategory
// @accept    application/json
// @Produce   application/json
// @Param     data  body      common.GetById                true  "分类id"
// @Success   200   {object}  response.Response{msg=string}  "删除分类"
// @Router    /attachmentCategory/deleteCategory [post]
func (a *AttachmentCategoryApi) DeleteCategory(ctx context.Context, c *app.RequestContext) {
	var req common.GetById
	if err := c.BindJSON(&req); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	if req.ID == 0 {
		response.FailWithMessage("参数错误", c)
		return
	}

	if err := attachmentCategoryService.DeleteCategory(&req.ID); err != nil {
		response.FailWithMessage("删除失败", c)
		return
	}

	response.OkWithMessage("删除成功", c)
}
