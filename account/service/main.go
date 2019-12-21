package service

import (
	"github.com/kyhsa93/go-rest-example/account/repository"
)

// Service account service struct
type Service struct {
	repository *repository.Repository
}

// NewService create account service instance
func NewService(repository *repository.Repository) *Service {
	return &Service{repository: repository}
}
