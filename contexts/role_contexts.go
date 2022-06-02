package contexts

type RoleData struct {
	*DATA
	RoleName string `json:"role_name" db:"role_name"` // 角色名称
	ParentId uint   `json:"parent_id" db:"parent_id"` // 父级角色ID
}

type RoleInsertRequest struct {
	RoleName string `json:"role_name" insert:"role_name"` // 角色名称
	ParentId uint   `json:"parent_id" insert:"parent_id"` // 父级角色ID
}

type RoleUpdateRequest struct {
	Id       uint   `json:"id" where:"id"`
	RoleName string `json:"role_name" update:"role_name"` // 角色名称
	ParentId uint   `json:"parent_id" update:"parent_id"` // 父级角色ID
}

type RoleListRequest struct {
	Id       uint `order:"id desc"`
	Status   int  `form:"status" where:"status"`  // 状态
	IsDelete int  `where:"is_delete" default:"2"` // 是否删除
	PAGEANDLIMIT
}
