package credential_service

import (
	"github.com/arpb2/C-3PO/api/model"
	"github.com/arpb2/C-3PO/api/service/credential"
)

func CreateService() credential_service.Service {
	return &credentialService{}
}

type credentialService struct{}

func (c credentialService) Store(user *model.AuthenticatedUser) error {
	// Stores userId / email / hash(password).
	// User service will probably need to consume it on create.
	panic("implement me")
}

func (c credentialService) Retrieve(email, password string) (uint, error) {
	// Asserts email / hash(password) -> returns userId, error
	panic("implement me")
}
