package service

// Delete account by accountID
func (service *Service) Delete(accountID string) {
	service.repository.Delete(accountID)
}
