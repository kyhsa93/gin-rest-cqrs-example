package service

// Update update account by accountID
func (service *Service) Update(accountID string, email string, provider string, socialID string, password string) {
	oldData := service.ReadAccountByID(accountID)
	if oldData == nil {
		return
	}
	service.repository.Save(accountID, email, provider, socialID, password)
}
