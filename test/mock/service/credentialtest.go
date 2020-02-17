package service

import (
	"github.com/stretchr/testify/mock"
)

type MockCredentialService struct {
	mock.Mock
}

func (c MockCredentialService) GetUserId(email, password string) (uint, error) {
	args := c.Called(email, password)
	return args.Get(0).(uint), args.Error(1)
}
