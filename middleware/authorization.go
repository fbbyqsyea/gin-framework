package middleware

import (
	"fmt"
	"net/http"

	"github.com/fbbyqsyea/gin-framework/contexts"
	"github.com/fbbyqsyea/gin-framework/global"
	"github.com/fbbyqsyea/gin-framework/service"

	"github.com/gin-gonic/gin"
)

// 授权校验中间件
func Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var resp contexts.RESPONSE
		// 从header中获取Authorization
		authorization := ctx.Request.Header.Get("Authorization")
		if authorization == "" {
			global.LOGGER.Error("fatal error no authorization.")
			resp.STATE(contexts.ERR_USER_NO_LOGIN)
		} else {
			us := service.NewUserService()
			// 解析token
			authorizationClaims, err := us.ParseAuthorization(authorization)
			if err != nil {
				global.LOGGER.Error(fmt.Sprintf("fatal error parse authorization [%s] - [%s]", authorization, err))
				resp.STATE(err)
			} else {
				// 登录态校验
				if us.GetLogin(authorizationClaims.Id) != "" {
					ctx.Set("Id", authorizationClaims.Id)
					ctx.Next()
					return
				}
				global.LOGGER.Error(fmt.Sprintf("fatal error lost login status [%s]", authorization))
				resp.STATE(contexts.ERR_USER_LOGIN_OUT_TIME)
			}
		}
		ctx.JSON(http.StatusNetworkAuthenticationRequired, resp)
		ctx.Abort()
	}
}
