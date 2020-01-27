package service

import (
	"github.com/arpb2/C-3PO/api/model"
	"github.com/stretchr/testify/mock"
)

type MockCodeService struct {
	mock.Mock
}

func (m MockCodeService) GetCode(userId uint, codeId uint) (code *model.Code, err error) {
	args := m.Called(userId, codeId)

	codeParam := args.Get(0)
	if codeParam != nil {
		code = codeParam.(*model.Code)
	}

	err = args.Error(1)
	return
}

func (m MockCodeService) CreateCode(userId uint, code string) (codeModel *model.Code, err error) {
	args := m.Called(userId, code)

	codeParam := args.Get(0)
	if codeParam != nil {
		codeModel = codeParam.(*model.Code)
	}

	err = args.Error(1)
	return
}

func (m MockCodeService) ReplaceCode(code *model.Code) error {
	args := m.Called(code)

	return args.Error(0)
}
