package command_test

import (
	"errors"
	"testing"

	model2 "github.com/arpb2/C-3PO/pkg/domain/user/model"

	"github.com/arpb2/C-3PO/pkg/presentation/user/command"

	"github.com/arpb2/C-3PO/pkg/domain/infrastructure/http"
	"github.com/arpb2/C-3PO/test/mock/service"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserCommand_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	cmd := command.CreateCreateUserCommand(nil)

	name := cmd.Name()

	assert.Equal(t, "create_user_command", name)
}

func TestCreateUserCommand_GivenOneAndAContextWithoutAuthenticatedUser_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	cmd := command.CreateCreateUserCommand(nil)

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestCreateUserCommand_GivenOneAndAFailingService_WhenRunning_ThenServiceError(t *testing.T) {
	expectedVal := model2.AuthenticatedUser{}
	ctx := gopipeline.CreateContext()
	ctx.Set(command.TagAuthenticatedUser, expectedVal)
	expectedErr := errors.New("some error")
	s := new(service.MockUserService)
	s.On("CreateUser", expectedVal).Return(expectedVal.User, expectedErr)
	cmd := command.CreateCreateUserCommand(s)

	err := cmd.Run(ctx)

	assert.Equal(t, expectedErr, err)
	s.AssertExpectations(t)
}

func TestCreateUserCommand_GivenOne_WhenRunning_ThenContextHasUserAndReturnsNoError(t *testing.T) {
	expectedVal := model2.AuthenticatedUser{}
	ctx := gopipeline.CreateContext()
	ctx.Set(command.TagAuthenticatedUser, expectedVal)
	s := new(service.MockUserService)
	s.On("CreateUser", expectedVal).Return(expectedVal.User, nil)
	cmd := command.CreateCreateUserCommand(s)

	err := cmd.Run(ctx)
	val, exists := ctx.Get(command.TagUser)

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, expectedVal.User, val)
	s.AssertExpectations(t)
}
