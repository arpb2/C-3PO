package repository

import (
	model2 "github.com/arpb2/C-3PO/pkg/domain/user/model"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetUser(userId uint) (user model2.User, err error) {
	args := m.Called(userId)

	firstParam := args.Get(0)
	if firstParam != nil {
		user = firstParam.(model2.User)
	}

	err = args.Error(1)
	return
}

func (m *MockUserRepository) CreateUser(user model2.AuthenticatedUser) (result model2.User, err error) {
	args := m.Called(user)

	result = args.Get(0).(model2.User)
	err = args.Error(1)
	return
}

func (m *MockUserRepository) UpdateUser(user model2.AuthenticatedUser) (result model2.User, err error) {
	args := m.Called(user)

	result = args.Get(0).(model2.User)
	err = args.Error(1)
	return
}

func (m *MockUserRepository) DeleteUser(userId uint) error {
	args := m.Called(userId)
	return args.Error(0)
}
