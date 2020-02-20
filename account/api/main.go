package api

import (
	"encoding/json"
	"net/http"

	"github.com/kyhsa93/gin-rest-cqrs-example/config"
	"github.com/kyhsa93/gin-rest-cqrs-example/file/model"
)

// Interface api interace
type Interface interface {
	GetFileByID(fileID string) (*model.File, error)
}

// API api struct
type API struct {
	fileAPIURL string
}

// New create api instance
func New(config *config.Config) *API {
	return &API{fileAPIURL: config.Server.FileAPIAddress}
}

// GetFileByID get file data from file endpoint using file id
func (api *API) GetFileByID(fileID string) (*model.File, error) {
	response, htttRequestError := http.Get(api.fileAPIURL + "/" + fileID)
	if htttRequestError != nil {
		return nil, htttRequestError
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	var file *model.File
	responseBodyDecodeError := decoder.Decode(&file)
	if responseBodyDecodeError != nil {
		return nil, responseBodyDecodeError
	}
	return file, nil
}
