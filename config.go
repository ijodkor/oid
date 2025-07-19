package oid

type Config struct {
	Url          string
	ClientId     string
	ClientSecret string
	RedirectUrl  *string
}

func (c *Config) GetRedirectUrl() string {
	if c.RedirectUrl == nil {
		return ""
	}

	return *c.RedirectUrl
}
