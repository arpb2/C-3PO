package service

import (
	"github.com/arpb2/C-3PO/api/model"
	"github.com/stretchr/testify/mock"
)

type MockUserLevelService struct {
	mock.Mock
}

func (m MockUserLevelService) GetUserLevel(userId uint, levelId uint) (result model.UserLevel, err error) {
	args := m.Called(userId, levelId)

	result = args.Get(0).(model.UserLevel)
	err = args.Error(1)
	return
}

func (m MockUserLevelService) StoreUserLevel(userLevel model.UserLevel) (result model.UserLevel, err error) {
	args := m.Called(userLevel)

	result = args.Get(0).(model.UserLevel)
	err = args.Error(1)
	return
}
