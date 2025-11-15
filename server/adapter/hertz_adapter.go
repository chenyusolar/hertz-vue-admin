package adapter

import (
	"mime/multipart"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/adaptor"
	"github.com/cloudwego/hertz/pkg/protocol"
)

// Context 适配器，封装Hertz的RequestContext，提供类似Gin的接口
type Context struct {
	ctx context.Context
	c   *app.RequestContext
}

// NewContext 创建新的Context适配器
func NewContext(ctx context.Context, c *app.RequestContext) *Context {
	return &Context{
		ctx: ctx,
		c:   c,
	}
}

// ShouldBindJSON 绑定JSON数据
func (ctx *Context) ShouldBindJSON(obj interface{}) error {
	return ctx.c.BindAndValidate(obj)
}

// ShouldBindQuery 绑定query参数
func (ctx *Context) ShouldBindQuery(obj interface{}) error {
	return ctx.c.BindAndValidate(obj)
}

// ShouldBind 绑定请求数据
func (ctx *Context) ShouldBind(obj interface{}) error {
	return ctx.c.BindAndValidate(obj)
}

// JSON 返回JSON响应
func (ctx *Context) JSON(code int, obj interface{}) {
	ctx.c.JSON(code, obj)
}

// String 返回字符串响应
func (ctx *Context) String(code int, format string, values ...interface{}) {
	ctx.c.String(code, format, values...)
}

// Abort 中止请求处理
func (ctx *Context) Abort() {
	ctx.c.Abort()
}

// AbortWithStatus 带状态码中止请求
func (ctx *Context) AbortWithStatus(code int) {
	ctx.c.AbortWithStatus(code)
}

// AbortWithStatusJSON 带状态码和JSON中止请求
func (ctx *Context) AbortWithStatusJSON(code int, jsonObj interface{}) {
	ctx.c.AbortWithStatusJSON(code, jsonObj)
}

// Next 执行下一个中间件
func (ctx *Context) Next() {
	ctx.c.Next(ctx.ctx)
}

// Set 设置键值对
func (ctx *Context) Set(key string, value interface{}) {
	ctx.c.Set(key, value)
}

// Get 获取键值
func (ctx *Context) Get(key string) (value interface{}, exists bool) {
	return ctx.c.Get(key)
}

// MustGet 必须获取值
func (ctx *Context) MustGet(key string) interface{} {
	return ctx.c.MustGet(key)
}

// GetString 获取字符串值
func (ctx *Context) GetString(key string) string {
	return ctx.c.GetString(key)
}

// GetHeader 获取请求头
func (ctx *Context) GetHeader(key string) string {
	return string(ctx.c.GetHeader(key))
}

// Header 设置响应头
func (ctx *Context) Header(key, value string) {
	ctx.c.Header(key, value)
}

// Query 获取query参数
func (ctx *Context) Query(key string) string {
	return ctx.c.Query(key)
}

// DefaultQuery 获取query参数，带默认值
func (ctx *Context) DefaultQuery(key, defaultValue string) string {
	return ctx.c.DefaultQuery(key, defaultValue)
}

// Param 获取路径参数
func (ctx *Context) Param(key string) string {
	return ctx.c.Param(key)
}

// PostForm 获取post form参数
func (ctx *Context) PostForm(key string) string {
	return ctx.c.PostForm(key)
}

// DefaultPostForm 获取post form参数，带默认值
func (ctx *Context) DefaultPostForm(key, defaultValue string) string {
	return ctx.c.DefaultPostForm(key, defaultValue)
}

// FormFile 获取上传文件
func (ctx *Context) FormFile(name string) (*multipart.FileHeader, error) {
	return ctx.c.FormFile(name)
}

// MultipartForm 获取multipart form
func (ctx *Context) MultipartForm() (*multipart.Form, error) {
	return ctx.c.MultipartForm()
}

// SaveUploadedFile 保存上传文件
func (ctx *Context) SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	return ctx.c.SaveUploadedFile(file, dst)
}

// ClientIP 获取客户端IP
func (ctx *Context) ClientIP() string {
	return ctx.c.ClientIP()
}

// Cookie 获取Cookie
func (ctx *Context) Cookie(name string) (string, error) {
	return string(ctx.c.Cookie(name)), nil
}

// SetCookie 设置Cookie
func (ctx *Context) SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool) {
	cookie := &protocol.Cookie{}
	cookie.SetKey(name)
	cookie.SetValue(value)
	cookie.SetMaxAge(maxAge)
	cookie.SetPath(path)
	cookie.SetDomain(domain)
	cookie.SetSecure(secure)
	cookie.SetHTTPOnly(httpOnly)
	ctx.c.Response.Header.SetCookie(cookie)
}

// SetSameSite 设置SameSite属性
func (ctx *Context) SetSameSite(samesite http.SameSite) {
	// Hertz的SameSite设置方式不同，这里需要在SetCookie时处理
}

// Request 获取原始请求
func (ctx *Context) Request() *http.Request {
	req, _ := adaptor.GetCompatRequest(&ctx.c.Request)
	return req
}

