package router

import (
	"github.com/fbbyqsyea/gin-framework/api"
	"github.com/fbbyqsyea/gin-framework/middleware"
	"github.com/gin-gonic/gin"
)

type RoleRouter struct{}

func NewRoleRouter() *RoleRouter {
	return &RoleRouter{}
}

// 注册角色模块路由
func (r *RoleRouter) Register(g *gin.RouterGroup) {
	// 角色路由分组
	roleg := g.Group("role")
	{
		roleApi := api.NewRoleApi()
		// 角色路由使用后台用户权限校验中间件
		roleg.Use(middleware.Authorization())

		roleg.POST("insert", roleApi.Insert)
		roleg.POST("update", roleApi.Update)
		roleg.POST("status", roleApi.Status)
		roleg.POST("remove", roleApi.Remove)
		roleg.POST("removes", roleApi.Removes)
		roleg.GET("get", roleApi.Get)
		roleg.GET("list", roleApi.List)

	}
}
