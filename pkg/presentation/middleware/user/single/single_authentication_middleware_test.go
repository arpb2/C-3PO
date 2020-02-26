package single_test

import (
	"github.com/arpb2/C-3PO/pkg/presentation/user"
	"net/http"
	"testing"

	http3 "github.com/arpb2/C-3PO/pkg/domain/architecture/http"
	token2 "github.com/arpb2/C-3PO/pkg/domain/session/token"
	http2 "github.com/arpb2/C-3PO/test/mock/http"
	"github.com/arpb2/C-3PO/test/mock/token"

	"github.com/arpb2/C-3PO/pkg/presentation/middleware"

	"github.com/arpb2/C-3PO/pkg/presentation/middleware/user/single"
	"github.com/stretchr/testify/assert"
)

func Test_Single_HandlingOfAuthentication_NoHeader(t *testing.T) {
	reader := new(http2.MockReader)
	reader.On("GetHeader", middleware.HeaderAuthorization).Return("")

	c, w := http2.CreateTestContext()
	c.Reader = reader

	middle := single.CreateMiddleware(&token.MockTokenHandler{})

	middle(c)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Equal(t, "{\"error\":\"unauthorized\"}", w.Body.String())
}

func Test_Single_HandlingOfAuthentication_BadHeader(t *testing.T) {
	tokenHandler := new(token.MockTokenHandler)
	tokenHandler.On("Retrieve", "token").Return(nil, http3.CreateBadRequestError("malformed token"))

	reader := new(http2.MockReader)
	reader.On("GetHeader", middleware.HeaderAuthorization).Return("token")

	c, recorder := http2.CreateTestContext()
	c.Reader = reader

	middle := single.CreateMiddleware(tokenHandler)

	middle(c)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.Equal(t, "{\"error\":\"malformed token\"}", recorder.Body.String())
	reader.AssertExpectations(t)
	tokenHandler.AssertExpectations(t)
}

func Test_Single_HandlingOfAuthentication_UnauthorizedUser(t *testing.T) {
	tokenHandler := new(token.MockTokenHandler)
	tokenHandler.On("Retrieve", "token").Return(&token2.Token{
		UserId: 1000,
	}, nil)

	reader := new(http2.MockReader)
	reader.On("GetParameter", user.ParamUserId).Return("1", nil).Once()
	reader.On("GetHeader", middleware.HeaderAuthorization).Return("token")

	c, recorder := http2.CreateTestContext()
	c.Reader = reader

	middle := single.CreateMiddleware(tokenHandler)

	middle(c)

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	assert.Equal(t, "{\"error\":\"unauthorized\"}", recorder.Body.String())
	reader.AssertExpectations(t)
	tokenHandler.AssertExpectations(t)
}

func Test_Single_HandlingOfAuthentication_Authorized_SameUser(t *testing.T) {
	tokenHandler := new(token.MockTokenHandler)
	tokenHandler.On("Retrieve", "token").Return(&token2.Token{
		UserId: 1000,
	}, nil)

	reader := new(http2.MockReader)
	reader.On("GetParameter", user.ParamUserId).Return("1000", nil).Once()
	reader.On("GetHeader", middleware.HeaderAuthorization).Return("token")

	c, recorder := http2.CreateTestContext()
	c.Reader = reader

	middle := single.CreateMiddleware(tokenHandler)

	middle(c)

	assert.Equal(t, http.StatusOK, recorder.Code)
	reader.AssertExpectations(t)
	tokenHandler.AssertExpectations(t)
}
