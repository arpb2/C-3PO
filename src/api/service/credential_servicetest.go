package service

import (
	"github.com/arpb2/C-3PO/src/api/model"
	"github.com/stretchr/testify/mock"
)

type MockCredentialService struct{
	mock.Mock
}

func (c MockCredentialService) Store(user *model.AuthenticatedUser) error {
	args := c.Called(user)
	return args.Error(0)
}

func (c MockCredentialService) Retrieve(email, password string) (uint, error) {
	args := c.Called(email, password)
	return args.Get(0).(uint), args.Error(1)
}
