# OneID package for Gin-Gonic - Single identification system (OneId) package

## Talablar (Requirements)

- Go ^v1.24
- Gin ^v1.10

## 2. Sozlash (Setup)

Muhit o&#8216;zgaruvchilari o&#8216;rnatilinadi (Set environment variables)

Majburiy (Mandatory)

```dotenv
ONE_ID_SSO_URL=<one_id_sso_url>
ONE_ID_CLIENT_ID=<client_id>
ONE_ID_CLIENT_SECRET=<client_secret>
```

Ixtiyoriy (Optional)

```dotenv
ONE_ID_REDIRECT_URL=<redirect_url>
```

## Foydalanish (Usage)

```go
package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/ijodkor/oid"
	"github.com/ijodkor/rest/response"
	"github.com/ijodkor/rest/validation"
)

// Auth module
// module.go

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

// ...
// handler.go

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

```

OneIDni bog&#8216;liq qism kiritilishi (DI) orqali ulash (Register OneID package with dependency injection - DI)

```go

package main

import (
	"go.uber.org/dig"
	"github.com/ijodkor/oid"
)

container := dig.New()

oid.Register()

_ = container.Provide(oneid.GetService)
_ = container.Provide(oneid.GetController)

```

### Foydalanilgan manbalar (References)