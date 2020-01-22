package user_command_test

import (
	"errors"
	"github.com/arpb2/C-3PO/src/api/controller/user/user_command"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

func TestFetchUserCommand_Name(t *testing.T) {
	assert.Equal(t, "fetch_user_command", user_command.CreateFetchAuthenticatedUserCommand(nil).Name())
}

func TestFetchUserCommand_Fallback_DoesNothingOnUnknownError(t *testing.T) {
	command := user_command.CreateFetchAuthenticatedUserCommand(nil)
	runErr := errors.New("run err")

	assert.Equal(t, runErr, command.Fallback(runErr))
}

func TestFetchUserCommand_Fallback_ReturnsNilAndNotifies_OnBadRequest(t *testing.T) {
	middleware := new(http_wrapper.MockMiddleware)
	middleware.On("AbortTransactionWithStatus", http.StatusBadRequest, http_wrapper.Json{
		"error": "some message",
	}).Once()

	command := user_command.CreateFetchAuthenticatedUserCommand(&http_wrapper.Context{
		Reader: nil,
		Writer: nil,
		Middleware: middleware,
	})

	runErr := http_wrapper.CreateBadRequestError("some message")

	assert.Nil(t, command.Fallback(runErr))
	middleware.AssertExpectations(t)
}

func TestFetchUserCommand_Run_OnBadRead_Halts_NoErrorReturned(t *testing.T) {
	reader := new(http_wrapper.MockReader)
	reader.On("ReadBody", mock.MatchedBy(func(user *model.AuthenticatedUser) bool {
		return true
	})).Return(errors.New("bad read error")).Once()

	command := user_command.CreateFetchAuthenticatedUserCommand(&http_wrapper.Context{
			Reader: reader,
			Writer: nil,
			Middleware: nil,
		})

	err := command.Run()

	assert.NotNil(t, err)
	assert.Equal(t, http_wrapper.CreateBadRequestError("malformed body"), err)
	assert.Zero(t, len(command.OutputStream))
	reader.AssertExpectations(t)
}

func TestFetchUserCommand_Run_OnGoodRead_PublishesUser(t *testing.T) {
	expectedEmail := "test@email.com"
	expectedPassword := "test password"

	reader := new(http_wrapper.MockReader)
	reader.On("ReadBody", mock.MatchedBy(func(user *model.AuthenticatedUser) bool {
		user.User = &model.User{
			Email: expectedEmail,
		}
		user.Password = expectedPassword
		return true
	})).Return(nil).Once()

	command := user_command.CreateFetchAuthenticatedUserCommand(&http_wrapper.Context{
		Reader: reader,
		Writer: nil,
		Middleware: nil,
	})

	err := command.Run()

	assert.Nil(t, err)
	reader.AssertExpectations(t)
}

