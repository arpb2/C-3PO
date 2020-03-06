package repository

import (
	"github.com/arpb2/C-3PO/pkg/domain/model/user"
	"github.com/stretchr/testify/mock"
)

type MockUserLevelRepository struct {
	mock.Mock
}

func (m *MockUserLevelRepository) GetUserLevel(userId uint, levelId uint) (result user.Level, err error) {
	args := m.Called(userId, levelId)

	result = args.Get(0).(user.Level)
	err = args.Error(1)
	return
}

func (m *MockUserLevelRepository) StoreUserLevel(userLevel user.Level) (result user.Level, err error) {
	args := m.Called(userLevel)

	result = args.Get(0).(user.Level)
	err = args.Error(1)
	return
}
