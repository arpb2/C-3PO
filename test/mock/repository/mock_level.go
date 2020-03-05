package repository

import (
	"github.com/arpb2/C-3PO/pkg/domain/model/level"
	"github.com/stretchr/testify/mock"
)

type MockLevelRepository struct {
	mock.Mock
}

func (m *MockLevelRepository) GetLevel(levelId uint) (result level.Level, err error) {
	args := m.Called(levelId)

	result = args.Get(0).(level.Level)
	err = args.Error(1)
	return
}

func (m *MockLevelRepository) StoreLevel(l level.Level) (result level.Level, err error) {
	args := m.Called(l)

	result = args.Get(0).(level.Level)
	err = args.Error(1)
	return
}
