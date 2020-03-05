package api

import (
	"encoding/json"
	"net/http"

	account "github.com/kyhsa93/gin-rest-cqrs-example/account/model"
	"github.com/kyhsa93/gin-rest-cqrs-example/config"
	file "github.com/kyhsa93/gin-rest-cqrs-example/file/model"
)

// Interface external api interface
type Interface interface {
	GetFileByID(fileID string) (*file.File, error)
	GetAccountByAccessToken(
		accessToken string,
	) (*account.Account, error)
}

// API api struct
type API struct {
	accountAPIURL string
	fileAPIURL    string
}

// New create api instance
func New(config *config.Config) *API {
	return &API{
		accountAPIURL: config.Server.AccountAPIAddress,
		fileAPIURL:    config.Server.FileAPIAddress,
	}
}

// GetAccountByAccessToken get account data from account service by accesstoken
func (api *API) GetAccountByAccessToken(
	accessToken string,
) (*account.Account, error) {
	accountServiceEndpoint := api.accountAPIURL
	request, createNewRequestError := http.NewRequest(
		"GET",
		accountServiceEndpoint,
		nil,
	)
	if createNewRequestError != nil {
		return nil, createNewRequestError
	}

	request.Header.Add("Authorization", accessToken)

	client := &http.Client{}
	response, httpRequestError := client.Do(request)
	if httpRequestError != nil {
		return nil, httpRequestError
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	var account *account.Account
	responseBodyDecodeError := decoder.Decode(&account)
	if responseBodyDecodeError != nil {
		return nil, responseBodyDecodeError
	}
	return account, nil
}

// GetFileByID get file data from file service usibg fileID
func (api *API) GetFileByID(fileID string) (*file.File, error) {
	response, httpRquestError := http.Get(api.fileAPIURL + "/" + fileID)
	if httpRquestError != nil {
		return nil, httpRquestError
	}

	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	var file *file.File
	responseBodyDecodeError := decoder.Decode(&file)
	if responseBodyDecodeError != nil {
		return nil, responseBodyDecodeError
	}

	return file, nil
}
