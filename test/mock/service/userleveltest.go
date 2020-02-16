package service

import (
	"github.com/arpb2/C-3PO/api/model"
	"github.com/stretchr/testify/mock"
)

type MockUserLevelService struct {
	mock.Mock
}

func (m MockUserLevelService) GetUserLevel(userId uint, codeId uint) (code *model.UserLevel, err error) {
	args := m.Called(userId, codeId)

	codeParam := args.Get(0)
	if codeParam != nil {
		code = codeParam.(*model.UserLevel)
	}

	err = args.Error(1)
	return
}

func (m MockUserLevelService) CreateUserLevel(userId uint, code string) (codeModel *model.UserLevel, err error) {
	args := m.Called(userId, code)

	codeParam := args.Get(0)
	if codeParam != nil {
		codeModel = codeParam.(*model.UserLevel)
	}

	err = args.Error(1)
	return
}

func (m MockUserLevelService) ReplaceUserLevel(code *model.UserLevel) error {
	args := m.Called(code)

	return args.Error(0)
}
