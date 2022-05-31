package service

import (
	"database/sql"
	"errors"

	"github.com/fbbyqsyea/gin-framework/contexts"
)

type Model interface {
	Get(req, resp interface{}) error
	Count(req, resp interface{}) error
	Select(req, resp interface{}) error
	Insert(req, resp interface{}) error
	Update(req, resp interface{}) error
}

type Service struct {
	Mdl Model
}

func NewService(mdl Model) *Service {
	return &Service{
		Mdl: mdl,
	}
}

func (s *Service) Get(req, data any) *contexts.RESPONSE {
	var resp contexts.RESPONSE
	err := s.Mdl.Get(req, data)
	if errors.Is(err, sql.ErrNoRows) {
		resp.STATE(contexts.ERR_USER_NOT_EXISTS)
	} else if err != nil {
		resp.STATE(contexts.ERR_SYSTEM)
	} else {
		resp.DATA(data)
	}
	return &resp
}

func (s *Service) List(req, data any) *contexts.RESPONSEWITHCOUNT {
	var resp contexts.RESPONSEWITHCOUNT
	var count contexts.COUNT
	err := s.Mdl.Count(req, &count)
	if err != nil {
		resp.STATE(contexts.ERR_SYSTEM)
	} else {
		resp.COUNT(count.Count)
		err := s.Mdl.Select(req, data)
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
	var id int64
	// 写入数据
	err := s.Mdl.Insert(req, &id)
	if err != nil {
		resp.STATE(contexts.ERR_INSERT)
	} else {
		resp.DATA(id)
	}
	return &resp
}

func (s *Service) Update(req any) *contexts.RESPONSE {
	var resp contexts.RESPONSE
	var rows int64
	err := s.Mdl.Update(req, &rows)
	if err != nil {
		resp.STATE(contexts.ERR_UPDATE)
	} else {
		resp.DATA(rows)
	}
	return &resp
}

func (s *Service) Remove(req any) *contexts.RESPONSE {
	var resp contexts.RESPONSE
	var rows int64
	err := s.Mdl.Update(req, &rows)
	if err != nil {
		resp.STATE(contexts.ERR_DELETE)
	} else {
		resp.DATA(rows)
	}
	return &resp
}
