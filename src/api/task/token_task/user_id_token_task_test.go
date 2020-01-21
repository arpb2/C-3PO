package token_task_test

import (
	"errors"
	"github.com/arpb2/C-3PO/src/api/auth"
	"github.com/arpb2/C-3PO/src/api/task/token_task"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

type MockTokenHandler struct{
	mock.Mock
}

func (m MockTokenHandler) Create(token auth.Token) (*string, *auth.TokenError) {
	args := m.Called(token)

	var tokenString *string
	if args.Get(0) != nil {
		tokenString = args.Get(0).(*string)
	}

	var err *auth.TokenError
	if args.Get(1) != nil {
		err = args.Get(1).(*auth.TokenError)
	}

	return tokenString, err
}

func (m MockTokenHandler) Retrieve(token string) (*auth.Token, *auth.TokenError) {
	args := m.Called(token)

	var tokenString *auth.Token
	if args.Get(0) != nil {
		tokenString = args.Get(0).(*auth.Token)
	}

	var err *auth.TokenError
	if args.Get(1) != nil {
		err = args.Get(1).(*auth.TokenError)
	}

	return tokenString, err
}

func TestCreateTokenTask_NilOnError(t *testing.T) {
	expectedToken := auth.Token{
		UserId: 1000,
	}

	tokenHandler := new(MockTokenHandler)
	tokenHandler.On("Create", expectedToken).Return(nil, &auth.TokenError{
		Error:  errors.New("test error"),
		Status: http.StatusInternalServerError,
	})

	token, err := token_task.CreateTokenTaskImpl(
		1000,
		tokenHandler,
	)

	assert.Nil(t, token)
	assert.NotNil(t, err)
	tokenHandler.AssertExpectations(t)
}

func TestCreateTokenTask_TokenOnSuccess(t *testing.T) {
	expectedToken := auth.Token{
		UserId: 1000,
	}

	expectedTokenStr := "token string"

	tokenHandler := new(MockTokenHandler)
	tokenHandler.On("Create", expectedToken).Return(&expectedTokenStr, nil)

	token, err := token_task.CreateTokenTaskImpl(
		1000,
		tokenHandler,
	)

	assert.Equal(t, &expectedTokenStr, token)
	assert.Nil(t, err)
	tokenHandler.AssertExpectations(t)
}