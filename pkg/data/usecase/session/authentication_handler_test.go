package session_test

import (
	"errors"
	"testing"

	"github.com/arpb2/C-3PO/pkg/data/repository/session"
	session2 "github.com/arpb2/C-3PO/pkg/data/usecase/session"

	http3 "github.com/arpb2/C-3PO/pkg/domain/http"
	mocktoken "github.com/arpb2/C-3PO/test/mock/token"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAuthenticationStrategy struct {
	mock.Mock
}

func (s *MockAuthenticationStrategy) Authenticate(token *session.Token, userId string) (authorized bool, err error) {
	args := s.Called(token, userId)
	return args.Bool(0), args.Error(1)
}

func Test_HandlingOfAuthentication_NoHeader(t *testing.T) {
	err := session2.HandleTokenizedAuthentication("", "", &mocktoken.MockTokenHandler{})

	assert.Equal(t, http3.CreateUnauthorizedError(), err)
}

func Test_HandlingOfAuthentication_BadHeader(t *testing.T) {
	tokenHandler := new(mocktoken.MockTokenHandler)
	tokenHandler.On("Retrieve", "token").Return(nil, http3.CreateBadRequestError("malformed token"))

	err := session2.HandleTokenizedAuthentication("token", "", tokenHandler)

	assert.Equal(t, http3.CreateBadRequestError("malformed token"), err)
	tokenHandler.AssertExpectations(t)
}

func Test_HandlingOfAuthentication_UnauthorizedUser(t *testing.T) {
	tokenHandler := new(mocktoken.MockTokenHandler)
	tokenHandler.On("Retrieve", "token").Return(&session.Token{
		UserId: 1000,
	}, nil)

	err := session2.HandleTokenizedAuthentication("token", "1", tokenHandler)

	assert.Equal(t, http3.CreateUnauthorizedError(), err)
	tokenHandler.AssertExpectations(t)
}

func Test_HandlingOfAuthentication_Authorized_SameUser(t *testing.T) {
	tokenHandler := new(mocktoken.MockTokenHandler)
	tokenHandler.On("Retrieve", "token").Return(&session.Token{
		UserId: 1000,
	}, nil)

	err := session2.HandleTokenizedAuthentication("token", "1000", tokenHandler)

	assert.Nil(t, err)
	tokenHandler.AssertExpectations(t)
}

func TestStrategy_Error_Halts(t *testing.T) {
	tokenHandler := new(mocktoken.MockTokenHandler)
	tokenHandler.On("Retrieve", "token").Return(&session.Token{
		UserId: 1000,
	}, nil)

	expectedErr := errors.New("whoops this fails")
	strategy := new(MockAuthenticationStrategy)
	strategy.On("Authenticate", mock.MatchedBy(func(token *session.Token) bool {
		return token.UserId == uint(1000)
	}), "1001").Return(false, expectedErr).Once()

	err := session2.HandleTokenizedAuthentication("token", "1001", tokenHandler, strategy)

	assert.Equal(t, expectedErr, err)
	strategy.AssertExpectations(t)
	tokenHandler.AssertExpectations(t)
}

func TestStrategy_Unauthorized_Halts(t *testing.T) {
	tokenHandler := new(mocktoken.MockTokenHandler)
	tokenHandler.On("Retrieve", "token").Return(&session.Token{
		UserId: 1000,
	}, nil)

	strategy := new(MockAuthenticationStrategy)
	strategy.On("Authenticate", mock.MatchedBy(func(token *session.Token) bool {
		return token.UserId == uint(1000)
	}), "1001").Return(false, nil).Once()

	err := session2.HandleTokenizedAuthentication("token", "1001", tokenHandler, strategy)

	assert.Equal(t, http3.CreateUnauthorizedError(), err)
	strategy.AssertExpectations(t)
	tokenHandler.AssertExpectations(t)
}

func TestStrategy_Authorized_Continues(t *testing.T) {
	tokenHandler := new(mocktoken.MockTokenHandler)
	tokenHandler.On("Retrieve", "token").Return(&session.Token{
		UserId: 1000,
	}, nil)

	strategy := new(MockAuthenticationStrategy)
	strategy.On("Authenticate", mock.MatchedBy(func(token *session.Token) bool {
		return token.UserId == uint(1000)
	}), "1001").Return(true, nil).Once()

	err := session2.HandleTokenizedAuthentication("token", "1001", tokenHandler, strategy)

	assert.Nil(t, err)
	strategy.AssertExpectations(t)
	tokenHandler.AssertExpectations(t)
}
