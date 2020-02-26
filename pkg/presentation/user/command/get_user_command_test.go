package command_test

import (
	"errors"
	"testing"

	model2 "github.com/arpb2/C-3PO/pkg/domain/user/model"

	"github.com/arpb2/C-3PO/pkg/presentation/user/command"

	"github.com/arpb2/C-3PO/pkg/domain/architecture/http"
	"github.com/arpb2/C-3PO/test/mock/repository"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestGetUserCommand_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	cmd := command.CreateGetUserCommand(nil)

	name := cmd.Name()

	assert.Equal(t, "get_user_command", name)
}

func TestGetUserCommand_GivenOneAndAContextWithoutAuthenticatedUser_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	cmd := command.CreateGetUserCommand(nil)

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestGetUserCommand_GivenOneAndAFailingRepository_WhenRunning_ThenRepositoryError(t *testing.T) {
	expectedVal := uint(1000)
	ctx := gopipeline.CreateContext()
	ctx.Set(command.TagUserId, expectedVal)
	expectedErr := errors.New("some error")
	s := new(repository.MockUserRepository)
	s.On("GetUser", expectedVal).Return(nil, expectedErr)
	cmd := command.CreateGetUserCommand(s)

	err := cmd.Run(ctx)

	assert.Equal(t, expectedErr, err)
	s.AssertExpectations(t)
}

func TestGetUserCommand_GivenOneAndANoUserCreatedRepository_WhenRunning_Then404(t *testing.T) {
	expectedVal := uint(1000)
	ctx := gopipeline.CreateContext()
	ctx.Set(command.TagUserId, expectedVal)
	expectedErr := http.CreateNotFoundError()
	s := new(repository.MockUserRepository)
	s.On("GetUser", expectedVal).Return(model2.User{}, http.CreateNotFoundError())
	cmd := command.CreateGetUserCommand(s)

	err := cmd.Run(ctx)

	assert.Equal(t, expectedErr, err)
	s.AssertExpectations(t)
}

func TestGetUserCommand_GivenOne_WhenRunning_ThenContextHasUserAndReturnsNoError(t *testing.T) {
	expectedVal := model2.User{}
	ctx := gopipeline.CreateContext()
	ctx.Set(command.TagUserId, uint(1000))
	s := new(repository.MockUserRepository)
	s.On("GetUser", uint(1000)).Return(expectedVal, nil)
	cmd := command.CreateGetUserCommand(s)

	err := cmd.Run(ctx)
	val, exists := ctx.Get(command.TagUser)

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, expectedVal, val)
	s.AssertExpectations(t)
}
