package api

import (
	"net/http"

	"github.com/fbbyqsyea/gin-framework/contexts"
	"github.com/gin-gonic/gin"
	"github.com/mcuadros/go-defaults"
)

type Service interface {
	Get(req, data interface{}) *contexts.RESPONSE
	List(req, data interface{}) *contexts.RESPONSEWITHCOUNT
	Insert(req interface{}) *contexts.RESPONSE
	Update(req interface{}) *contexts.RESPONSE
	Remove(req interface{}) *contexts.RESPONSE
}

type Api struct {
	Svc Service
}

func NewApi(s Service) *Api {
	return &Api{
		Svc: s,
	}
}

func (a *Api) ShouldBindJSON(c *gin.Context, req interface{}) bool {
	defaults.SetDefaults(req)
	err := c.ShouldBindJSON(req)
	if err != nil {
		a.JSON(c, contexts.NewRESPONSE().STATE(contexts.ERR_PARAMS))
		return false
	}
	return true
}

func (a *Api) ShouldBindQuery(c *gin.Context, req interface{}) bool {
	defaults.SetDefaults(req)
	err := c.ShouldBindQuery(req)
	if err != nil {
		a.JSON(c, contexts.NewRESPONSE().STATE(contexts.ERR_PARAMS))
		return false
	}
	return true
}

func (a *Api) JSON(c *gin.Context, resp interface{}) {
	defaults.SetDefaults(resp)
	c.JSON(http.StatusOK, resp)
}

func (a *Api) Get(c *gin.Context, req, data interface{}) {
	if a.ShouldBindQuery(c, req) {
		a.JSON(c, a.Svc.Get(req, data))
	}
}

func (a *Api) List(c *gin.Context, req, data interface{}) {
	if a.ShouldBindQuery(c, req) {
		a.JSON(c, a.Svc.List(req, data))
	}
}

func (a *Api) Insert(c *gin.Context, req interface{}) {
	if a.ShouldBindJSON(c, req) {
		a.JSON(c, a.Svc.Insert(req))
	}
}

func (a *Api) Update(c *gin.Context, req interface{}) {
	if a.ShouldBindJSON(c, req) {
		a.JSON(c, a.Svc.Update(req))
	}
}
func (a *Api) Status(c *gin.Context, req interface{}) {
	if a.ShouldBindJSON(c, req) {
		a.JSON(c, a.Svc.Update(req))
	}
}

func (a *Api) Remove(c *gin.Context, req interface{}) {
	if a.ShouldBindJSON(c, req) {
		a.JSON(c, a.Svc.Remove(req))
	}
}

func (a *Api) Removes(c *gin.Context, req interface{}) {
	if a.ShouldBindJSON(c, req) {
		a.JSON(c, a.Svc.Remove(req))
	}
}

func (a *Api) GetHandleFunc(req, data interface{}) func(*gin.Context) {
	return func(c *gin.Context) {
		if a.ShouldBindQuery(c, req) {
			a.JSON(c, a.Svc.Get(req, data))
		}
	}
}

func (a *Api) ListHandleFunc(req, data interface{}) func(*gin.Context) {
	return func(c *gin.Context) {
		if a.ShouldBindQuery(c, req) {
			a.JSON(c, a.Svc.List(req, data))
		}
	}
}

func (a *Api) InsertHandleFunc(req interface{}) func(*gin.Context) {
	return func(c *gin.Context) {
		if a.ShouldBindJSON(c, req) {
			a.JSON(c, a.Svc.Insert(req))
		}
	}
}

func (a *Api) UpdateHandleFunc(req interface{}) func(*gin.Context) {
	return func(c *gin.Context) {
		if a.ShouldBindJSON(c, req) {
			a.JSON(c, a.Svc.Update(req))
		}
	}
}
func (a *Api) StatusHandleFunc(req interface{}) func(*gin.Context) {
	return func(c *gin.Context) {
		if a.ShouldBindJSON(c, req) {
			a.JSON(c, a.Svc.Update(req))
		}
	}
}

func (a *Api) RemoveHandleFunc(req interface{}) func(*gin.Context) {
	return func(c *gin.Context) {
		if a.ShouldBindJSON(c, req) {
			a.JSON(c, a.Svc.Remove(req))
		}
	}
}

func (a *Api) RemovesHandleFunc(req interface{}) func(*gin.Context) {
	return func(c *gin.Context) {
		if a.ShouldBindJSON(c, req) {
			a.JSON(c, a.Svc.Remove(req))
		}
	}
}
