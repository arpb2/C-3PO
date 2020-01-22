package user_command_test

import (
	"errors"
	"github.com/arpb2/C-3PO/src/api/controller/session/session_command"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

func TestFetchUserCommand_Name(t *testing.T) {
	assert.Equal(t, "fetch_user_command", session_command.CreateFetchUserCommand(nil).Name())
}

func TestFetchUserCommand_Fallback_DoesNothing(t *testing.T) {
	command := session_command.CreateFetchUserCommand(nil)
	runErr := errors.New("run err")

	assert.Equal(t, runErr, command.Fallback(runErr))
}

func TestFetchUserCommand_Run_OnBadRead_Halts_NoErrorReturned(t *testing.T) {
	reader := new(http_wrapper.MockReader)
	reader.On("ReadBody", mock.MatchedBy(func(user *model.AuthenticatedUser) bool {
		return true
	})).Return(errors.New("bad read error")).Once()

	middleware := new(http_wrapper.MockMiddleware)
	middleware.On("AbortTransactionWithStatus", http.StatusBadRequest, http_wrapper.Json{
		"error": "malformed body",
	})

	command := session_command.CreateFetchUserCommand(&http_wrapper.Context{
			Reader: reader,
			Writer: nil,
			Middleware: middleware,
		})

	err := command.Run()

	assert.Nil(t, err)
	assert.Zero(t, len(command.OutputStream))
	middleware.AssertExpectations(t)
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

	command := session_command.CreateFetchUserCommand(&http_wrapper.Context{
		Reader: reader,
		Writer: nil,
		Middleware: nil,
	})

	err := command.Run()

	assert.Nil(t, err)
	reader.AssertExpectations(t)
}

