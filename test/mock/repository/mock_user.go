package repository

import (
	"github.com/arpb2/C-3PO/pkg/domain/model/user"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetUser(userId uint) (u user.User, err error) {
	args := m.Called(userId)

	firstParam := args.Get(0)
	if firstParam != nil {
		u = firstParam.(user.User)
	}

	err = args.Error(1)
	return
}

func (m *MockUserRepository) CreateUser(u user.AuthenticatedUser) (result user.User, err error) {
	args := m.Called(u)

	result = args.Get(0).(user.User)
	err = args.Error(1)
	return
}

func (m *MockUserRepository) UpdateUser(u user.AuthenticatedUser) (result user.User, err error) {
	args := m.Called(u)

	result = args.Get(0).(user.User)
	err = args.Error(1)
	return
}

func (m *MockUserRepository) DeleteUser(userId uint) error {
	args := m.Called(userId)
	return args.Error(0)
}
