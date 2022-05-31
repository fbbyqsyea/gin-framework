package contexts

import (
	"fmt"
	"time"
)

// state context
type State struct {
	Code int
	Msg  string
}

func (s *State) Error() string {
	return fmt.Sprintf("err code:[%d] - err msg:[%s]", s.Code, s.Msg)
}

var (
	SUC_LOGIN  = &State{0, "登录成功"}
	ERR_PARAMS = &State{10001, "参数错误"}
	ERR_SYSTEM = &State{10002, "系统异常，请稍后重试。"}
	// common
	ERR_NO_DATA = &State{30001, "没有数据"}
	ERR_INSERT  = &State{30002, "新增失败"}
	ERR_UPDATE  = &State{30003, "更新失败"}
	ERR_DELETE  = &State{30004, "删除失败"}
)

// response context
type RESPONSE struct {
	Code int         `json:"code" default:"0"`      // 状态码 0:成功 非0:失败 失败可直接提示msg信息
	Msg  string      `json:"msg" default:"success"` // 状态信息
	Data interface{} `json:"data"`                  // 数据 成功返回指定数据 失败返回nil
}

func NewRESPONSE() *RESPONSE {
	return &RESPONSE{}
}

func (resp *RESPONSE) STATE(s *State) *RESPONSE {
	resp.Code = s.Code
	resp.Msg = s.Msg
	return resp
}

func (resp *RESPONSE) DATA(data interface{}) {
	resp.Data = data
}

type RESPONSEWITHCOUNT struct {
	RESPONSE
	Count int `json:"count"` // 数量
}

func (resp *RESPONSEWITHCOUNT) COUNT(count int) {
	resp.Count = count
}

// base context
type DATA struct {
	ID       uint      `json:"id" db:"id"`               // id
	Status   int       `json:"status" db:"status"`       // 状态 1:启用 2:禁用
	IsDelete int       `json:"-" db:"is_delete"`         // 是否删除  1:是 2:否
	InsertAt time.Time `json:"insert_at" db:"insert_at"` // 入库时间
	UpdateAt time.Time `json:"update_at" db:"update_at"` // 更新时间
}

type PAGEANDLIMIT struct {
	Page  uint64 `json:"page" form:"page" default:"1" page:"page"`      // 页数
	Limit uint64 `json:"limit" form:"limit" default:"20" limit:"limit"` // 页数量
}

type COUNT struct {
	Count int `json:"count" db:"count"` // 数量
}

type StatusRequest struct {
	Id     uint `json:"id" where:"id"`          // 用户ID
	Status int  `json:"status" update:"status"` // 状态 1:启用 2:禁用
}

type RemoveRequest struct {
	Id       uint `json:"id" where:"id"`                    // 用户ID
	IsDelete int  `json:"-" default:"1" update:"is_delete"` // 是否删除 1:是 0:否
}

type RemovesRequest struct {
	Ids      []uint `json:"ids" where:"id"`                   // 用户IDS
	IsDelete int    `json:"-" default:"1" update:"is_delete"` // 是否删除 1:是 0:否
}
