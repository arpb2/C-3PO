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

func TestUpdateUserCommand_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	cmd := user.CreateUpdateUserCommand(nil)

	name := cmd.Name()

	assert.Equal(t, "update_user_command", name)
}

func TestUpdateUserCommand_GivenOneAndAContextWithoutAuthenticatedUser_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagUserId, uint(1000))
	cmd := user.CreateUpdateUserCommand(nil)

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestUpdateUserCommand_GivenOneAndAContextWithoutUserId_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagAuthenticatedUser, model.AuthenticatedUser{})
	cmd := user.CreateUpdateUserCommand(nil)

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestUpdateUserCommand_GivenOneAndAFailingService_WhenRunning_ThenServiceError(t *testing.T) {
	expectedVal := model.AuthenticatedUser{
		User: model.User{
			Id: uint(1000),
		},
	}
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagUserId, uint(1000))
	ctx.Set(user.TagAuthenticatedUser, expectedVal)
	expectedErr := errors.New("some error")
	s := new(service.MockUserService)
	s.On("UpdateUser", expectedVal.User).Return(expectedVal.User, expectedErr)
	cmd := user.CreateUpdateUserCommand(s)

	err := cmd.Run(ctx)

	assert.Equal(t, expectedErr, err)
	s.AssertExpectations(t)
}

func TestUpdateUserCommand_GivenOne_WhenRunning_ThenContextHasUserAndReturnsNoError(t *testing.T) {
	expectedVal := model.AuthenticatedUser{
		User: model.User{
			Id: uint(1000),
		},
	}
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagUserId, uint(1000))
	ctx.Set(user.TagAuthenticatedUser, expectedVal)
	s := new(service.MockUserService)
	s.On("UpdateUser", expectedVal.User).Return(expectedVal.User, nil)
	cmd := user.CreateUpdateUserCommand(s)

	err := cmd.Run(ctx)
	val, exists := ctx.Get(user.TagUser)

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, expectedVal.User, val)
	s.AssertExpectations(t)
}
