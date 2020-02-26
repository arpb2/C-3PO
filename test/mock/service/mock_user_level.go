package service

import (
	model2 "github.com/arpb2/C-3PO/pkg/domain/user_level/model"
	"github.com/stretchr/testify/mock"
)

type MockUserLevelService struct {
	mock.Mock
}

func (m *MockUserLevelService) GetUserLevel(userId uint, levelId uint) (result model2.UserLevel, err error) {
	args := m.Called(userId, levelId)

	result = args.Get(0).(model2.UserLevel)
	err = args.Error(1)
	return
}

func (m *MockUserLevelService) StoreUserLevel(userLevel model2.UserLevel) (result model2.UserLevel, err error) {
	args := m.Called(userLevel)

	result = args.Get(0).(model2.UserLevel)
	err = args.Error(1)
	return
}
