package oid

type OneIdIdentityRequest struct {
	Code string `json:"code" g:"required"`
}

type TokenDto struct {
	Code        string
	RedirectUri string
}

type OneIdToken struct {
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

type OneIdUrlRequest struct {
	State string `json:"state" form:"state" g:"required"`
}
