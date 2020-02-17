package credential

import (
	credentialservice "github.com/arpb2/C-3PO/api/service/credential"
)

func CreateService() credentialservice.Service {
	return &credentialService{}
}

type credentialService struct{}

func (c credentialService) GetUserId(email, password string) (uint, error) {
	// Asserts email / hash(password) -> returns userId, error
	panic("implement me")
}
