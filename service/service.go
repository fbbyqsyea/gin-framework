package service

import (
	"database/sql"
	"errors"

	"github.com/fbbyqsyea/gin-framework/contexts"
)

type Model interface {
	Get(req, resp interface{}, isMaster bool) error
	Count(req, resp interface{}, isMaster bool) error
	Select(req, resp interface{}, isMaster bool) error
	Insert(req interface{}) (int64, error)
	Update(req interface{}) (int64, error)
}

type Service struct {
	Mdl Model
}

func NewService(mdl Model) *Service {
	return &Service{
		Mdl: mdl,
	}
}

func (s *Service) Get(req, data any, isMaster bool) *contexts.RESPONSE {
	var resp contexts.RESPONSE
	err := s.Mdl.Get(req, data, isMaster)
	if errors.Is(err, sql.ErrNoRows) {
		resp.STATE(contexts.ERR_USER_NOT_EXISTS)
	} else if err != nil {
		resp.STATE(contexts.ERR_SYSTEM)
	} else {
		resp.DATA(data)
	}
	return &resp
}

func (s *Service) List(req, data any, isMaster bool) *contexts.RESPONSEWITHCOUNT {
	var resp contexts.RESPONSEWITHCOUNT
	var count contexts.COUNT
	err := s.Mdl.Count(req, &count, isMaster)
	if err != nil {
		resp.STATE(contexts.ERR_SYSTEM)
	} else {
		resp.COUNT(count.Count)
		err := s.Mdl.Select(req, data, isMaster)
		if err != nil {
			resp.STATE(contexts.ERR_SYSTEM)
		} else {
			resp.DATA(data)
		}
	}
	return &resp
}

func (s *Service) Insert(req any) *contexts.RESPONSE {
	var resp contexts.RESPONSE
	// 写入数据
	id, err := s.Mdl.Insert(req)
	if err != nil {
		resp.STATE(contexts.ERR_INSERT)
	} else {
		resp.DATA(id)
	}
	return &resp
}

func (s *Service) Update(req any) *contexts.RESPONSE {
	var resp contexts.RESPONSE
	rows, err := s.Mdl.Update(req)
	if err != nil {
		resp.STATE(contexts.ERR_UPDATE)
	} else {
		resp.DATA(rows)
	}
	return &resp
}

func (s *Service) Remove(req any) *contexts.RESPONSE {
	var resp contexts.RESPONSE
	rows, err := s.Mdl.Update(req)
	if err != nil {
		resp.STATE(contexts.ERR_DELETE)
	} else {
		resp.DATA(rows)
	}
	return &resp
}
