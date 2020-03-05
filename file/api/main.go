package api

import (
	"encoding/json"
	"net/http"

	"github.com/kyhsa93/gin-rest-cqrs-example/account/model"
	"github.com/kyhsa93/gin-rest-cqrs-example/config"
)

// Interface api interface
type Interface interface {
	GetAccountByAccessToken(
		accessToken string,
	) (*model.Account, error)
}

// API api struct
type API struct {
	accountAPIURL string
}

// New create api instance
func New(config *config.Config) *API {
	return &API{accountAPIURL: config.Server.AccountAPIAddress}
}

// GetAccountByAccessToken get account data from account service by accesstoken
func (api *API) GetAccountByAccessToken(
	accessToken string,
) (*model.Account, error) {
	accountServiceEndpoint := api.accountAPIURL
	request, createNewReqeustError := http.NewRequest(
		"GET",
		accountServiceEndpoint,
		nil,
	)
	if createNewReqeustError != nil {
		return nil, createNewReqeustError
	}

	request.Header.Add("Authorization", accessToken)

	client := &http.Client{}
	response, httpRequestError := client.Do(request)
	if httpRequestError != nil {
		return nil, httpRequestError
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	var account *model.Account
	responseBodyDecodeError := decoder.Decode(&account)
	if responseBodyDecodeError != nil {
		return nil, responseBodyDecodeError
	}
	return account, nil
}
