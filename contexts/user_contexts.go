package contexts

// 定义用户相关状态码
var (
	ERR_USER_NOT_EXISTS              = &State{20001, "该用户不存在。"}
	ERR_USER_FORBIDDEN               = &State{20002, "该用户被禁止登录系统。"}
	ERR_USER_PASSWORD                = &State{20003, "密码错误。"}
	ERR_USER_PASSWORD_INCONFORMITY   = &State{20003, "密码和确认密码不一致。"}
	ERR_USER_CAN_NOT_DELETE_LOGIN    = &State{20004, "不能删除当前登录用户。"}
	ERR_USER_NO_LOGIN                = &State{20005, "未登录，请先登录。"}
	ERR_USER_SYSTEM_UNSAFE           = &State{20006, "系统异常，请重新登录。"}
	ERR_USER_LOGIN_OUT_TIME          = &State{20007, "登录失效，请重新登录。"}
	ERR_USER_CAN_NOT_FORBIDDEN_LOGIN = &State{20004, "不能禁用当前登录用户。"}
)

// 用户表数据
type UserData struct {
	*DATA
	Account  string `json:"account" db:"account"`   // 账户
	UserName string `json:"username" db:"username"` // 用户名
	Password string `json:"-" db:"password"`
	Salt     string `json:"-" db:"salt"`
}

type UserLoginRequest struct {
	Account  string `json:"account" where:"account,"` // 账号
	Password string `json:"password"`                 // 密码
}

type UserLoginResponse struct {
	RESPONSE
	Data struct {
		Authorization string `json:"authorization"` // 登录授权令牌
	} `json:"data"` // 数据
}

type UserListRequest struct {
	Id       uint   `order:"id desc"`
	Account  string `form:"account" where:"account,like"` // 账户
	Status   int    `form:"status" where:"status"`        // 状态
	IsDelete int    `where:"is_delete" default:"2"`       // 是否删除
	PAGEANDLIMIT
}

type UserInsertRequest struct {
	Account         string `json:"account" insert:"account"`           // 账户
	UserName        string `json:"username" insert:"username"`         // 用户名
	Password        string `json:"password" insert:"password"`         // 密码
	ComfirmPassword string `json:"confirmPassword"`                    // 确认密码
	Salt            string `insert:"salt"`                             // 盐
	Status          int    `json:"status" default:"2" insert:"status"` // 状态 1:启用 2:禁用
	IsDelete        int    `default:"2" insert:"is_delete"`            // 是否删除 1:是 0:否
}

type UserUpdateRequest struct {
	Id              uint   `json:"id" where:"id"`              // 用户ID
	Account         string `json:"account" update:"account"`   // 账户
	UserName        string `json:"username" update:"username"` // 用户名
	Password        string `json:"password" update:"password"` // 密码
	ComfirmPassword string `json:"confirmPassword"`            // 确认密码
	IsDelete        int    `default:"2" where:"is_delete"`     // 是否删除 1:是 0:否
}
