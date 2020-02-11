package user_test

import (
	"errors"
	"testing"

	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/pkg/command/user"
	"github.com/arpb2/C-3PO/test/mock/service"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestDeleteUserCommand_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	cmd := user.CreateDeleteUserCommand(nil)

	name := cmd.Name()

	assert.Equal(t, "delete_user_command", name)
}

func TestDeleteUserCommand_GivenOneAndAContextWithoutUserId_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	cmd := user.CreateDeleteUserCommand(nil)

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestDeleteUserCommand_GivenOneAndAFailingService_WhenRunning_ThenServiceError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagUserId, uint(1000))
	expectedErr := errors.New("some error")
	s := new(service.MockUserService)
	s.On("DeleteUser", uint(1000)).Return(expectedErr)
	cmd := user.CreateDeleteUserCommand(s)

	err := cmd.Run(ctx)

	assert.Equal(t, expectedErr, err)
	s.AssertExpectations(t)
}

func TestDeleteUserCommand_GivenOne_WhenRunning_ThenReturnsNoError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagUserId, uint(1000))
	s := new(service.MockUserService)
	s.On("DeleteUser", uint(1000)).Return(nil)
	cmd := user.CreateDeleteUserCommand(s)

	err := cmd.Run(ctx)

	assert.Nil(t, err)
	s.AssertExpectations(t)
}
