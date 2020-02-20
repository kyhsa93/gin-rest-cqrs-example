package api

import "github.com/kyhsa93/gin-rest-cqrs-example/config"

// Interface api interace
type Interface interface {
	GetFileByID(fileID string)
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
func (api *API) GetFileByID(fileID string) {

}
