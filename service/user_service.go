package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/fbbyqsyea/gin-framework/contexts"
	"github.com/fbbyqsyea/gin-framework/global"
	"github.com/fbbyqsyea/gin-framework/model"
	"github.com/fbbyqsyea/gin-framework/utils"
)

type UserService struct {
	*Service
	Mdl *model.UserModel
}

type AuthorizationClaims struct {
	jwt.StandardClaims
	Id uint
}

func NewUserService() *UserService {
	mdl := model.NewUserModel()
	return &UserService{
		Service: NewService(mdl),
		Mdl:     mdl,
	}
}

// 登录操作
func (us *UserService) Login(req *contexts.UserLoginRequest) *contexts.UserLoginResponse {
	var resp contexts.UserLoginResponse
	var userData contexts.UserData
	// 获取用户信息
	err := us.Mdl.Get(req, &userData, false)
	fmt.Println(utils.MD5(req.Password+userData.Salt), "122222222")
	if errors.Is(err, sql.ErrNoRows) {
		resp.STATE(contexts.ERR_USER_NOT_EXISTS)
	} else if err != nil { // 其他异常错误处理
		resp.STATE(contexts.ERR_SYSTEM)
	} else if userData.Status == 2 || userData.IsDelete == 1 { // 状态校验
		resp.STATE(contexts.ERR_USER_FORBIDDEN)
	} else if utils.MD5(req.Password+userData.Salt) != userData.Password { // 校验用户密码是否正确
		resp.STATE(contexts.ERR_USER_PASSWORD)
	} else {
		// 生成授权令牌
		authorization, err := us.CreateAuthorization(userData.ID)
		if err != nil {
			global.LOGGER.Error(fmt.Sprintf("fatal create token: [%s]", err))
			resp.STATE(contexts.ERR_SYSTEM)
		} else {
			// 设置登录标识
			us.SetLogin(userData.ID, global.CONFIG.Jwt.ExpireTime)
			resp.STATE(contexts.SUC_LOGIN)
			// 组装返回数据
			resp.Data.Authorization = authorization
		}
	}
	return &resp
}

// 生成授权令牌
func (us *UserService) CreateAuthorization(id uint) (s string, e error) {
	unix := time.Now().Unix()
	// 生成登录token
	c := AuthorizationClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: unix + global.CONFIG.Jwt.ExpireTime,
			IssuedAt:  unix,
			Issuer:    global.CONFIG.Jwt.Issuer,
			NotBefore: unix - 1,
			Subject:   global.CONFIG.Jwt.Subject,
		},
		Id: id,
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return t.SignedString([]byte(global.CONFIG.Jwt.SecretKey))
}

// 解析授权令牌
func (us *UserService) ParseAuthorization(s string) (*AuthorizationClaims, *contexts.State) {
	token, err := jwt.ParseWithClaims(s, &AuthorizationClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.CONFIG.Jwt.SecretKey), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok && ve.Errors&jwt.ValidationErrorExpired != 0 {
			return nil, contexts.ERR_USER_LOGIN_OUT_TIME
		}
	} else {
		// 将token中的claims信息解析出来并断言成用户自定义的有效载荷结构
		if claims, ok := token.Claims.(*AuthorizationClaims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, contexts.ERR_SYSTEM
}

// 写入用户信息
func (us *UserService) Insert(req any) *contexts.RESPONSE {
	var resp contexts.RESPONSE
	req1, _ := req.(*contexts.UserInsertRequest)
	// 密码和确认密码校验
	if req1.Password != req1.ComfirmPassword {
		resp.STATE(contexts.ERR_USER_PASSWORD_INCONFORMITY)
		return &resp
	}
	// 随机生成盐
	req1.Salt = utils.RandString(16)
	// 重新加密密码
	req1.Password = utils.MD5(req1.Password + req1.Salt)
	return us.Service.Insert(req1)
}

// 更新用户信息
func (us *UserService) Update(req any) *contexts.RESPONSE {
	var resp contexts.RESPONSE
	req1, _ := req.(*contexts.UserUpdateRequest)
	if req1.Password != "" {
		// 密码和确认密码校验
		if req1.Password != req1.ComfirmPassword {
			resp.STATE(contexts.ERR_USER_PASSWORD_INCONFORMITY)
			return &resp
		}
		var u contexts.UserData
		// 用户信息获取
		err := us.Mdl.Get(req, &u, true)
		if errors.Is(err, sql.ErrNoRows) {
			resp.STATE(contexts.ERR_USER_NOT_EXISTS)
			return &resp
		} else if err != nil {
			resp.STATE(contexts.ERR_SYSTEM)
			return &resp
		}
		// 重新加密密码
		req1.Password = utils.MD5(req1.Password + u.Salt)
	}
	return us.Service.Update(req1)
}

var ctx = context.Background()

// 用户登录标识
const USER_LOGIN_TPL = "user:login:%d"

// 设置登录标识(用户登录后设置)
func (us *UserService) SetLogin(id uint, expireTime int64) {
	global.REDIS.Set(ctx, fmt.Sprintf(USER_LOGIN_TPL, id), 1, time.Duration(expireTime)*time.Second)
}

// 获取登录标识
func (us *UserService) GetLogin(id uint) string {
	result, _ := global.REDIS.Get(ctx, fmt.Sprintf(USER_LOGIN_TPL, id)).Result()
	return result
}

// 删除登录标识(账号被禁用或者被删除后删除登录标识让用户被动下线)
func (us *UserService) DelLogin(id uint) {
	global.REDIS.Del(ctx, fmt.Sprintf(USER_LOGIN_TPL, id))
}
