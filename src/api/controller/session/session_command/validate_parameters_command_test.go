package session_command_test

import (
	"errors"
	"github.com/arpb2/C-3PO/src/api/controller/session/session_command"
	"github.com/arpb2/C-3PO/src/api/controller/session/session_validation"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/model"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestValidateParametersCommand_Name(t *testing.T) {
	assert.Equal(t, "validate_parameters_command", session_command.CreateValidateParametersCommand(nil, nil, nil).Name())
}

func TestValidateParametersCommand_Fallback_DoesNothing_OnInternalError(t *testing.T) {
	command := session_command.CreateValidateParametersCommand(nil, nil, nil)
	runErr := http_wrapper.CreateInternalError()

	assert.Equal(t, runErr, command.Fallback(runErr))
}

func TestValidateParametersCommand_Fallback_DoesNothing_OnNonHttpError(t *testing.T) {
	command := session_command.CreateValidateParametersCommand(nil, nil, nil)
	runErr := errors.New("run err")

	assert.Equal(t, runErr, command.Fallback(runErr))
}

func TestValidateParametersCommand_Fallback_Halts_OnHttpError_NotInternal(t *testing.T) {
	middleware := new(http_wrapper.MockMiddleware)
	middleware.On("AbortTransactionWithStatus", http.StatusBadRequest, http_wrapper.Json{
		"error": "some message",
	})

	command := session_command.CreateValidateParametersCommand(&http_wrapper.Context{
		Reader: nil,
		Writer: nil,
		Middleware: middleware,
	}, nil, nil)

	runErr := http_wrapper.CreateBadRequestError("some message")

	assert.Nil(t, command.Fallback(runErr))
	middleware.AssertExpectations(t)
}

func TestValidateParametersCommand_Run_OnWrongValidation_ReturnsError(t *testing.T) {
	expectedErr := http_wrapper.CreateBadRequestError("bad error")

	input := make(chan *model.AuthenticatedUser, 1)

	middleware := new(http_wrapper.MockMiddleware)
	middleware.On("AbortTransactionWithStatus", http.StatusBadRequest, http_wrapper.Json{
		"error": expectedErr.Error(),
	})

	finalFuncCalled := false

	command := session_command.CreateValidateParametersCommand(
		&http_wrapper.Context{
			Reader: nil,
			Writer: nil,
			Middleware: middleware,
		},
		input,
		[]session_validation.Validation{
			func(user *model.AuthenticatedUser) error {
				return nil
			},
			func(user *model.AuthenticatedUser) error {
				return expectedErr
			},
			func(user *model.AuthenticatedUser) error {
				finalFuncCalled = true
				return nil
			},
		})

	input <- &model.AuthenticatedUser{}

	err := command.Run()

	assert.NotNil(t, err)
	assert.Equal(t, expectedErr, err)
	assert.False(t, finalFuncCalled)
	assert.Zero(t, len(command.Stream))
}

func TestValidateParametersCommand_Run_OnGoodValidations_PublishesSameUser(t *testing.T) {
	expectedUser := &model.AuthenticatedUser{}
	finalFuncCalled := false

	input := make(chan *model.AuthenticatedUser, 1)
	command := session_command.CreateValidateParametersCommand(
		&http_wrapper.Context{
			Reader: nil,
			Writer: nil,
			Middleware: nil,
		},
		input,
		[]session_validation.Validation{
			func(user *model.AuthenticatedUser) error {
				return nil
			},
			func(user *model.AuthenticatedUser) error {
				return nil
			},
			func(user *model.AuthenticatedUser) error {
				finalFuncCalled = true
				return nil
			},
		})

	input <- expectedUser

	err := command.Run()

	assert.Nil(t, err)
	assert.True(t, finalFuncCalled)
	assert.Equal(t, expectedUser, <-command.Stream)
}
