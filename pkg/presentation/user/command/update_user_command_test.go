package command_test

import (
	"errors"
	"testing"

	model2 "github.com/arpb2/C-3PO/pkg/domain/user/model"

	"github.com/arpb2/C-3PO/pkg/presentation/user/command"

	"github.com/arpb2/C-3PO/pkg/domain/architecture/http"
	"github.com/arpb2/C-3PO/test/mock/service"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestUpdateUserCommand_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	cmd := command.CreateUpdateUserCommand(nil)

	name := cmd.Name()

	assert.Equal(t, "update_user_command", name)
}

func TestUpdateUserCommand_GivenOneAndAContextWithoutAuthenticatedUser_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(command.TagUserId, uint(1000))
	cmd := command.CreateUpdateUserCommand(nil)

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestUpdateUserCommand_GivenOneAndAContextWithoutUserId_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(command.TagAuthenticatedUser, model2.AuthenticatedUser{})
	cmd := command.CreateUpdateUserCommand(nil)

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestUpdateUserCommand_GivenOneAndAFailingService_WhenRunning_ThenServiceError(t *testing.T) {
	expectedVal := model2.AuthenticatedUser{
		User: model2.User{
			Id: uint(1000),
		},
	}
	ctx := gopipeline.CreateContext()
	ctx.Set(command.TagUserId, uint(1000))
	ctx.Set(command.TagAuthenticatedUser, expectedVal)
	expectedErr := errors.New("some error")
	s := new(service.MockUserService)
	s.On("UpdateUser", expectedVal).Return(expectedVal.User, expectedErr)
	cmd := command.CreateUpdateUserCommand(s)

	err := cmd.Run(ctx)

	assert.Equal(t, expectedErr, err)
	s.AssertExpectations(t)
}

func TestUpdateUserCommand_GivenOne_WhenRunning_ThenContextHasUserAndReturnsNoError(t *testing.T) {
	expectedVal := model2.AuthenticatedUser{
		User: model2.User{
			Id: uint(1000),
		},
	}
	ctx := gopipeline.CreateContext()
	ctx.Set(command.TagUserId, uint(1000))
	ctx.Set(command.TagAuthenticatedUser, expectedVal)
	s := new(service.MockUserService)
	s.On("UpdateUser", expectedVal).Return(expectedVal.User, nil)
	cmd := command.CreateUpdateUserCommand(s)

	err := cmd.Run(ctx)
	val, exists := ctx.Get(command.TagUser)

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, expectedVal.User, val)
	s.AssertExpectations(t)
}
