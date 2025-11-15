package core

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cloudwego/hertz/pkg/app/server"
	"go.uber.org/zap"
)

type serverInterface interface {
	Spin() error
	Shutdown(context.Context) error
}

// initServer 启动服务并实现优雅关闭
func initServer(address string, router *server.Hertz, readTimeout, writeTimeout time.Duration) {
	// Hertz服务器已经在Routers()中创建并配置好了，这里直接启动

	// 在goroutine中启动服务
	go func() {
		router.Spin() // Hertz的Spin()是阻塞调用
	}()

	// 等待中断信号以优雅地关闭服务器
	quit := make(chan os.Signal, 1)
	// kill (无参数) 默认发送 syscall.SIGTERM
	// kill -2 发送 syscall.SIGINT
	// kill -9 发送 syscall.SIGKILL，但是无法被捕获，所以不需要添加
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Info("关闭WEB服务...")

	// 设置5秒的超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := router.Shutdown(ctx); err != nil {
		zap.L().Fatal("WEB服务关闭异常", zap.Error(err))
	}

	zap.L().Info("WEB服务已关闭")
}
