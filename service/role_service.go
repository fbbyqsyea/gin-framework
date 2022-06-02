package service

import (
	"github.com/fbbyqsyea/gin-framework/model"
)

type RoleService struct {
	*Service
	Mdl *model.RoleModel
}

func NewRoleService() *RoleService {
	mdl := model.NewRoleModel()
	return &RoleService{
		Service: NewService(mdl),
		Mdl:     mdl,
	}
}
