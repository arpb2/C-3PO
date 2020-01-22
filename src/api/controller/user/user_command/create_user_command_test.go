package user_command_test

import (
	"errors"
	"github.com/arpb2/C-3PO/src/api/controller/user/user_command"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestCreateUserCommand_Name(t *testing.T) {
	assert.Equal(t, "create_user_command", user_command.CreateCreateUserCommand(nil,  nil, nil).Name())
}

func TestCreateUserCommand_Fallback_Consumes_InternalError(t *testing.T) {
	runErr := http_wrapper.CreateInternalError()

	middleware := new(http_wrapper.MockMiddleware)
	middleware.On("AbortTransactionWithStatus", http.StatusInternalServerError, http_wrapper.Json{
		"error": runErr.Error(),
	})

	command := user_command.CreateCreateUserCommand(&http_wrapper.Context{
		Reader: nil,
		Writer: nil,
		Middleware: middleware,
	}, nil, nil)

	assert.Equal(t, nil, command.Fallback(runErr))
	middleware.AssertExpectations(t)
}

func TestCreateUserCommand_Fallback_DoesNothing_OnNonHttpError(t *testing.T) {
	command := user_command.CreateCreateUserCommand(nil, nil, nil)
	runErr := errors.New("run err")

	assert.Equal(t, runErr, command.Fallback(runErr))
}

func TestCreateUserCommand_Fallback_Halts_OnHttpError_NotInternal(t *testing.T) {
	middleware := new(http_wrapper.MockMiddleware)
	middleware.On("AbortTransactionWithStatus", http.StatusBadRequest, http_wrapper.Json{
		"error": "some message",
	})

	command := user_command.CreateCreateUserCommand(&http_wrapper.Context{
		Reader: nil,
		Writer: nil,
		Middleware: middleware,
	}, nil, nil)

	runErr := http_wrapper.CreateBadRequestError("some message")

	assert.Nil(t, command.Fallback(runErr))
	middleware.AssertExpectations(t)
}
