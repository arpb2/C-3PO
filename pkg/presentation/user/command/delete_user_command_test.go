package command_test

import (
	"errors"
	"testing"

	"github.com/arpb2/C-3PO/pkg/presentation/user/command"

	"github.com/arpb2/C-3PO/pkg/domain/architecture/http"
	"github.com/arpb2/C-3PO/test/mock/service"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestDeleteUserCommand_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	cmd := command.CreateDeleteUserCommand(nil)

	name := cmd.Name()

	assert.Equal(t, "delete_user_command", name)
}

func TestDeleteUserCommand_GivenOneAndAContextWithoutUserId_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	cmd := command.CreateDeleteUserCommand(nil)

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestDeleteUserCommand_GivenOneAndAFailingService_WhenRunning_ThenServiceError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(command.TagUserId, uint(1000))
	expectedErr := errors.New("some error")
	s := new(service.MockUserService)
	s.On("DeleteUser", uint(1000)).Return(expectedErr)
	cmd := command.CreateDeleteUserCommand(s)

	err := cmd.Run(ctx)

	assert.Equal(t, expectedErr, err)
	s.AssertExpectations(t)
}

func TestDeleteUserCommand_GivenOne_WhenRunning_ThenReturnsNoError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(command.TagUserId, uint(1000))
	s := new(service.MockUserService)
	s.On("DeleteUser", uint(1000)).Return(nil)
	cmd := command.CreateDeleteUserCommand(s)

	err := cmd.Run(ctx)

	assert.Nil(t, err)
	s.AssertExpectations(t)
}
