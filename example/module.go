package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/ijodkor/oid"
)

func Register(
	api *gin.RouterGroup,
) {
	var oneHandler = oid.GetController()
	var oneSrv = oid.GetService()
	var handler = CrtHandler(oneSrv)

	router := api.Group("/auth")
	router.GET("/one-id/url", oneHandler.GetUrl)  // get url
	router.POST("/one-id/access", handler.Access) // redirect url
}
