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

func TestCreateUserCommand_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	cmd := user.CreateCreateUserCommand(nil)

	name := cmd.Name()

	assert.Equal(t, "create_user_command", name)
}

func TestCreateUserCommand_GivenOneAndAContextWithoutAuthenticatedUser_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	cmd := user.CreateCreateUserCommand(nil)

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestCreateUserCommand_GivenOneAndAFailingService_WhenRunning_ThenServiceError(t *testing.T) {
	expectedVal := model.AuthenticatedUser{}
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagAuthenticatedUser, expectedVal)
	expectedErr := errors.New("some error")
	s := new(service.MockUserService)
	s.On("CreateUser", expectedVal.User).Return(expectedVal.User, expectedErr)
	cmd := user.CreateCreateUserCommand(s)

	err := cmd.Run(ctx)

	assert.Equal(t, expectedErr, err)
	s.AssertExpectations(t)
}

func TestCreateUserCommand_GivenOne_WhenRunning_ThenContextHasUserAndReturnsNoError(t *testing.T) {
	expectedVal := model.AuthenticatedUser{}
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagAuthenticatedUser, expectedVal)
	s := new(service.MockUserService)
	s.On("CreateUser", expectedVal.User).Return(expectedVal.User, nil)
	cmd := user.CreateCreateUserCommand(s)

	err := cmd.Run(ctx)
	val, exists := ctx.Get(user.TagUser)

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, expectedVal.User, val)
	s.AssertExpectations(t)
}
