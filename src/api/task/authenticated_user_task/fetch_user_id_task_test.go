package authenticated_user_task_test

import (
	"errors"
	"github.com/arpb2/C-3PO/src/api/model"
	"github.com/arpb2/C-3PO/src/api/task/authenticated_user_task"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type credentialService struct{
	mock.Mock
}

func (c credentialService) Store(user *model.AuthenticatedUser) error {
	args := c.Called(user)
	return args.Error(0)
}

func (c credentialService) Retrieve(email, password string) (uint, error) {
	args := c.Called(email, password)
	return args.Get(0).(uint), args.Error(1)
}

func TestFetchUserIdTaskImpl_FailsOnServiceFailure(t *testing.T) {
	service := new(credentialService)
	service.On("Retrieve", "test@email.com", "testpassword").Return(uint(0), errors.New("error")).Once()

	userId, err := authenticated_user_task.FetchUserIdTaskImpl(
		service,
		&model.AuthenticatedUser{
			User:     &model.User{
				Email:   "test@email.com",
			},
			Password: "testpassword",
		},
	)

	assert.Zero(t, userId)
	assert.NotNil(t, err)
	service.AssertExpectations(t)
}

func TestFetchUserIdTaskImpl_ReturnsServiceResult(t *testing.T) {
	service := new(credentialService)
	service.On("Retrieve", "test@email.com", "testpassword").Return(uint(1000), nil).Once()

	userId, err := authenticated_user_task.FetchUserIdTaskImpl(
		service,
		&model.AuthenticatedUser{
			User:     &model.User{
				Email:   "test@email.com",
			},
			Password: "testpassword",
		},
	)

	assert.Equal(t, uint(1000), userId)
	assert.Nil(t, err)
	service.AssertExpectations(t)
}