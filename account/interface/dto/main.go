package dto

// Provider account service provider map
func Provider() map[string]string {
	return map[string]string{
		"email": "email",
		"gmail": "gmail",
	}
}

// FilterAccountAttributeByProvider remove socialID or password by provider
func FilterAccountAttributeByProvider(data *Account) {
	if data.Provider == Provider()["email"] {
		data.SocialID = ""
		return
	}
	data.Password = ""
	return
}

// ValidateAccountAttributeByProvider validate account attribute by provider
func ValidateAccountAttributeByProvider(data *Account) bool {
	if data.Provider == Provider()["email"] && data.Password == "" {
		return false
	}
	if data.Provider != Provider()["email"] && data.SocialID == "" {
		return false
	}
	return true
}

// Gender user gender map
func Gender() map[string]string {
	return map[string]string{
		"male":   "male",
		"female": "female",
	}
}

// ValidateAccountGender validation account gender attribute
func ValidateAccountGender(data *Account) bool {
	if Gender()[data.Gender] == "" {
		return false
	}
	return true
}

// Intereste part of user interested
func Intereste() map[string]string {
	return map[string]string{
		"develop": "develop",
		"design":  "design",
		"manage":  "manage",
	}
}

// ValidateInteresteAttribute validation account's intereste
func ValidateInteresteAttribute(data *Account) bool {
	if Intereste()[data.Intereste] == "" {
		return false
	}
	return true
}

// Account account dto for command action
type Account struct {
	Email     string `json:"email" example:"test@gmail.com" binding:"required"`
	Provider  string `json:"provider" example:"gmail" binding:"required"`
	SocialID  string `json:"social_id" example:"social_id"`
	Password  string `json:"password" example:"password"`
	Gender    string `json:"gender" example:"male"`
	Intereste string `json:"intereste" example:"develop"`
}
