package api

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/kyhsa93/gin-rest-cqrs-example/config"
	"github.com/kyhsa93/gin-rest-cqrs-example/file/model"
	file "github.com/kyhsa93/gin-rest-cqrs-example/file/model"
	"github.com/kyhsa93/gin-rest-cqrs-example/profile/dto"
	profile "github.com/kyhsa93/gin-rest-cqrs-example/profile/model"
)

// Interface api interace
type Interface interface {
	CreateProfile(
		accessToken string,
		accountID string,
		email string,
		gender string,
		interestedField string,
		interestedFieldDetail []string,
	) (*profile.Profile, error)
	GetFileByID(fileID string) (*model.File, error)
}

// API api struct
type API struct {
	fileAPIURL    string
	profileAPIURL string
}

// New create api instance
func New(config config.Interface) *API {
	return &API{
		fileAPIURL:    config.Server().FileServiceEndPoint(),
		profileAPIURL: config.Server().ProfileServiceEndPoint(),
	}
}

// CreateProfile request create profile to profile service
func (api *API) CreateProfile(
	accessToken string,
	accountID string,
	email string,
	gender string,
	interestedField string,
	interestedFieldDetail []string,
) (*profile.Profile, error) {
	profileDto := dto.CreateProfile{
		AccountID:             accountID,
		Email:                 email,
		Gender:                gender,
		InterestedField:       interestedField,
		InterestedFieldDetail: interestedFieldDetail,
	}

	byteData, err := json.Marshal(profileDto)
	if err != nil {
		panic(err)
	}

	request, err := http.NewRequest(
		"POST",
		api.profileAPIURL,
		bytes.NewBuffer(byteData),
	)

	if err != nil {
		panic(err)
	}

	request.Header.Add("Authorization", accessToken)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	var profile *profile.Profile
	responseBodyDecodeError := decoder.Decode(&profile)
	if responseBodyDecodeError != nil {
		return nil, responseBodyDecodeError
	}
	return profile, nil
}

// GetFileByID get file data from file endpoint using file id
func (api *API) GetFileByID(fileID string) (*file.File, error) {
	response, httpRequestError := http.Get(api.fileAPIURL + "/" + fileID)
	if httpRequestError != nil {
		return nil, httpRequestError
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
