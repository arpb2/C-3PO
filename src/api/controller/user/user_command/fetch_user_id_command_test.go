package user_command_test

import (
	"errors"
	"github.com/arpb2/C-3PO/src/api/controller/user/user_command"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestFetchUserIdCommand_Name(t *testing.T) {
	assert.Equal(t, "fetch_user_id_command", user_command.CreateFetchUserIdCommand(nil).Name())
}

func TestFetchUserIdCommand_Fallback_DoesNothingOnUnknownError(t *testing.T) {
	command := user_command.CreateFetchUserIdCommand(nil)
	runErr := errors.New("run err")

	assert.Equal(t, runErr, command.Fallback(runErr))
}

func TestFetchUserIdCommand_Fallback_ReturnsNilAndNotifies_OnBadRequest(t *testing.T) {
	middleware := new(http_wrapper.MockMiddleware)
	middleware.On("AbortTransactionWithStatus", http.StatusBadRequest, http_wrapper.Json{
		"error": "some message",
	})

	command := user_command.CreateFetchUserIdCommand(&http_wrapper.Context{
		Reader: nil,
		Writer: nil,
		Middleware: middleware,
	})

	runErr := http_wrapper.CreateBadRequestError("some message")

	assert.Nil(t, command.Fallback(runErr))
	middleware.AssertExpectations(t)
}

func TestFetchUserId_RetrievesFromParam(t *testing.T) {
	reader := new(http_wrapper.MockReader)
	reader.On("GetParameter", "user_id").Return("1234").Once()

	fetchCommand := user_command.CreateFetchUserIdCommand(&http_wrapper.Context{
		Reader:     reader,
		Writer:     nil,
		Middleware: nil,
	})

	err := fetchCommand.Run()

	assert.Nil(t, err)
	assert.Equal(t, uint(1234), <-fetchCommand.OutputStream)
	reader.AssertExpectations(t)
}

func TestFetchUserId_RetrievesFromParam_400IfMalformed(t *testing.T) {
	reader := new(http_wrapper.MockReader)
	reader.On("GetParameter", "user_id").Return("not a number").Once()

	fetchCommand := user_command.CreateFetchUserIdCommand(&http_wrapper.Context{
		Reader:     reader,
		Writer:     nil,
		Middleware: nil,
	})

	err := fetchCommand.Run()

	assert.Equal(t, http_wrapper.CreateBadRequestError("'user_id' malformed, expecting a positive number"), err)
	assert.Zero(t, len(fetchCommand.OutputStream))
	reader.AssertExpectations(t)
}

func TestFetchUserId_HaltsWith400OnError(t *testing.T) {
	reader := new(http_wrapper.MockReader)
	reader.On("GetParameter", "user_id").Return("").Once()

	fetchCommand := user_command.CreateFetchUserIdCommand(&http_wrapper.Context{
		Reader:     reader,
		Writer:     nil,
		Middleware: nil,
	})

	err := fetchCommand.Run()

	assert.Equal(t, http_wrapper.CreateBadRequestError("'user_id' empty"), err)
	assert.Zero(t, len(fetchCommand.OutputStream))
	reader.AssertExpectations(t)
}

