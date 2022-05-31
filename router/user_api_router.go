package router

import (
	"github.com/fbbyqsyea/gin-framework/api"
	"github.com/fbbyqsyea/gin-framework/middleware"
	"github.com/gin-gonic/gin"
)

// 用户api router
type UserRouter struct{}

func NewUserRouter() *UserRouter {
	return &UserRouter{}
}

func (u *UserRouter) Register(g *gin.RouterGroup) {
	// 用户路由
	userg := g.Group("user")
	{
		// 用户api模块
		userApi := api.NewUserApi()
		// 需要授权的路由
		auth := userg.Group("")
		{
			auth.Use(middleware.Authorization())
			auth.GET("get", userApi.Get)
			auth.GET("list", userApi.List)
			auth.POST("insert", userApi.Insert)
			auth.POST("update", userApi.Update)
			auth.POST("status", userApi.Status)
			auth.POST("remove", userApi.Remove)
			auth.POST("removes", userApi.Removes)
		}
		// 不需要授权的路由
		noAuth := userg.Group("")
		{
			noAuth.POST("login", userApi.Login)
		}
	}
}
