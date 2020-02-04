package application

// Delete account by accountID
func (service *Service) Delete(accountID string) {
	service.infrastructure.Repository.Delete(accountID)
}
