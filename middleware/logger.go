package middleware

import (
	"time"

	"github.com/fbbyqsyea/gin-framework/global"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 全局日志中间件
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 生产环境记录访问日志
		if global.CONFIG.System.Mode == gin.ReleaseMode {
			start := time.Now()
			path := c.Request.URL.Path
			query := c.Request.URL.RawQuery
			global.LOGGER.Info(path,
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("ip", c.ClientIP()),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
				zap.Duration("cost", time.Since(start)),
			)
			c.Next()
		}
	}
}
