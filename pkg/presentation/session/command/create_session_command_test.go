package command_test

import (
	"errors"
	"testing"

	"github.com/arpb2/C-3PO/pkg/presentation/session/command"
	usercommand "github.com/arpb2/C-3PO/pkg/presentation/user/command"

	auth2 "github.com/arpb2/C-3PO/pkg/domain/auth"
	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/domain/model"
	"github.com/arpb2/C-3PO/test/mock/auth"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestCreateTokenCommand_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	cmd := command.CreateCreateSessionCommand(nil)

	name := cmd.Name()

	assert.Equal(t, "create_session_command", name)
}

func TestCreateTokenCommand_GivenOneAndAContextWithoutUserID_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	cmd := command.CreateCreateSessionCommand(nil)

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestCreateTokenCommand_GivenOneAndAFailingService_WhenRunning_ThenServiceError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(usercommand.TagUserId, uint(1000))
	expectedErr := errors.New("some error")
	s := new(auth.MockTokenHandler)
	s.On("Create", &auth2.Token{
		UserId: uint(1000),
	}).Return("", expectedErr)
	cmd := command.CreateCreateSessionCommand(s)

	err := cmd.Run(ctx)

	assert.Equal(t, expectedErr, err)
	s.AssertExpectations(t)
}

func TestCreateTokenCommand_GivenOne_WhenRunning_ThenContextHasUserIDAndReturnsNoError(t *testing.T) {
	expectedVal := model.Session{
		UserId: uint(1000),
		Token:  "token",
	}
	ctx := gopipeline.CreateContext()
	ctx.Set(usercommand.TagUserId, expectedVal.UserId)
	s := new(auth.MockTokenHandler)
	s.On("Create", &auth2.Token{
		UserId: expectedVal.UserId,
	}).Return(expectedVal.Token, nil)
	cmd := command.CreateCreateSessionCommand(s)

	err := cmd.Run(ctx)
	val, exists := ctx.Get(command.TagSession)

	assert.Nil(t, err)
	assert.Equal(t, expectedVal, val)
	assert.True(t, exists)
	s.AssertExpectations(t)
}
