package initialize

import "github.com/gin-gonic/gin"

// 全局初始化
func Init(g *gin.Engine) {
	// 初始化配置
	initConfig()
	// 初始化日志
	initLogger()
	// 初始化数据库
	initDB()
	// 初始化cache redis
	initRedis()
	// 初始化路由
	initRouters(g)
}
