package session_test

import (
	"errors"
	"testing"

	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/api/model"
	"github.com/arpb2/C-3PO/pkg/command/session"
	"github.com/arpb2/C-3PO/pkg/command/user"
	"github.com/arpb2/C-3PO/test/mock/service"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticateCommand_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	cmd := session.CreateAuthenticateCommand(nil)

	name := cmd.Name()

	assert.Equal(t, "authenticate_command", name)
}

func TestAuthenticateCommand_GivenOneAndAContextWithoutAuthenticatedUser_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	cmd := session.CreateAuthenticateCommand(nil)

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestAuthenticateCommand_GivenOneAndAFailingService_WhenRunning_ThenServiceError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagAuthenticatedUser, model.AuthenticatedUser{
		User: &model.User{},
	})
	expectedErr := errors.New("some error")
	s := new(service.MockCredentialService)
	s.On("Retrieve", "", "").Return(uint(0), expectedErr)
	cmd := session.CreateAuthenticateCommand(s)

	err := cmd.Run(ctx)

	assert.Equal(t, expectedErr, err)
	s.AssertExpectations(t)
}

func TestAuthenticateCommand_GivenOne_WhenRunning_ThenContextHasUserIDAndReturnsNoError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagAuthenticatedUser, model.AuthenticatedUser{
		User: &model.User{},
	})
	expectedVal := uint(1000)
	s := new(service.MockCredentialService)
	s.On("Retrieve", "", "").Return(expectedVal, nil)
	cmd := session.CreateAuthenticateCommand(s)

	err := cmd.Run(ctx)
	val, exists := ctx.Get(user.TagUserId)

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, expectedVal, val)
	s.AssertExpectations(t)
}
