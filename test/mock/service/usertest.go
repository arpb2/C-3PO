package service

import (
	"github.com/arpb2/C-3PO/pkg/domain/model"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m MockUserService) GetUser(userId uint) (user model.User, err error) {
	args := m.Called(userId)

	firstParam := args.Get(0)
	if firstParam != nil {
		user = firstParam.(model.User)
	}

	err = args.Error(1)
	return
}

func (m MockUserService) CreateUser(user model.AuthenticatedUser) (result model.User, err error) {
	args := m.Called(user)

	result = args.Get(0).(model.User)
	err = args.Error(1)
	return
}

func (m MockUserService) UpdateUser(user model.AuthenticatedUser) (result model.User, err error) {
	args := m.Called(user)

	result = args.Get(0).(model.User)
	err = args.Error(1)
	return
}

func (m MockUserService) DeleteUser(userId uint) error {
	args := m.Called(userId)
	return args.Error(0)
}
