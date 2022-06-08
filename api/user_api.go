package api

import (
	"github.com/fbbyqsyea/gin-framework/contexts"
	"github.com/fbbyqsyea/gin-framework/service"
	"github.com/gin-gonic/gin"
)

type UserApi struct {
	*Api
	Svc *service.UserService
}

func NewUserApi() *UserApi {
	svc := service.NewUserService()
	return &UserApi{
		Api: NewApi(svc),
		Svc: svc,
	}
}

// 登录
// @Summary 登录接口
// @Schemes http https
// @Description 用于登录admin
// @Tags 用户相关
// @Accept json
// @Produce json
// @param object body contexts.UserLoginRequest true "登录参数"
// @Response 200 {object} contexts.UserLoginResponse
// @Router /api/user/login [post]
func (u *UserApi) Login(c *gin.Context) {
	var req contexts.UserLoginRequest
	if u.ShouldBindJSON(c, &req) {
		u.JSON(c, u.Svc.Login(&req))
	}
}

// 获取用户信息
// @Summary 获取用户信息接口
// @Schemes http https
// @Description 用于获取用户信息
// @Tags 用户相关
// @Produce json
// @param id query uint true "用户ID"
// @param Authorization header string true "授权令牌"
// @Response 200 {object} contexts.RESPONSE{data=contexts.UserData}
// @Router /api/user/get [get]
func (u *UserApi) Get(c *gin.Context) {
	u.Api.Get(c, &contexts.GetRequest{}, &contexts.UserData{}, false)
}

// 获取用户列表
// @Summary 获取用户列表接口
// @Schemes http https
// @Description 用于获取用户列表
// @Tags 用户相关
// @Produce json
// @param account query string false "账户"
// @param status query int false "状态 1:启用 2:禁用 默认:1"
// @param page query uint64 false "页数 默认:1"
// @param limit query uint64 false "页数量 默认:20"
// @param Authorization header string true "授权令牌"
// @Response 200 {object} contexts.RESPONSE{data=[]contexts.UserData{}}
// @Router /api/user/list [get]
func (u *UserApi) List(c *gin.Context) {
	u.Api.List(c, &contexts.UserListRequest{}, &[]contexts.UserData{}, false)
}

// 新增用户信息
// @Summary 新增用户信息接口
// @Schemes http https
// @Description 用于新增用户信息
// @Tags 用户相关
// @Produce json
// @param object body contexts.UserInsertRequest true "用户信息"
// @param Authorization header string true "授权令牌"
// @Response 200 {object} contexts.RESPONSE{data=uint} "用户id"
// @Router /api/user/insert [post]
func (u *UserApi) Insert(c *gin.Context) {
	u.Api.Insert(c, &contexts.UserInsertRequest{})
}

// 更新用户信息
// @Summary 更新用户信息接口
// @Schemes http https
// @Description 用于更新用户信息
// @Tags 用户相关
// @Produce json
// @param object body contexts.UserUpdateRequest true "用户信息"
// @param Authorization header string true "授权令牌"
// @Response 200 {object} contexts.RESPONSE{data=uint} "用户信息更新是否成功"
// @Router /api/user/update [post]
func (u *UserApi) Update(c *gin.Context) {
	u.Api.Update(c, &contexts.UserUpdateRequest{})
}

// 更新用户状态
// @Summary 更新用户状态接口
// @Schemes http https
// @Description 用于更新用户状态
// @Tags 用户相关
// @Produce json
// @param object body contexts.StatusRequest true "用户状态信息"
// @param Authorization header string true "授权令牌"
// @Response 200 {object} contexts.RESPONSE{data=uint} "状态变更是否成功 1:是 0:否"
// @Router /api/user/status [post]
func (u *UserApi) Status(c *gin.Context) {
	var req contexts.StatusRequest
	if u.ShouldBindJSON(c, &req) {
		loginId, exists := c.Get("Id")
		if exists && req.Id == loginId {
			resp := &contexts.RESPONSE{}
			resp.STATE(contexts.ERR_USER_CAN_NOT_FORBIDDEN_LOGIN)
			u.JSON(c, resp)
		} else {
			u.Svc.DelLogin(req.Id)
			u.JSON(c, u.Svc.Service.Update(&req))
		}
	}
}

// 删除用户
// @Summary 删除用户接口
// @Schemes http https
// @Description 用于删除用户
// @Tags 用户相关
// @Produce json
// @param object body contexts.RemoveRequest true "用户id信息"
// @param Authorization header string true "授权令牌"
// @Response 200 {object} contexts.RESPONSE{data=uint} "删除用户数"
// @Router /api/user/remove [post]
func (u *UserApi) Remove(c *gin.Context) {
	var req contexts.RemoveRequest
	if u.ShouldBindJSON(c, &req) {
		loginId, exists := c.Get("Id")
		if exists && req.Id == loginId {
			resp := &contexts.RESPONSE{}
			resp.STATE(contexts.ERR_USER_CAN_NOT_DELETE_LOGIN)
			u.JSON(c, resp)
		} else {
			u.Svc.DelLogin(req.Id)
			u.JSON(c, u.Svc.Remove(&req))
		}
	}
}

// 批量删除用户
// @Summary 批量删除用户接口
// @Schemes http https
// @Description 用于批量删除用户
// @Tags 用户相关
// @Produce json
// @param object body contexts.RemovesRequest true "用户ids信息"
// @param Authorization header string true "授权令牌"
// @Response 200 {object} contexts.RESPONSE{data=uint} "批量删除用户数"
// @Router /api/user/removes [post]
func (u *UserApi) Removes(c *gin.Context) {
	var req contexts.RemovesRequest
	var canRemove = true
	loginId, exists := c.Get("Id")
	if u.ShouldBindJSON(c, &req) {
		for _, Id := range req.Ids {
			if exists && Id == loginId {
				canRemove = false
				break
			}
		}
		if !canRemove {
			resp := &contexts.RESPONSE{}
			resp.STATE(contexts.ERR_USER_CAN_NOT_DELETE_LOGIN)
			u.JSON(c, resp)
		} else {
			for _, id := range req.Ids {
				u.Svc.DelLogin(id)
			}
			u.JSON(c, u.Svc.Remove(&req))
		}
	}
}
