package user_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/arpb2/C-3PO/pkg/presentation/middleware/auth"

	"github.com/arpb2/C-3PO/pkg/domain/session/repository"

	user2 "github.com/arpb2/C-3PO/pkg/presentation/user"

	http3 "github.com/arpb2/C-3PO/pkg/domain/architecture/http"
	http2 "github.com/arpb2/C-3PO/test/mock/http"
	mocktoken "github.com/arpb2/C-3PO/test/mock/token"

	"github.com/arpb2/C-3PO/pkg/presentation/middleware/auth/user"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAuthenticationStrategy struct {
	mock.Mock
}

func (s *MockAuthenticationStrategy) Authenticate(token *repository.Token, userId string) (authorized bool, err error) {
	args := s.Called(token, userId)
	return args.Bool(0), args.Error(1)
}

func Test_HandlingOfAuthentication_NoHeader(t *testing.T) {
	reader := new(http2.MockReader)
	reader.On("GetHeader", auth.HeaderAuthorization).Return("")

	c, recorder := http2.CreateTestContext()
	c.Reader = reader

	user.HandleAuthentication(c, &mocktoken.MockTokenHandler{})

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	assert.Equal(t, "{\"error\":\"unauthorized\"}", recorder.Body.String())
	reader.AssertExpectations(t)
}

func Test_HandlingOfAuthentication_BadHeader(t *testing.T) {
	tokenHandler := new(mocktoken.MockTokenHandler)
	tokenHandler.On("Retrieve", "token").Return(nil, http3.CreateBadRequestError("malformed token"))

	reader := new(http2.MockReader)
	reader.On("GetHeader", auth.HeaderAuthorization).Return("token")

	c, recorder := http2.CreateTestContext()
	c.Reader = reader

	user.HandleAuthentication(c, tokenHandler)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.Equal(t, "{\"error\":\"malformed token\"}", recorder.Body.String())
	reader.AssertExpectations(t)
	tokenHandler.AssertExpectations(t)
}

func Test_HandlingOfAuthentication_UnauthorizedUser(t *testing.T) {
	tokenHandler := new(mocktoken.MockTokenHandler)
	tokenHandler.On("Retrieve", "token").Return(&repository.Token{
		UserId: 1000,
	}, nil)

	reader := new(http2.MockReader)
	reader.On("GetParameter", user2.ParamUserId).Return("1", nil).Once()
	reader.On("GetHeader", auth.HeaderAuthorization).Return("token")

	c, recorder := http2.CreateTestContext()
	c.Reader = reader

	user.HandleAuthentication(c, tokenHandler)

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	assert.Equal(t, "{\"error\":\"unauthorized\"}", recorder.Body.String())
	reader.AssertExpectations(t)
	tokenHandler.AssertExpectations(t)
}

func Test_HandlingOfAuthentication_Authorized_SameUser(t *testing.T) {
	tokenHandler := new(mocktoken.MockTokenHandler)
	tokenHandler.On("Retrieve", "token").Return(&repository.Token{
		UserId: 1000,
	}, nil)

	reader := new(http2.MockReader)
	reader.On("GetParameter", user2.ParamUserId).Return("1000", nil).Once()
	reader.On("GetHeader", auth.HeaderAuthorization).Return("token")

	c, recorder := http2.CreateTestContext()
	c.Reader = reader

	user.HandleAuthentication(c, tokenHandler)

	assert.Equal(t, http.StatusOK, recorder.Code)
	reader.AssertExpectations(t)
	tokenHandler.AssertExpectations(t)
}

func TestStrategy_Error_Halts(t *testing.T) {
	tokenHandler := new(mocktoken.MockTokenHandler)
	tokenHandler.On("Retrieve", "token").Return(&repository.Token{
		UserId: 1000,
	}, nil)

	strategy := new(MockAuthenticationStrategy)
	strategy.On("Authenticate", mock.MatchedBy(func(token *repository.Token) bool {
		return token.UserId == uint(1000)
	}), "1001").Return(false, errors.New("whoops this fails")).Once()

	reader := new(http2.MockReader)
	reader.On("GetParameter", user2.ParamUserId).Return("1001", nil).Once()
	reader.On("GetHeader", auth.HeaderAuthorization).Return("token")

	c, recorder := http2.CreateTestContext()
	c.Reader = reader

	user.HandleAuthentication(c, tokenHandler, strategy)

	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	assert.Equal(t, "{\"error\":\"internal error\"}", recorder.Body.String())
	strategy.AssertExpectations(t)
	reader.AssertExpectations(t)
	tokenHandler.AssertExpectations(t)
}

func TestStrategy_Unauthorized_Halts(t *testing.T) {
	tokenHandler := new(mocktoken.MockTokenHandler)
	tokenHandler.On("Retrieve", "token").Return(&repository.Token{
		UserId: 1000,
	}, nil)

	strategy := new(MockAuthenticationStrategy)
	strategy.On("Authenticate", mock.MatchedBy(func(token *repository.Token) bool {
		return token.UserId == uint(1000)
	}), "1001").Return(false, nil).Once()

	reader := new(http2.MockReader)
	reader.On("GetParameter", user2.ParamUserId).Return("1001", nil).Once()
	reader.On("GetHeader", auth.HeaderAuthorization).Return("token")

	c, recorder := http2.CreateTestContext()
	c.Reader = reader

	user.HandleAuthentication(c, tokenHandler, strategy)

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	assert.Equal(t, "{\"error\":\"unauthorized\"}", recorder.Body.String())
	strategy.AssertExpectations(t)
	reader.AssertExpectations(t)
	tokenHandler.AssertExpectations(t)
}

func TestStrategy_Authorized_Continues(t *testing.T) {
	tokenHandler := new(mocktoken.MockTokenHandler)
	tokenHandler.On("Retrieve", "token").Return(&repository.Token{
		UserId: 1000,
	}, nil)

	strategy := new(MockAuthenticationStrategy)
	strategy.On("Authenticate", mock.MatchedBy(func(token *repository.Token) bool {
		return token.UserId == uint(1000)
	}), "1001").Return(true, nil).Once()

	reader := new(http2.MockReader)
	reader.On("GetParameter", user2.ParamUserId).Return("1001", nil).Once()
	reader.On("GetHeader", auth.HeaderAuthorization).Return("token")

	c, recorder := http2.CreateTestContext()
	c.Reader = reader

	user.HandleAuthentication(c, tokenHandler, strategy)

	assert.Equal(t, http.StatusOK, recorder.Code)
	strategy.AssertExpectations(t)
	reader.AssertExpectations(t)
	tokenHandler.AssertExpectations(t)
}
