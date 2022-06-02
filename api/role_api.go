package api

import (
	"github.com/fbbyqsyea/gin-framework/contexts"
	"github.com/fbbyqsyea/gin-framework/service"
	"github.com/gin-gonic/gin"
)

type RoleApi struct {
	*Api
	Svc *service.RoleService
}

func NewRoleApi() *RoleApi {
	svc := service.NewRoleService()
	return &RoleApi{
		Api: NewApi(svc),
		Svc: svc,
	}
}

// 新增角色信息
// @Summary 新增角色信息接口
// @Schemes http https
// @Description 用于新增角色信息
// @Tags 角色相关
// @Produce json
// @param object body contexts.RoleInsertRequest true "角色信息"
// @param Authorization header string true "授权令牌"
// @Response 200 {object} contexts.RESPONSE{data=uint} "新增的角色id"
// @Router /api/role/insert [post]
func (r *RoleApi) Insert(c *gin.Context) {
	r.Api.Insert(c, &contexts.RoleInsertRequest{})
}

// 更新角色信息
// @Summary 更新角色信息接口
// @Schemes http https
// @Description 用于更新角色信息
// @Tags 角色相关
// @Produce json
// @param object body contexts.RoleUpdateRequest true "角色信息"
// @param Authorization header string true "授权令牌"
// @Response 200 {object} contexts.RESPONSE{data=uint} "更新影响行数"
// @Router /api/role/update [post]
func (r *RoleApi) Update(c *gin.Context) {
	r.Api.Update(c, &contexts.RoleUpdateRequest{})
}

// 更新角色状态信息
// @Summary 更新角色状态信息接口
// @Schemes http https
// @Description 用于更新角色状态信息
// @Tags 角色相关
// @Produce json
// @param object body contexts.StatusRequest true "状态信息"
// @param Authorization header string true "授权令牌"
// @Response 200 {object} contexts.RESPONSE{data=uint} "更新影响行数"
// @Router /api/role/status [post]
func (r *RoleApi) Status(c *gin.Context) {
	r.Api.Status(c, &contexts.StatusRequest{})
}

// 删除角色
// @Summary 删除角色接口
// @Schemes http https
// @Description 用于删除角色
// @Tags 角色相关
// @Produce json
// @param object body contexts.RemoveRequest true "角色信息"
// @param Authorization header string true "授权令牌"
// @Response 200 {object} contexts.RESPONSE{data=uint} "删除影响行数"
// @Router /api/role/remove [post]
func (r *RoleApi) Remove(c *gin.Context) {
	r.Api.Remove(c, &contexts.RemoveRequest{})
}

// 批量删除角色
// @Summary 批量删除角色接口
// @Schemes http https
// @Description 用于批量删除角色
// @Tags 角色相关
// @Produce json
// @param object body contexts.RemovesRequest true "角色信息"
// @param Authorization header string true "授权令牌"
// @Response 200 {object} contexts.RESPONSE{data=uint} "删除影响行数"
// @Router /api/role/removes [post]
func (r *RoleApi) Removes(c *gin.Context) {
	r.Api.Removes(c, &contexts.RemovesRequest{})
}

// 获取角色信息
// @Summary 获取角色信息接口
// @Schemes http https
// @Description 用于获取角色信息
// @Tags 角色相关
// @Produce json
// @param id query uint true "角色ID"
// @param Authorization header string true "授权令牌"
// @Response 200 {object} contexts.RESPONSE{data=contexts.RoleData}
// @Router /api/role/get [get]
func (r *RoleApi) Get(c *gin.Context) {
	r.Api.Get(c, &contexts.GetRequest{}, &contexts.RoleData{})
}

// 获取角色列表
// @Summary 获取角色列表接口
// @Schemes http https
// @Description 用于获取角色列表
// @Tags 角色相关
// @Produce json
// @param id query uint false "角色id"
// @param status query int false "状态 1:启用 2:禁用 默认:1"
// @param page query uint64 false "页数 默认:1"
// @param limit query uint64 false "页数量 默认:20"
// @param Authorization header string true "授权令牌"
// @Response 200 {object} contexts.RESPONSE{data=[]contexts.RoleData{}}
// @Router /api/role/list [get]
func (r *RoleApi) List(c *gin.Context) {
	r.Api.List(c, &contexts.RoleListRequest{}, &[]contexts.RoleData{})
}
