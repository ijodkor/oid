package oid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_GetUrl(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	redirectUrl := "http://localhost"
	srv := CrtOneIdService(&Config{
		Url:          "https://bir.joriy.uz",
		ClientId:     "string",
		ClientSecret: "string",
		RedirectUrl:  &redirectUrl,
	})

	url := srv.GetUrl(nil, "test")

	is.NotEmpty(url)
}
