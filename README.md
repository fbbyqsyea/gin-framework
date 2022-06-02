# gin-framework

<img src="https://img.shields.io/badge/license-MIT-green" alt="license" /> <img src="https://img.shields.io/badge/version-1.0.0-blue" alt="version" /> <img src="https://img.shields.io/badge/golang-1.18-orange" alt="golang" />

### 项目介绍
这是一个基于<a href="https://gin-gonic.com/">gin</a>的web api开发框架，封装了<b>Config</b>、<b>Logger</b>、<b>Middleware</b>、<b>Api</b>、<b>Service</b>、<b>Model</b>、<b>Router</b>、<b>Utils</b>等模块。而且框架的ASM(api,service,model)模块封装了一套简易的CURD逻辑，只需要简单的定义几个struct和variable就可以实现一个CURD。对于后台系统开发尤为方便。项目目前处于维护中，欢迎PR。

### 项目结构
```tree
gin-framework:
            |- api                                            # api入口
            |   |- api.go
            |   |- user_api.go
            |- config                                         # 配置信息
            |   |- config.go
            |- contexts                                       # 请求响应等上下文定义
            |   |- contexts.go
            |   |- user_contexts.go
            |- docs                                           # swagger 生成的文档 不要编辑
            |- global                                         # 全局变量定义 config、logger、db、redis 等全局使用的变量统一在此定义
            |   |- global.go
            |- initialize                                     # 框架初始化模块
            |   |- config.go
            |   |- db.go
            |   |- initialize.go
            |   |- logger.go
            |   |- redis.go
            |   |- router.go
            |- log                                            # 日志文件
            |- middleware                                     # 中间件
            |- model                                          # 数据库操作Model
            |   |- model.go
            |   |- user_model.go
            |- router                                         # 路由定义
            |   |- router.go
            |   |- user_api_router.go
            |- service                                        # 业务逻辑处理 
            |   |- service.go
            |   |- user_service.go
            |- sql                                            # 数据表结构  ddl & dml
            |- tmp                                            # fresh 热重启生成的临时文件目录 不要编辑
            |- utils                                          # 工具包 公共函数模块
            |- config.yaml                                    # 配置信息yaml文件
            |- fresh.conf                                     # fresh 热重启配置文件
            |- go.mod                                         # go mod
            |- go.sum                                         # go sum
            |- main.go                                        # 程序的统一入口
```

### 开始使用

##### 拉取项目
```shell
$ git clone https://github.com/fbbyqsyea/gin-framework.git
```

##### 依赖安装
```shell
$ cd gin-framework && go mod tidy
```

##### 启动项目(研发环境)
```
$ fresh -c fresh.conf
```

### 开发示例(CURD)
我们以开发一个角色管理的后台api为例来介绍一下框架的开发流程。

##### 新建角色表
    在sql下新建role.sql
```sql
use `operation`;

CREATE TABLE `tb_operation_role` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `role_name` varchar(50) NOT NULL DEFAULT '' COMMENT '角色名称',
  `parent_id` int(11) unsigned not null default 0 comment '父级角色id',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '状态  1:启用  2:禁用',
  `is_delete` tinyint(3) unsigned NOT NULL DEFAULT '2' COMMENT '是否删除  1:是  2:否',
  `insert_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '写入时间',
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最近一次更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT="角色表";
```
##### model开发
    在model目录下新建role_model.go。

```go
package model

// 定义角色 model  角色model集成model struct的方法
type RoleModel struct {
	*Model
}

// 实例化RoleModel
func NewRoleModel() *RoleModel {
	return &RoleModel{
		NewModel("tb_operation_role"), // tb_operation_role 数据表
	}
}

// 如果有新的特殊的角色数据库操作 可在此文件定义

```
##### service开发
    在service目录下新建role_service.go。

```go
package service

import (
	"github.com/fbbyqsyea/gin-framework/model"
)

// 定义角色 service 处理角色api的逻辑 继承了service struct的方法 
type RoleService struct {
	*Service
	Mdl *model.RoleModel
}

// 实例化service 内部包含Role 用于操作数据库
func NewRoleService() *RoleService {
	mdl := model.NewRoleModel()
	return &RoleService{
		Service: NewService(mdl),
		Mdl:     mdl,
	}
}

// 如果有新的特殊的角色方法 可在此文件定义

```
##### api模块开发
    在api目录下新建role_api.go文件，用来管理角色的api。
    在contexts目录下新建role_contexts.go，用来管理role模块的上下文结构体。
```go
// role_api.go
package api

import (
	"github.com/fbbyqsyea/gin-framework/contexts"
	"github.com/fbbyqsyea/gin-framework/service"
	"github.com/gin-gonic/gin"
)

// 角色api 
type RoleApi struct {
	*Api
	Svc *service.RoleService
}

// 实例化角色api 内嵌role service 逻辑实现
func NewRoleApi() *RoleApi {
	svc := service.NewRoleService()
	return &RoleApi{
		Api: NewApi(svc),
		Svc: svc,
	}
}
```
###### insert
    新增角色api

```go
// role_contexts.go 新增insert struct  
// json Tag用来接收数据  insert Tag是db需要写入的信息
type RoleInsertRequest struct {
	RoleName string `json:"role_name" insert:"role_name"` // 角色名称
	ParentId uint   `json:"parent_id" insert:"parent_id"` // 父级角色ID
}
```
```go
// role_api.go 文件新增Insert 方法 调用继承的api的insert方法完成写入

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
```
###### update
    更新角色api
```go
// role_contexts.go 新增update struct  
// json Tag用来接收数据  update Tag是db需要更新的信息
type RoleUpdateRequest struct {
	Id       uint   `json:"id" where:"id"`
	RoleName string `json:"role_name" update:"role_name"` // 角色名称
	ParentId string `json:"parent_id" update:"parent_id"` // 父级角色ID
}
```
```go
// role_api.go 文件新增update 方法 调用继承的api的update方法完成写入

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
```

###### status
    修改角色状态api
```go
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
```

###### remove
    删除角色api
```go
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
```

###### removes
    批量删除角色api
```go
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
```

###### get
    获取角色信息api
```go
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
```

###### list
    获取角色列表api
```go
// list request struct 用来接收列表筛选条件
// form tag用来接收get请求参数 where Tag用来限定查询条件 order 用来限定排序 
// page,limit用来限制limit offset批量取数
type RoleListRequest struct {
	Id       uint `order:"id desc"`
	Status   int  `form:"status" where:"status"`  // 状态
	IsDelete int  `where:"is_delete" default:"2"` // 是否删除
	PAGEANDLIMIT
}
```
```go
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

```

##### router开发
    开发角色模块路由映射，在router模块下新建role_api_router.go文件
```go
// role_api_router.go
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
```
```go
// 在router.go中注册RoleRouter

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
```

##### 生成新的swagger文档
```sh
# 文档地址:http://localhost:8888/swagger/index.html#/
$ swag init
```

##### fresh 热更新 
    研发环境为了方便调试代码，我们使用fresh插件来进行热更新，当文件有改动的时候，自动build and run
```sh
$ fresh
# 或者
$ fresh -c fresh.conf
```
