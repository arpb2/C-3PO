package session_command_test

import (
	"errors"
	"github.com/arpb2/C-3PO/src/api/auth"
	"github.com/arpb2/C-3PO/src/api/controller/session/session_command"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/model"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestTokenCommand_Name(t *testing.T) {
	assert.Equal(t, "create_token_command", session_command.CreateCreateTokenCommand(nil, nil, nil).Name())
}

func TestTokenCommand_Fallback_DoesNothing(t *testing.T) {
	command := session_command.CreateCreateTokenCommand(nil, nil, nil)
	runErr := errors.New("run err")

	assert.Equal(t, runErr, command.Fallback(runErr))
}

func TestTokenCommand_Run_OnBadToken_InternalError_ReturnsError(t *testing.T) {
	expectedUserId := uint(1000)
	expectedErr := errors.New("err")

	input := make(chan *model.AuthenticatedUser, 1)
	tokenHandler := new(auth.MockTokenHandler)
	tokenHandler.On("Create", &auth.Token{
		UserId: expectedUserId,
	}).Return("", &auth.TokenError{
		Error:  expectedErr,
		Status: http.StatusInternalServerError,
	}).Once()

	command := session_command.CreateCreateTokenCommand(
		&http_wrapper.Context{
			Reader: nil,
			Writer: nil,
			Middleware: nil,
		},
		tokenHandler,
		input)

	input <- &model.AuthenticatedUser{
		User: &model.User{
			Id: expectedUserId,
		},
	}

	err := command.Run()

	assert.Equal(t, err, expectedErr)
	assert.Zero(t, len(command.OutputStream))
	tokenHandler.AssertExpectations(t)
}

func TestTokenCommand_Run_OnBadToken_OtherError_Halts_ReturnsNothing(t *testing.T) {
	expectedUserId := uint(1000)
	expectedErr := errors.New("err")

	input := make(chan *model.AuthenticatedUser, 1)
	tokenHandler := new(auth.MockTokenHandler)
	tokenHandler.On("Create", &auth.Token{
		UserId: expectedUserId,
	}).Return("", &auth.TokenError{
		Error:  expectedErr,
		Status: http.StatusBadRequest,
	}).Once()

	middleware := new(http_wrapper.MockMiddleware)
	middleware.On("AbortTransactionWithStatus", http.StatusBadRequest, http_wrapper.Json{
		"error": expectedErr.Error(),
	})

	command := session_command.CreateCreateTokenCommand(
		&http_wrapper.Context{
			Reader: nil,
			Writer: nil,
			Middleware: middleware,
		},
		tokenHandler,
		input)

	input <- &model.AuthenticatedUser{
		User: &model.User{
			Id: expectedUserId,
		},
	}

	err := command.Run()

	assert.Nil(t, err)
	assert.Zero(t, len(command.OutputStream))
	tokenHandler.AssertExpectations(t)
	middleware.AssertExpectations(t)
}

func TestTokenCommand_Run_OnGoodToken_PublishesToken(t *testing.T) {
	expectedUserId := uint(1000)
	expectedToken := "some token"

	input := make(chan *model.AuthenticatedUser, 1)
	tokenHandler := new(auth.MockTokenHandler)
	tokenHandler.On("Create", &auth.Token{
		UserId: expectedUserId,
	}).Return(expectedToken, nil).Once()

	command := session_command.CreateCreateTokenCommand(
		&http_wrapper.Context{
			Reader: nil,
			Writer: nil,
			Middleware: nil,
		},
		tokenHandler,
		input)

	input <- &model.AuthenticatedUser{
		User: &model.User{
			Id: expectedUserId,
		},
	}

	err := command.Run()

	assert.Nil(t, err)

	token := <-command.OutputStream
	assert.NotNil(t, token)
	assert.Equal(t, expectedToken, token)
	tokenHandler.AssertExpectations(t)
}

