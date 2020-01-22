package session_command_test

import (
	"errors"
	"github.com/arpb2/C-3PO/src/api/controller/session/session_command"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/model"
	"github.com/arpb2/C-3PO/src/api/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthenticateCommand_Name(t *testing.T) {
	assert.Equal(t, "authenticate_command", session_command.CreateAuthenticateCommand(nil, nil, nil).Name())
}

func TestAuthenticateCommand_Fallback_DoesNothing(t *testing.T) {
	command := session_command.CreateAuthenticateCommand(nil, nil, nil)
	runErr := errors.New("run err")

	assert.Equal(t, runErr, command.Fallback(runErr))
}

func TestAuthenticateCommand_Run_OnServiceError_ReturnsError(t *testing.T) {
	expectedErr := errors.New("service error")
	expectedEmail := "test@email.com"
	expectedPassword := "test password"

	input := make(chan *model.AuthenticatedUser, 1)
	credentialService := new(service.MockCredentialService)
	credentialService.On("Retrieve", expectedEmail, expectedPassword).Return(uint(0), expectedErr)

	command := session_command.CreateAuthenticateCommand(
		&http_wrapper.Context{
			Reader: nil,
			Writer: nil,
			Middleware: nil,
		},
		credentialService,
		input)

	input <- &model.AuthenticatedUser{
		User: &model.User{
			Email: expectedEmail,
		},
		Password: expectedPassword,
	}

	err := command.Run()

	assert.Equal(t, err, expectedErr)
	assert.Zero(t, len(command.Stream))
	credentialService.AssertExpectations(t)
}

func TestAuthenticateCommand_Run_OnServiceSuccess_PublishesDecoratedUser(t *testing.T) {
	expectedUserId := uint(1000)
	expectedEmail := "test@email.com"
	expectedPassword := "test password"

	input := make(chan *model.AuthenticatedUser, 1)
	credentialService := new(service.MockCredentialService)
	credentialService.On("Retrieve", expectedEmail, expectedPassword).Return(expectedUserId, nil)

	command := session_command.CreateAuthenticateCommand(
		&http_wrapper.Context{
			Reader: nil,
			Writer: nil,
			Middleware: nil,
		},
		credentialService,
		input)

	input <- &model.AuthenticatedUser{
		User: &model.User{
			Email: expectedEmail,
		},
		Password: expectedPassword,
	}

	err := command.Run()

	assert.Nil(t, err)

	user := <-input
	assert.Equal(t, expectedUserId, user.Id)
	assert.Equal(t, expectedEmail, user.Email)
	assert.Equal(t, expectedPassword, user.Password)
	credentialService.AssertExpectations(t)
}
