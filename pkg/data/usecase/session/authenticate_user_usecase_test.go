package session_test

import (
	"testing"

	session2 "github.com/arpb2/C-3PO/pkg/data/repository/session"
	"github.com/arpb2/C-3PO/pkg/data/usecase/session"

	http3 "github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/test/mock/token"

	"github.com/stretchr/testify/assert"
)

func Test_Single_HandlingOfAuthentication_NoHeader(t *testing.T) {
	middle := session.CreateUserAuthenticationUseCase(&token.MockTokenHandler{})

	err := middle("", "")

	assert.Equal(t, http3.CreateUnauthorizedError(), err)
}

func Test_Single_HandlingOfAuthentication_BadHeader(t *testing.T) {
	tokenHandler := new(token.MockTokenHandler)
	tokenHandler.On("Retrieve", "token").Return(nil, http3.CreateBadRequestError("malformed token"))

	middle := session.CreateUserAuthenticationUseCase(tokenHandler)

	err := middle("token", "")

	assert.Equal(t, http3.CreateBadRequestError("malformed token"), err)
	tokenHandler.AssertExpectations(t)
}

func Test_Single_HandlingOfAuthentication_UnauthorizedUser(t *testing.T) {
	tokenHandler := new(token.MockTokenHandler)
	tokenHandler.On("Retrieve", "token").Return(&session2.Token{
		UserId: 1000,
	}, nil)

	middle := session.CreateUserAuthenticationUseCase(tokenHandler)

	err := middle("token", "1")

	assert.Equal(t, http3.CreateUnauthorizedError(), err)
	tokenHandler.AssertExpectations(t)
}

func Test_Single_HandlingOfAuthentication_Authorized_SameUser(t *testing.T) {
	tokenHandler := new(token.MockTokenHandler)
	tokenHandler.On("Retrieve", "token").Return(&session2.Token{
		UserId: 1000,
	}, nil)

	middle := session.CreateUserAuthenticationUseCase(tokenHandler)

	err := middle("token", "1000")

	assert.Nil(t, err)
	tokenHandler.AssertExpectations(t)
}
