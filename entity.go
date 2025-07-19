package oid

import (
	"fmt"
	"github.com/ijodkor/num"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

const (
	ResponseTypeCode  = "one_code"
	GrandTypeToken    = "one_authorization_code"
	GrandTypeIdentity = "one_access_token_identify"
	GrandLogout       = "one_log_out"
	OneScope          = "legal"
)

type LegalEntity struct {
	IsBasic bool   `json:"is_basic"`
	Tin     string `json:"tin"`
	AcronUZ string `json:"acron_UZ"`
	LeTin   string `json:"le_tin"`
	LeName  string `json:"le_name"`
}

type Identity struct {
	Pin        string        `json:"pin"`
	PPortNo    string        `json:"pport_no"`
	UserId     string        `json:"user_id"`
	SurName    string        `json:"sur_name"`
	FirstName  string        `json:"first_name"`
	MidName    string        `json:"mid_name"`
	FullName   string        `json:"full_name"`
	UserType   string        `json:"user_type"`
	AuthMethod string        `json:"auth_method"`
	Valid      bool          `json:"valid"`
	SessId     string        `json:"sess_id"` // It is same with access token
	RetCd      string        `json:"ret_cd"`
	LegalInfo  []LegalEntity `json:"legal_info"`
	// ValidMethods []string      `json:"valid_methods"`

	Scope string `json:"scope"`
}

func (i Identity) GetFullName() string {
	names := strings.Split(i.MidName, " ")
	midName := cases.Title(language.Uzbek).String(names[0])
	if len(names) > 1 {
		midName += " " + strings.ToLower(names[1])
	}

	return fmt.Sprintf("%s %s %s",
		cases.Title(language.Uzbek).String(i.SurName),
		cases.Title(language.Uzbek).String(i.FirstName),
		midName,
	)
}

func (i Identity) GetPin() int64 {
	return num.ToInt64(i.Pin)
}
