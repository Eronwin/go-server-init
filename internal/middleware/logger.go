package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func ZapLoggerMiddleware(logger *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		c.Next()
		latency := time.Since(start)
		status := c.Writer.Status()

		logger.Info("request log",
			"method", c.Request.Method,
			"path", path,
			"status", status,
			"latency", latency,
			"clientIP", c.ClientIP(),
		)
	}
}
