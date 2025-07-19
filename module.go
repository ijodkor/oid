package oid

import (
	"os"
	"sync"
)

var (
	once sync.Once
)

func getConfig() *Config {
	redirect := os.Getenv("ONE_ID_REDIRECT_URL")
	return &Config{
		Url:          os.Getenv("ONE_ID_SSO_URL"),
		ClientId:     os.Getenv("ONE_ID_CLIENT_ID"),
		ClientSecret: os.Getenv("ONE_ID_CLIENT_SECRET"),
		RedirectUrl:  &redirect,
	}
}

func Register() {
	// Once initialize instance
	once.Do(func() {
		cnf := getConfig()
		srv := CrtOneIdService(cnf)
		crtController(srv)
	})
}

func RegisterAsync(cnf *Config) {
	// Once initialize instance
	once.Do(func() {
		srv := CrtOneIdService(cnf)
		crtController(srv)
	})
}
