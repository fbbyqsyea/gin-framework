package initialize

import (
	"github.com/fbbyqsyea/gin-framework/router"
	"github.com/gin-gonic/gin"
)

// 初始化路由 router模块管理
func initRouters(g *gin.Engine) {
	router.InitRouoters(g)
}
