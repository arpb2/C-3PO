package user_test

import (
	"errors"
	"testing"

	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/api/model"
	"github.com/arpb2/C-3PO/pkg/command/user"
	"github.com/arpb2/C-3PO/test/mock/service"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestGetUserCommand_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	cmd := user.CreateGetUserCommand(nil)

	name := cmd.Name()

	assert.Equal(t, "get_user_command", name)
}

func TestGetUserCommand_GivenOneAndAContextWithoutAuthenticatedUser_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	cmd := user.CreateGetUserCommand(nil)

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestGetUserCommand_GivenOneAndAFailingService_WhenRunning_ThenServiceError(t *testing.T) {
	expectedVal := uint(1000)
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagUserId, expectedVal)
	expectedErr := errors.New("some error")
	s := new(service.MockUserService)
	s.On("GetUser", expectedVal).Return(nil, expectedErr)
	cmd := user.CreateGetUserCommand(s)

	err := cmd.Run(ctx)

	assert.Equal(t, expectedErr, err)
	s.AssertExpectations(t)
}

func TestGetUserCommand_GivenOneAndANoUserCreatedService_WhenRunning_Then404(t *testing.T) {
	expectedVal := uint(1000)
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagUserId, expectedVal)
	expectedErr := http.CreateNotFoundError()
	s := new(service.MockUserService)
	s.On("GetUser", expectedVal).Return(model.User{}, http.CreateNotFoundError())
	cmd := user.CreateGetUserCommand(s)

	err := cmd.Run(ctx)

	assert.Equal(t, expectedErr, err)
	s.AssertExpectations(t)
}

func TestGetUserCommand_GivenOne_WhenRunning_ThenContextHasUserAndReturnsNoError(t *testing.T) {
	expectedVal := model.User{}
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagUserId, uint(1000))
	s := new(service.MockUserService)
	s.On("GetUser", uint(1000)).Return(expectedVal, nil)
	cmd := user.CreateGetUserCommand(s)

	err := cmd.Run(ctx)
	val, exists := ctx.Get(user.TagUser)

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, expectedVal, val)
	s.AssertExpectations(t)
}
