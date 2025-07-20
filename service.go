package oid

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
)

type IService interface {
	GetUrl(scope *string, state string) string
	GetToken(dto TokenDto) (OneIdToken, error)
	GetIdentity(dto OneIdIdentityRequest) *Identity
	Logout(dto OneIdIdentityRequest) bool
}

type service struct {
	conf *Config
}

var (
	sOnce       sync.Once
	srvInstance IService
)

func (s *service) GetUrl(scope *string, state string) string {
	if scope == nil {
		sc := OneScope
		scope = &sc
	}

	params := url.Values{}
	params.Add("response_type", ResponseTypeCode)
	params.Add("client_id", s.conf.ClientId)
	params.Add("redirect_uri", s.conf.GetRedirectUrl())
	params.Add("scope", *scope)
	params.Add("state", state)

	return fmt.Sprintf("%s?%s", s.conf.Url, params.Encode())
}

func (s *service) GetToken(dto TokenDto) (OneIdToken, error) {
	payload := url.Values{}
	payload.Set("grant_type", GrandTypeToken)
	payload.Set("client_id", s.conf.ClientId)
	payload.Set("client_secret", s.conf.ClientSecret)
	payload.Set("redirect_uri", dto.RedirectUri)
	payload.Set("code", dto.Code)

	body := payload.Encode()

	res, _ := http.Post(
		s.conf.Url,
		"application/x-www-form-urlencoded",
		bytes.NewBufferString(body),
	)

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	if res.StatusCode >= http.StatusBadRequest {
		return OneIdToken{}, nil
	}

	bodyBytes, _ := io.ReadAll(res.Body)

	var token OneIdToken
	err := json.Unmarshal(bodyBytes, &token)

	return token, err
}

func (s *service) GetIdentity(dto OneIdIdentityRequest) *Identity {
	token, err := s.GetToken(TokenDto{
		RedirectUri: *s.conf.RedirectUrl,
		Code:        dto.Code,
	})

	if err != nil {
		return nil
	}

	payload := url.Values{}
	payload.Set("grant_type", GrandTypeIdentity)
	payload.Set("client_id", s.conf.ClientId)
	payload.Set("client_secret", s.conf.ClientSecret)
	payload.Set("access_token", token.AccessToken)
	payload.Set("scope", token.Scope)

	body := payload.Encode()

	res, _ := http.Post(
		s.conf.Url,
		"application/x-www-form-urlencoded",
		bytes.NewBufferString(body),
	)

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	bodyBytes, _ := io.ReadAll(res.Body)

	var identity Identity
	_ = json.Unmarshal(bodyBytes, &identity)

	identity.Scope = token.Scope
	if len(identity.Pin) < 1 {
		return nil
	}

	return &identity
}

func (s *service) Logout(dto OneIdIdentityRequest) bool {
	token, err := s.GetToken(TokenDto{
		RedirectUri: *s.conf.RedirectUrl,
		Code:        dto.Code,
	})

	if err != nil {
		return false
	}

	payload := url.Values{}
	payload.Set("grant_type", GrandLogout)
	payload.Set("client_id", s.conf.ClientId)
	payload.Set("client_secret", s.conf.ClientSecret)
	payload.Set("access_token", token.AccessToken)
	payload.Set("scope", token.Scope)

	body := payload.Encode()

	res, _ := http.Post(
		s.conf.Url,
		"application/x-www-form-urlencoded",
		bytes.NewBufferString(body),
	)

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	// io.ReadAll(res.Body)

	return true
}

func CrtOneIdService(conf *Config) IService {
	sOnce.Do(func() {
		srvInstance = &service{
			conf: conf,
		}
	})

	return srvInstance
}

func GetService() IService {
	return srvInstance
}
