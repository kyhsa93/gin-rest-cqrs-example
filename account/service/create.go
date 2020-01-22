package service

// Create create account
func (service *Service) Create(email string, provider string, socialID string, password string) {
	service.repository.Save("", email, provider, socialID, password)
}
