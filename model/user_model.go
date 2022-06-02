package model

type UserModel struct {
	*Model
}

func NewUserModel() *UserModel {
	return &UserModel{
		NewModel("tb_operation_user"),
	}
}
