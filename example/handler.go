package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/ijodkor/oid"
	"github.com/ijodkor/rest/response"
	"github.com/ijodkor/rest/validation"
)

type Handler struct {
	oneSrv *oid.Service
}

func (h *Handler) Access(c *gin.Context) {
	// Validate
	req, e := validation.ValidatedBody[oid.OneIdIdentityRequest](c)
	if e {
		return
	}

	identity := h.oneSrv.GetIdentity(req)
	if identity == nil {
		response.Fail(c, "Identity not verified")
		return
	}

	// Write own logic here

	response.Success(c, gin.H{
		"token": "token",
	})
}

func CrtHandler(oneSrv *oid.Service) *Handler {
	return &Handler{
		oneSrv: oneSrv,
	}
}
