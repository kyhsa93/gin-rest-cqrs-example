package controller

import "github.com/kyhsa93/go-rest-example/account/repository"

// Delete account by accountID
func Delete(accountID string) {
	repository.Delete(accountID)
}
