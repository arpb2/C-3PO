package service

import (
	"github.com/arpb2/C-3PO/src/api/model"
)

type CredentialService interface {

	Store(user *model.AuthenticatedUser) error

	Retrieve(email, password string) (uint, error)

}
