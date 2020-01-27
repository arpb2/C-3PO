package service

import (
	"github.com/arpb2/C-3PO/api/model"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m MockUserService) GetUser(userId uint) (user *model.User, err error) {
	args := m.Called(userId)

	firstParam := args.Get(0)
	if firstParam != nil {
		user = firstParam.(*model.User)
	}

	err = args.Error(1)
	return
}

func (m MockUserService) CreateUser(authenticatedUser *model.AuthenticatedUser) (user *model.User, err error) {
	args := m.Called(authenticatedUser)

	firstParam := args.Get(0)
	if firstParam != nil {
		user = firstParam.(*model.AuthenticatedUser).User
	}

	err = args.Error(1)
	return
}

func (m MockUserService) UpdateUser(authenticatedUser *model.AuthenticatedUser) (user *model.User, err error) {
	args := m.Called(authenticatedUser)

	firstParam := args.Get(0)
	if firstParam != nil {
		user = firstParam.(*model.User)
	}

	err = args.Error(1)
	return
}

func (m MockUserService) DeleteUser(userId uint) error {
	args := m.Called(userId)
	return args.Error(0)
}
