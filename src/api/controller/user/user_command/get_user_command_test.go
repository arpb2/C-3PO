package user_command_test

import (
	"errors"
	"github.com/arpb2/C-3PO/src/api/controller/user/user_command"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetUserCommand_Name(t *testing.T) {
	assert.Equal(t, "get_user_command", user_command.CreateGetUserCommand(nil,  nil, nil).Name())
}

func TestGetUserCommand_Fallback_Consumes_InternalError(t *testing.T) {
	runErr := http_wrapper.CreateInternalError()

	middleware := new(http_wrapper.MockMiddleware)
	middleware.On("AbortTransactionWithStatus", http.StatusInternalServerError, http_wrapper.Json{
		"error": runErr.Error(),
	})

	command := user_command.CreateGetUserCommand(&http_wrapper.Context{
		Reader: nil,
		Writer: nil,
		Middleware: middleware,
	}, nil, nil)

	assert.Equal(t, nil, command.Fallback(runErr))
	middleware.AssertExpectations(t)
}

func TestGetUserCommand_Fallback_DoesNothing_OnNonHttpError(t *testing.T) {
	command := user_command.CreateGetUserCommand(nil, nil, nil)
	runErr := errors.New("run err")

	assert.Equal(t, runErr, command.Fallback(runErr))
}

func TestGetUserCommand_Fallback_Halts_OnHttpError_NotInternal(t *testing.T) {
	middleware := new(http_wrapper.MockMiddleware)
	middleware.On("AbortTransactionWithStatus", http.StatusBadRequest, http_wrapper.Json{
		"error": "some message",
	})

	command := user_command.CreateGetUserCommand(&http_wrapper.Context{
		Reader: nil,
		Writer: nil,
		Middleware: middleware,
	}, nil, nil)

	runErr := http_wrapper.CreateBadRequestError("some message")

	assert.Nil(t, command.Fallback(runErr))
	middleware.AssertExpectations(t)
}