// Writer 获取响应Writer
func (ctx *Context) Writer() http.ResponseWriter {
	return adaptor.GetCompatResponseWriter(&ctx.c.Response)
}

// FullPath 获取完整路径
func (ctx *Context) FullPath() string {
	return ctx.c.FullPath()
}

// GetRawData 获取原始请求体数据
func (ctx *Context) GetRawData() ([]byte, error) {
	body, err := ctx.c.Body()
	return body, err
}

// Bind 绑定数据
func (ctx *Context) Bind(obj interface{}) error {
	return ctx.c.BindAndValidate(obj)
}

// BindJSON 绑定JSON数据
func (ctx *Context) BindJSON(obj interface{}) error {
	return ctx.c.BindJSON(obj)
}

// BindQuery 绑定Query参数
func (ctx *Context) BindQuery(obj interface{}) error {
	return ctx.c.BindQuery(obj)
}

// Status 设置响应状态码
func (ctx *Context) Status(code int) {
	ctx.c.SetStatusCode(code)
}

// GetRequest 获取底层Hertz RequestContext
func (ctx *Context) GetRequest() *app.RequestContext {
	return ctx.c
}

// GetContext 获取context.Context
func (ctx *Context) GetContext() context.Context {
	return ctx.ctx
}

// File 发送文件响应
func (ctx *Context) File(filepath string) {
	ctx.c.File(filepath)
}

// FileAttachment 以附件形式发送文件
func (ctx *Context) FileAttachment(filepath, filename string) {
	ctx.c.FileAttachment(filepath, filename)
}

// Data 发送原始数据响应
func (ctx *Context) Data(code int, contentType string, data []byte) {
	ctx.c.Data(code, contentType, data)
}

// Stream 流式响应
func (ctx *Context) Stream(step func(w http.ResponseWriter) bool) bool {
	// Hertz的流式响应需要特殊处理
	return false
}

// SSEvent Server-Sent Event
func (ctx *Context) SSEvent(name string, message interface{}) {
	// Hertz需要通过其他方式实现SSE
	ctx.c.SetStatusCode(http.StatusOK)
	ctx.c.Response.Header.Set("Content-Type", "text/event-stream")
	ctx.c.Response.Header.Set("Cache-Control", "no-cache")
	ctx.c.Response.Header.Set("Connection", "keep-alive")
}

// GetQuery 获取query参数
func (ctx *Context) GetQuery(key string) (string, bool) {
	val := ctx.c.Query(key)
	return val, val != ""
}

// GetPostForm 获取post form参数
func (ctx *Context) GetPostForm(key string) (string, bool) {
	val := ctx.c.PostForm(key)
	return val, val != ""
}

// QueryArray 获取query数组参数
func (ctx *Context) QueryArray(key string) []string {
	var result []string
	ctx.c.QueryArgs().VisitAll(func(k, v []byte) {
		if string(k) == key {
			result = append(result, string(v))
		}
	})
	return result
}

// PostFormArray 获取post form数组参数
func (ctx *Context) PostFormArray(key string) []string {
	var result []string
	ctx.c.PostArgs().VisitAll(func(k, v []byte) {
		if string(k) == key {
			result = append(result, string(v))
		}
	})
	return result
}

// Redirect 重定向
func (ctx *Context) Redirect(code int, location string) {
	ctx.c.Redirect(code, []byte(location))
}

// QueryMap 获取query map参数
func (ctx *Context) QueryMap(key string) map[string]string {
	// 需要自己解析
	return make(map[string]string)
}

// PostFormMap 获取post form map参数
func (ctx *Context) PostFormMap(key string) map[string]string {
	// 需要自己解析
	return make(map[string]string)
}

// ContentType 返回请求Content-Type
func (ctx *Context) ContentType() string {
	return string(ctx.c.ContentType())
}

// IsWebsocket 检查是否为websocket请求
func (ctx *Context) IsWebsocket() bool {
	return string(ctx.c.GetHeader("Upgrade")) == "websocket"
}

// RemoteIP 获取远程IP
func (ctx *Context) RemoteIP() string {
	return ctx.c.RemoteAddr().String()
}

// HandlerName 获取处理器名称
func (ctx *Context) HandlerName() string {
	return ctx.c.HandlerName()
}

// Error 添加错误信息
func (ctx *Context) Error(err error) error {
	// Hertz没有直接的Error方法，需要自己处理
	return err
}

// Value 获取context中的值
func (ctx *Context) Value(key interface{}) interface{} {
	return ctx.ctx.Value(key)
}

// Keys 获取所有设置的键值对
func (ctx *Context) Keys() map[string]interface{} {
	keys := make(map[string]interface{})
	ctx.c.ForEachKey(func(k string, v interface{}) {
		keys[k] = v
	})
	return keys
}

// GetInt 获取int值
func (ctx *Context) GetInt(key string) int {
	val, _ := ctx.Get(key)
	if i, ok := val.(int); ok {
		return i
	}
	return 0
}

// GetInt64 获取int64值
func (ctx *Context) GetInt64(key string) int64 {
	val, _ := ctx.Get(key)
	if i, ok := val.(int64); ok {
		return i
	}
	return 0
}

