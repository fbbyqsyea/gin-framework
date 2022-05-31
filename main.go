package main

import (
	"github.com/fbbyqsyea/gin-framework/global"
	"github.com/fbbyqsyea/gin-framework/initialize"
	"github.com/fbbyqsyea/gin-framework/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	// 运行环境
	gin.SetMode(global.CONFIG.System.Mode)
	// 实例化gin框架
	g := gin.New()
	// 加载全局中间件
	g.Use(middleware.GinLogger(), gin.Recovery())
	// 全局初始化
	initialize.Init(g)
	// 运行服务
	g.Run(global.CONFIG.System.Attr())
}
