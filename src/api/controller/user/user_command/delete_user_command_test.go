package user_command_test

import (
	"errors"
	"github.com/arpb2/C-3PO/src/api/controller/user/user_command"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteUserCommand_Name(t *testing.T) {
	assert.Equal(t, "delete_user_command", user_command.CreateDeleteUserCommand(nil,  nil, nil).Name())
}

func TestDeleteUserCommand_Fallback_DoesntConsume_HttpError(t *testing.T) {
	runErr := http_wrapper.CreateInternalError()

	command := user_command.CreateDeleteUserCommand(&http_wrapper.Context{
		Reader: nil,
		Writer: nil,
		Middleware: nil,
	}, nil, nil)

	assert.Equal(t, runErr, command.Fallback(runErr))
}

func TestDeleteUserCommand_Fallback_DoesNothing_OnNonHttpError(t *testing.T) {
	command := user_command.CreateDeleteUserCommand(nil, nil, nil)
	runErr := errors.New("run err")

	assert.Equal(t, runErr, command.Fallback(runErr))
}
