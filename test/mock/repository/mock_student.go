package repository

import (
	"github.com/arpb2/C-3PO/pkg/domain/model/user"
	"github.com/stretchr/testify/mock"
)

type MockStudentRepository struct {
	mock.Mock
}

func (m *MockStudentRepository) GetStudents(userId uint) (students []user.User, err error) {
	args := m.Called(userId)
	return args.Get(0).([]user.User), args.Error(1)
}
