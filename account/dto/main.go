package dto

// Provider account service provider map
func Provider() map[string]string {
	return map[string]string{
		"email": "email",
		"gmail": "gmail",
	}
}

// Gender user gender map
func Gender() map[string]string {
	return map[string]string{
		"male":   "male",
		"female": "female",
	}
}

// InterestedField part of user InterestedFielded
func InterestedField() map[string]string {
	return map[string]string{
		"develop": "develop",
		"design":  "design",
		"manage":  "manage",
	}
}

// CreateAccount account dto for create command
type CreateAccount struct {
	Email                 string
	Provider              string
	SocialID              string
	Password              string
	FCMToken              string
	Gender                string
	InterestedField       string
	InterestedFieldDetail []string
}

// ValidateAccountGender validation account gender attribute
func (dto *CreateAccount) ValidateAccountGender() bool {
	if Gender()[dto.Gender] == "" {
		return false
	}
	return true
}

// ValidateInterestedFieldAttribute validation account's InterestedField
func (dto *CreateAccount) ValidateInterestedFieldAttribute() bool {
	if InterestedField()[dto.InterestedField] == "" {
		return false
	}
	return true
}

// FilterAccountAttributeByProvider remove socialID or password by provider
func (dto *CreateAccount) FilterAccountAttributeByProvider() {
	if dto.Provider == Provider()["email"] {
		dto.SocialID = ""
		return
	}
	dto.Password = ""
	return
}

// ValidateProvider validate provider
func (dto *CreateAccount) ValidateProvider() bool {
	_, validate := Provider()[dto.Provider]
	return validate
}

// ValidateAccountAttributeByProvider validate account attribute by provider
func (dto *CreateAccount) ValidateAccountAttributeByProvider() bool {
	if dto.Provider == Provider()["email"] && dto.Password == "" {
		return false
	}
	if dto.Provider != Provider()["email"] && dto.SocialID == "" {
		return false
	}
	return true
}

// UpdateAccount account dto for update command
type UpdateAccount struct {
	FCMToken              string
	Password              string
	InterestedField       string
	InterestedFieldDetail []string
}

// ValidateInterestedFieldAttribute validation account's InterestedField
func (dto *UpdateAccount) ValidateInterestedFieldAttribute() bool {
	if InterestedField()[dto.InterestedField] == "" {
		return false
	}
	return true
}

// ReadAccount account dto for query
type ReadAccount struct {
	Email    string
	Provider string
	SocialID string
	Password string
}

// ValidateAccountAttributeByProvider validate account attribute by provider
func (dto *ReadAccount) ValidateAccountAttributeByProvider() bool {
	if dto.Provider == Provider()["email"] && dto.Password == "" {
		return false
	}
	if dto.Provider != Provider()["email"] && dto.SocialID == "" {
		return false
	}
	return true
}

// FilterAccountAttributeByProvider remove socialID or password by provider
func (dto *ReadAccount) FilterAccountAttributeByProvider() {
	if dto.Provider == Provider()["email"] {
		dto.SocialID = ""
		return
	}
	dto.Password = ""
	return
}

// ValidateProvider validate provider
func (dto *ReadAccount) ValidateProvider() bool {
	_, validate := Provider()[dto.Provider]
	return validate
}
