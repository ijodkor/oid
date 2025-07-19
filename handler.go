package oid

import (
	"github.com/gin-gonic/gin"
	"github.com/ijodkor/rest/response"
	"github.com/ijodkor/rest/validation"
)

type Handler struct {
	srv *Service
}

var Controller *Handler

func (h Handler) GetUrl(c *gin.Context) {
	r := validation.ValidatedQuery[OneIdUrlRequest](c)
	if r == nil {
		return
	}

	state, _ := c.GetQuery("state")

	response.Success(c, gin.H{
		"url": h.srv.GetUrl(nil, state),
	})
}

func crtController(srv *Service) *Handler {
	Controller = &Handler{
		srv: srv,
	}

	return Controller
}

func GetController() *Handler {
	return Controller
}
