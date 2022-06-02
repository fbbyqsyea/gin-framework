package model

type RoleModel struct {
	*Model
}

func NewRoleModel() *RoleModel {
	return &RoleModel{
		NewModel("tb_operation_role"),
	}
}
