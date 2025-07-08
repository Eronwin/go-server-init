package router

import (
	"github.com/gin-gonic/gin"
	v1 "go-server-init/internal/api/v1"
	"go-server-init/internal/config"
	"go-server-init/internal/middleware"
	"go-server-init/internal/repository"
	"go-server-init/internal/service"
	"go.uber.org/zap"
	"net/http"
)

func SetupRouter(cfg *config.Config, logger *zap.SugaredLogger) *gin.Engine {
	r := gin.New()

	// 注册全局中间件
	r.Use(
		gin.Recovery(),
		middleware.ZapLoggerMiddleware(logger),
	)

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// 业务依赖
	repo := repository.New()
	svc := service.New(repo)

	// API v1

	v1Group := r.Group("/api/v1")
	{
		v1Group.GET("/ping", v1.NewPingHandler(svc).Ping)
	}
	return r
}