// GetUint 获取uint值
func (ctx *Context) GetUint(key string) uint {
	val, _ := ctx.Get(key)
	if i, ok := val.(uint); ok {
		return i
	}
	return 0
}

// GetUint64 获取uint64值
func (ctx *Context) GetUint64(key string) uint64 {
	val, _ := ctx.Get(key)
	if i, ok := val.(uint64); ok {
		return i
	}
	return 0
}

// GetFloat64 获取float64值
func (ctx *Context) GetFloat64(key string) float64 {
	val, _ := ctx.Get(key)
	if f, ok := val.(float64); ok {
		return f
	}
	return 0
}

// GetBool 获取bool值
func (ctx *Context) GetBool(key string) bool {
	val, _ := ctx.Get(key)
	if b, ok := val.(bool); ok {
		return b
	}
	return false
}

// GetTime 获取time.Time值
func (ctx *Context) GetTime(key string) interface{} {
	val, _ := ctx.Get(key)
	return val
}

// GetDuration 获取Duration值
func (ctx *Context) GetDuration(key string) interface{} {
	val, _ := ctx.Get(key)
	return val
}

// GetStringSlice 获取字符串切片
func (ctx *Context) GetStringSlice(key string) []string {
	val, _ := ctx.Get(key)
	if s, ok := val.([]string); ok {
		return s
	}
	return nil
}

// GetStringMap 获取字符串map
func (ctx *Context) GetStringMap(key string) map[string]interface{} {
	val, _ := ctx.Get(key)
	if m, ok := val.(map[string]interface{}); ok {
		return m
	}
	return nil
}

// GetStringMapString 获取字符串string map
func (ctx *Context) GetStringMapString(key string) map[string]string {
	val, _ := ctx.Get(key)
	if m, ok := val.(map[string]string); ok {
		return m
	}
	return nil
}

// GetStringMapStringSlice 获取字符串string slice map
func (ctx *Context) GetStringMapStringSlice(key string) map[string][]string {
	val, _ := ctx.Get(key)
	if m, ok := val.(map[string][]string); ok {
		return m
	}
	return nil
}

// SetAccepted 设置接受的格式
func (ctx *Context) SetAccepted(formats ...string) {
	// Hertz需要特殊处理
}

// Negotiate 协商响应格式
func (ctx *Context) Negotiate(code int, config interface{}) {
	// Hertz需要特殊处理
}

// NegotiateFormat 协商格式
func (ctx *Context) NegotiateFormat(offered ...string) string {
	// Hertz需要特殊处理
	return ""
}

// GetQueryArray 获取query数组
func (ctx *Context) GetQueryArray(key string) ([]string, bool) {
	arr := ctx.QueryArray(key)
	return arr, len(arr) > 0
}

// GetQueryMap 获取query map
func (ctx *Context) GetQueryMap(key string) (map[string]string, bool) {
	m := ctx.QueryMap(key)
	return m, len(m) > 0
}

// GetPostFormArray 获取post form数组
func (ctx *Context) GetPostFormArray(key string) ([]string, bool) {
	arr := ctx.PostFormArray(key)
	return arr, len(arr) > 0
}

// GetPostFormMap 获取post form map
func (ctx *Context) GetPostFormMap(key string) (map[string]string, bool) {
	m := ctx.PostFormMap(key)
	return m, len(m) > 0
}

// AsciiJSON 返回ASCII JSON响应
func (ctx *Context) AsciiJSON(code int, obj interface{}) {
	ctx.c.JSON(code, obj)
}

// PureJSON 返回纯JSON响应
func (ctx *Context) PureJSON(code int, obj interface{}) {
	ctx.c.PureJSON(code, obj)
}

// IndentedJSON 返回格式化的JSON响应
func (ctx *Context) IndentedJSON(code int, obj interface{}) {
	ctx.c.IndentedJSON(code, obj)
}

// SecureJSON 返回安全的JSON响应
func (ctx *Context) SecureJSON(code int, obj interface{}) {
	ctx.c.JSON(code, obj)
}

// JSONP 返回JSONP响应
func (ctx *Context) JSONP(code int, obj interface{}) {
	ctx.c.JSON(code, obj)
}

// HTML 返回HTML响应
func (ctx *Context) HTML(code int, name string, obj interface{}) {
	ctx.c.HTML(code, name, obj)
}

// XML 返回XML响应
func (ctx *Context) XML(code int, obj interface{}) {
	ctx.c.XML(code, obj)
}

// YAML 返回YAML响应
func (ctx *Context) YAML(code int, obj interface{}) {
	// Hertz需要特殊处理
	ctx.c.JSON(code, obj)
}

// ProtoBuf 返回ProtoBuf响应
func (ctx *Context) ProtoBuf(code int, obj interface{}) {
	ctx.c.ProtoBuf(code, obj)
}

// Copy 复制context
func (ctx *Context) Copy() *Context {
	return &Context{
		ctx: ctx.ctx,
		c:   ctx.c.Copy(),
	}
}

// IsAborted 检查是否已中止
func (ctx *Context) IsAborted() bool {
	return ctx.c.IsAborted()
}
