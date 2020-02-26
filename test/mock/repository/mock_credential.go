package repository

import (
	"github.com/stretchr/testify/mock"
)

type MockCredentialRepository struct {
	mock.Mock
}

func (c *MockCredentialRepository) GetUserId(email, password string) (uint, error) {
	args := c.Called(email, password)
	return args.Get(0).(uint), args.Error(1)
}
