package auth

import (
	"errors"
	"github.com/savaki/httpctx"
	"net/http"
	"os"
)

type ApiKey struct {
	Id     string
	Secret string
	Url    string
}

func GetAuth(id, secret string) (apiKey ApiKey, err error) {
	if id != "" && secret != "" {
		return ApiKey{Id: id, Secret: secret}, nil
	}

	return EnvAuth()
}

func EnvAuth() (apiKey ApiKey, err error) {
	apiKey.Id = os.Getenv("STORMPATH_API_KEY_ID")
	if apiKey.Id == "" {
		apiKey.Id = os.Getenv("STORMPATH_API_ID")
	}

	apiKey.Secret = os.Getenv("STORMPATH_API_KEY_SECRET")
	if apiKey.Secret == "" {
		apiKey.Secret = os.Getenv("STORMPATH_API_SECRET")
	}

	apiKey.Url = os.Getenv("STORMPATH_URL")

	if apiKey.Id == "" {
		err = errors.New("STORMPATH_API_KEY_ID or STORMPATH_API_ID not found in environment")
	}
	if apiKey.Secret == "" {
		err = errors.New("STORMPATH_API_KEY_SECRET or STORMPATH_API_SECRET not found in environment")
	}
	return
}

func Authenticator(apiKey ApiKey) httpctx.AuthFunc {
	_apiKey := apiKey
	return func(req *http.Request) *http.Request {
		req.SetBasicAuth(_apiKey.Id, _apiKey.Secret)
		return req
	}
}
