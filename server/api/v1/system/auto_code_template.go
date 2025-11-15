package system

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"go.uber.org/zap"
)

type AutoCodeTemplateApi struct{}

// Preview
// @Tags      AutoCodeTemplate
// @Summary   预览创建后的代码
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.AutoCode                                      true  "预览创建代码"
// @Success   200   {object}  response.Response{data=map[string]interface{},msg=string}  "预览创建后的代码"
// @Router    /autoCode/preview [post]
func (a *AutoCodeTemplateApi) Preview(ctx context.Context, c *app.RequestContext) {
	var info request.AutoCode
	err := c.BindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(info, utils.AutoCodeVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = info.Pretreatment()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	info.PackageT = utils.FirstUpper(info.Package)
	autoCode, err := autoCodeTemplateService.Preview(ctx, info)
	if err != nil {
		global.GVA_LOG.Error(err.Error(), zap.Error(err))
		response.FailWithMessage("预览失败:"+err.Error(), c)
	} else {
		response.OkWithDetailed(map[string]interface{}{"autoCode": autoCode}, "预览成功", c)
	}
}

// Create
// @Tags      AutoCodeTemplate
// @Summary   自动代码模板
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.AutoCode  true  "创建自动代码"
// @Success   200   {string}  string                 "{"success":true,"data":{},"msg":"创建成功"}"
// @Router    /autoCode/createTemp [post]
func (a *AutoCodeTemplateApi) Create(ctx context.Context, c *app.RequestContext) {
	var info request.AutoCode
	err := c.BindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(info, utils.AutoCodeVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = info.Pretreatment()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = autoCodeTemplateService.Create(ctx, info)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// AddFunc
// @Tags      AddFunc
// @Summary   增加方法
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.AutoCode  true  "增加方法"
// @Success   200   {string}  string                 "{"success":true,"data":{},"msg":"创建成功"}"
// @Router    /autoCode/addFunc [post]
func (a *AutoCodeTemplateApi) AddFunc(ctx context.Context, c *app.RequestContext) {
	var info request.AutoFunc
	err := c.BindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var tempMap map[string]string
	if info.IsPreview {
		info.Router = "填充router"
		info.FuncName = "填充funcName"
		info.Method = "填充method"
		info.Description = "填充description"
		tempMap, err = autoCodeTemplateService.GetApiAndServer(info)
	} else {
		err = autoCodeTemplateService.AddFunc(info)
	}
	if err != nil {
		global.GVA_LOG.Error("注入失败!", zap.Error(err))
		response.FailWithMessage("注入失败", c)
	} else {
		if info.IsPreview {
			response.OkWithDetailed(tempMap, "注入成功", c)
			return
		}
		response.OkWithMessage("注入成功", c)
	}
}
