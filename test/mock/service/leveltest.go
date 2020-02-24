package service

import (
	"github.com/arpb2/C-3PO/pkg/domain/model"
	"github.com/stretchr/testify/mock"
)

type MockLevelService struct {
	mock.Mock
}

func (m *MockLevelService) GetLevel(levelId uint) (result model.Level, err error) {
	args := m.Called(levelId)

	result = args.Get(0).(model.Level)
	err = args.Error(1)
	return
}

func (m *MockLevelService) StoreLevel(level model.Level) (result model.Level, err error) {
	args := m.Called(level)

	result = args.Get(0).(model.Level)
	err = args.Error(1)
	return
}
