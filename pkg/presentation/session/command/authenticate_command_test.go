package command_test

import (
	"errors"
	"github.com/arpb2/C-3PO/pkg/presentation/session/command"
	command2 "github.com/arpb2/C-3PO/pkg/presentation/user/command"
	"testing"

	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/domain/model"
	"github.com/arpb2/C-3PO/test/mock/service"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticateCommand_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	cmd := command.CreateAuthenticateCommand(nil)

	name := cmd.Name()

	assert.Equal(t, "authenticate_command", name)
}

func TestAuthenticateCommand_GivenOneAndAContextWithoutAuthenticatedUser_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	cmd := command.CreateAuthenticateCommand(nil)

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestAuthenticateCommand_GivenOneAndAFailingService_WhenRunning_ThenServiceError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(command2.TagAuthenticatedUser, model.AuthenticatedUser{
		User: model.User{},
	})
	expectedErr := errors.New("some error")
	s := new(service.MockCredentialService)
	s.On("GetUserId", "", "").Return(uint(0), expectedErr)
	cmd := command.CreateAuthenticateCommand(s)

	err := cmd.Run(ctx)

	assert.Equal(t, expectedErr, err)
	s.AssertExpectations(t)
}

func TestAuthenticateCommand_GivenOne_WhenRunning_ThenContextHasUserIDAndReturnsNoError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(command2.TagAuthenticatedUser, model.AuthenticatedUser{
		User: model.User{},
	})
	expectedVal := uint(1000)
	s := new(service.MockCredentialService)
	s.On("GetUserId", "", "").Return(expectedVal, nil)
	cmd := command.CreateAuthenticateCommand(s)

	err := cmd.Run(ctx)
	val, exists := ctx.Get(command2.TagUserId)

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, expectedVal, val)
	s.AssertExpectations(t)
}
