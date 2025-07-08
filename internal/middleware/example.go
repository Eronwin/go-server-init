package middleware

import (
	"github.com/gin-gonic/gin"
)

func ExampleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 拦截 + 日志 + 鉴权
		c.Next()
	}
}
