package credential

import (
	credentialservice "github.com/arpb2/C-3PO/api/service/credential"
)

func CreateService() credentialservice.Service {
	return &credentialService{}
}

type credentialService struct{}

func (c credentialService) StoreCredentials(email, password string, userId uint) error {
	// Stores userId / email / hash(password).
	// User service will probably need to consume it on create.
	panic("implement me")
}

func (c credentialService) GetUserId(email, password string) (uint, error) {
	// Asserts email / hash(password) -> returns userId, error
	panic("implement me")
}
