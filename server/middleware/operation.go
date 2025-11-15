package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"go.uber.org/zap"
)

var respPool sync.Pool
var bufferSize = 1024

func init() {
	respPool.New = func() interface{} {
		return make([]byte, bufferSize)
	}
}

func OperationRecord() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var body []byte
		var userId int
		if string(c.Request.Method()) != http.MethodGet {
			body = c.Request.Body()
			// Hertz的Body()方法返回body的拷贝，不需要额外处理
		} else {
			query := string(c.Request.URI().QueryString())
			query, _ = url.QueryUnescape(query)
			split := strings.Split(query, "&")
			m := make(map[string]string)
			for _, v := range split {
				kv := strings.Split(v, "=")
				if len(kv) == 2 {
					m[kv[0]] = kv[1]
				}
			}
			body, _ = json.Marshal(&m)
		}
		claims, _ := utils.GetClaims(ctx, c)
		if claims != nil && claims.BaseClaims.ID != 0 {
			userId = int(claims.BaseClaims.ID)
		} else {
			id, err := strconv.Atoi(string(c.Request.Header.Get("x-user-id")))
			if err != nil {
				userId = 0
			}
			userId = id
		}
		record := system.SysOperationRecord{
			Ip:     c.ClientIP(),
			Method: string(c.Request.Method()),
			Path:   string(c.Request.URI().Path()),
			Agent:  string(c.Request.Header.Get("User-Agent")),
			Body:   "",
			UserID: userId,
		}

		// 上传文件时候 中间件日志进行裁断操作
		if strings.Contains(string(c.GetHeader("Content-Type")), "multipart/form-data") {
			record.Body = "[文件]"
		} else {
			if len(body) > bufferSize {
				record.Body = "[超出记录长度]"
			} else {
				record.Body = string(body)
			}
		}

		// Hertz不支持直接拦截响应体，简化处理
		now := time.Now()

		c.Next(ctx)

		latency := time.Since(now)
		record.ErrorMessage = "" // Hertz不支持gin.Errors机制
		record.Status = c.Response.StatusCode()
		record.Latency = latency
		record.Resp = "" // Hertz难以直接获取响应体

		// 判断是否为文件下载
		contentType := string(c.Response.Header.Get("Content-Type"))
		contentDisposition := string(c.Response.Header.Get("Content-Disposition"))
		if strings.Contains(contentType, "application/octet-stream") ||
			strings.Contains(contentType, "application/force-download") ||
			strings.Contains(contentType, "application/vnd.ms-excel") ||
			strings.Contains(contentType, "application/download") ||
			strings.Contains(contentDisposition, "attachment") {
			// 文件下载，不记录响应体
			record.Resp = "[文件下载]"
		}
		if err := global.GVA_DB.Create(&record).Error; err != nil {
			global.GVA_LOG.Error("create operation record error:", zap.Error(err))
		}
	}
}

type responseBodyWriter struct {
	http.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
