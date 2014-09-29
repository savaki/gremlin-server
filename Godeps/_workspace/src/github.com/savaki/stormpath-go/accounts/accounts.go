package accounts

import (
	"bytes"
	"code.google.com/p/go.net/context"
	"encoding/base64"
	"github.com/savaki/httpctx"
	"github.com/savaki/stormpath-go/auth"
	. "github.com/savaki/stormpath-go/types"
	"net/url"
)

type Accounts struct {
	client         httpctx.HttpClient
	ctx            context.Context
	applicationUrl string
}

func FromUrl(apiKey auth.ApiKey, applicationUrl string) *Accounts {
	authFunc := auth.Authenticator(apiKey)
	return &Accounts{
		client:         httpctx.WithAuthFunc(authFunc),
		ctx:            context.Background(),
		applicationUrl: applicationUrl,
	}
}

func (a *Accounts) WithContext(ctx context.Context) *Accounts {
	return &Accounts{
		client:         a.client,
		ctx:            ctx,
		applicationUrl: a.applicationUrl,
	}
}

func (a *Accounts) Create(account Account) (created *Account, err error) {
	created = &Account{}
	err = a.client.Post(a.ctx, a.applicationUrl+"/accounts", account, created)
	return
}

func (a *Accounts) LoginViaGoogle(accessCode string) (created *Account, err error) {
	created = &Account{}
	body := map[string]Provider{
		"providerData": Provider{
			ProviderId: "google",
			Code:       accessCode,
		},
	}
	err = a.client.Post(a.ctx, a.applicationUrl+"/accounts", body, created)
	return
}

func (a *Accounts) LoginViaFacebook(accessToken string) (created *Account, err error) {
	created = &Account{}
	body := map[string]Provider{
		"providerData": Provider{
			ProviderId:  "facebook",
			AccessToken: accessToken,
		},
	}
	err = a.client.Post(a.ctx, a.applicationUrl+"/accounts", body, created)
	return
}

func (a *Accounts) Search(email string) ([]Account, error) {
	results := struct {
		Items []Account `json:"items"`
	}{}
	params := &url.Values{}
	if email != "" {
		params.Add("email", email)
	}
	err := a.client.Get(a.ctx, a.applicationUrl+"/accounts", params, &results)
	return results.Items, err
}

func encode(username, password string) string {
	buf := &bytes.Buffer{}
	w := base64.NewEncoder(base64.StdEncoding, buf)
	w.Write([]byte(username + ":" + password))
	w.Close()

	return buf.String()
}

func (a *Accounts) Login(username, password string) (Url, error) {
	// request
	body := struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	}{
		Type:  "basic",
		Value: encode(username, password),
	}

	// response
	response := struct {
		Account Map `json:"account"`
	}{}

	err := a.client.Post(a.ctx, a.applicationUrl+"/loginAttempts", body, &response)
	if err != nil {
		return nil, err
	}

	return response.Account, nil
}

func (a *Accounts) Delete(account Url) error {
	return a.client.Delete(a.ctx, account.Url())
}
