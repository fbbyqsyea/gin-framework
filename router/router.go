package router

import (
	"net/http"

	"github.com/fbbyqsyea/gin-framework/docs"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 静态文件路由
type StaticRouters struct{}

func NewStaticRouters() *StaticRouters {
	return &StaticRouters{}
}

func (s *StaticRouters) Register(g *gin.Engine) {
}

// swagger路由
type SwaggerRouters struct{}

func NewSwaggerRouters() *SwaggerRouters {
	return &SwaggerRouters{}
}

func (s *SwaggerRouters) Register(g *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/"
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// api接口路由
type ApiRouters struct {
	user *UserRouter
	role *RoleRouter
}

func NewApiRouters() *ApiRouters {
	return &ApiRouters{
		user: NewUserRouter(),
		role: NewRoleRouter(),
	}
}

func (a *ApiRouters) Register(g *gin.Engine) {
	apig := g.Group("api")
	{
		a.user.Register(apig)
		a.role.Register(apig)
	}
}

// 全局路由
func InitRouoters(g *gin.Engine) {
	// ping 测试路由
	g.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})
	// 静态路由
	NewStaticRouters().Register(g)
	// swagger路由
	NewSwaggerRouters().Register(g)
	// api路由
	NewApiRouters().Register(g)
}
