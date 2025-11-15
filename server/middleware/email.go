package middleware

import (
	"context"
	"bytes"
	"io"
	"strconv"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/email/utils"
	utils2 "github.com/flipped-aurora/gin-vue-admin/server/utils"


	"github.com/cloudwego/hertz/pkg/app"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"go.uber.org/zap"
)

func ErrorToEmail() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var username string
		claims, _ := utils2.GetClaims(ctx, c)
		if claims.Username != "" {
			username = claims.Username
		} else {
			id, _ := strconv.Atoi(string(c.Request.Header.Get("x-user-id")))
			var u system.SysUser
			err := global.GVA_DB.Where("id = ?", id).First(&u).Error
			if err != nil {
				username = "Unknown"
			}
			username = u.Username
		}
		body, _ := io.ReadAll(c.Request.BodyStream())
		// 再重新写回请求体body中，ioutil.ReadAll会清空c.Request.Body中的数据
		c.Request.SetBodyStream(bytes.NewBuffer(body), len(body))
		record := system.SysOperationRecord{
			Ip:     c.ClientIP(),
			Method: string(c.Request.Method()),
			Path:   string(c.Request.URI().Path()),
			Agent:  string(c.Request.Header.Get("User-Agent")),
			Body:   string(body),
		}
		now := time.Now()

		c.Next(ctx)

		latency := time.Since(now)
		status := c.Response.StatusCode()
		record.ErrorMessage = "" // Hertz不支持类似Gin的Errors机制
		str := "接收到的请求为" + record.Body + "\n" + "请求方式为" + record.Method + "\n" + "报错信息如下" + record.ErrorMessage + "\n" + "耗时" + latency.String() + "\n"
		if status != 200 {
			subject := username + "" + record.Ip + "调用了" + record.Path + "报错了"
			if err := utils.ErrorToEmail(subject, str); err != nil {
				global.GVA_LOG.Error("ErrorToEmail Failed, err:", zap.Error(err))
			}
		}
	}
}
