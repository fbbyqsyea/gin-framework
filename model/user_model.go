package model

type UserModel struct {
	*Model
}

// 数据表
var tableName = "tb_operation_user"

func NewUserModel() *UserModel {
	return &UserModel{
		NewModel(tableName),
	}
}
