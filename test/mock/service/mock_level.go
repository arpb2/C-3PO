package service

import (
	model2 "github.com/arpb2/C-3PO/pkg/domain/level/model"
	"github.com/stretchr/testify/mock"
)

type MockLevelService struct {
	mock.Mock
}

func (m *MockLevelService) GetLevel(levelId uint) (result model2.Level, err error) {
	args := m.Called(levelId)

	result = args.Get(0).(model2.Level)
	err = args.Error(1)
	return
}

func (m *MockLevelService) StoreLevel(level model2.Level) (result model2.Level, err error) {
	args := m.Called(level)

	result = args.Get(0).(model2.Level)
	err = args.Error(1)
	return
}
