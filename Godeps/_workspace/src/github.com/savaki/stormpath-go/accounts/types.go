package accounts

import (
	. "github.com/savaki/stormpath-go/types"
)

type Account struct {
	Href                   string `json:"href,omitempty,omitempty"`
	Username               string `json:"username,omitempty"`
	Email                  string `json:"email,omitempty"`
	Password               string `json:"password,omitempty"`
	FullName               string `json:"fullName,omitempty"`
	GivenName              string `json:"givenName,omitempty"`
	MiddleName             string `json:"middleName,omitempty"`
	Surname                string `json:"surname,omitempty"`
	Status                 string `json:"status,omitempty"`
	CustomData             Map    `json:"customData,omitempty"`
	Groups                 Map    `json:"groups,omitempty"`
	GroupMemberships       Map    `json:"groupMemberships,omitempty"`
	Directory              Map    `json:"directory,omitempty"`
	Tenant                 Map    `json:"tenant,omitempty"`
	EmailVerificationToken string `json:"emailVerificationToken,omitempty"`
}

func (a Account) Url() string {
	return a.Href
}

type Provider struct {
	ProviderId  string `json:"providerId"`
	Code        string `json:"code,omitempty"`
	AccessToken string `json:"accessToken,omitempty"`
}
