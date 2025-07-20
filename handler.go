package oid

import (
	"github.com/gin-gonic/gin"
	"github.com/ijodkor/rest/response"
	"github.com/ijodkor/rest/validation"
)

type Handler struct {
	srv IService
}

var Controller *Handler

func (h Handler) GetUrl(c *gin.Context) {
	//state, _ := c.GetQuery("state")
	query := validation.ValidatedQuery[OneIdUrlRequest](c)
	if query == nil {
		return
	}

	response.Success(c, gin.H{
		"url": h.srv.GetUrl(query.Scope, query.State),
	})
}

func crtController(srv IService) *Handler {
	Controller = &Handler{
		srv: srv,
	}

	return Controller
}

func GetController() *Handler {
	return Controller
}
