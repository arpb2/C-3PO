package middleware_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/arpb2/C-3PO/pkg/presentation/auth/middleware"

	internalauth "github.com/arpb2/C-3PO/pkg/data/jwt"
	"github.com/arpb2/C-3PO/pkg/domain/auth"
	"github.com/arpb2/C-3PO/pkg/domain/controller"
	httpwrapper "github.com/arpb2/C-3PO/pkg/domain/http"
	ginengine "github.com/arpb2/C-3PO/pkg/infra/engine/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAuthenticationStrategy struct {
	mock.Mock
}

func (s MockAuthenticationStrategy) Authenticate(token *auth.Token, userId string) (authorized bool, err error) {
	args := s.Called(token, userId)
	return args.Bool(0), args.Error(1)
}

func performRequest(r http.Handler, method, path, body string, headers map[string][]string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, strings.NewReader(body))
	req.Header = headers

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func Test_HandlingOfAuthentication_NoHeader(t *testing.T) {
	e := ginengine.CreateEngine()
	e.Register(controller.Controller{
		Method: "GET",
		Path:   "/test",
		Middleware: []httpwrapper.Handler{
			func(context *httpwrapper.Context) {
				middleware.HandleAuthentication(context, internalauth.CreateTokenHandler())
			},
		},
		Body: func(ctx *httpwrapper.Context) {
			panic("Shouldn't reach here!")
		},
	})

	recorder := performRequest(e, "GET", "/test", "", map[string][]string{})

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	assert.Equal(t, "{\"error\":\"unauthorized\"}\n", recorder.Body.String())
}

func Test_HandlingOfAuthentication_BadHeader(t *testing.T) {
	e := ginengine.CreateEngine()
	e.Register(controller.Controller{
		Method: "GET",
		Path:   "/test",
		Middleware: []httpwrapper.Handler{
			func(context *httpwrapper.Context) {
				middleware.HandleAuthentication(context, internalauth.CreateTokenHandler())
			},
		},
		Body: func(ctx *httpwrapper.Context) {
			panic("Shouldn't reach here!")
		},
	})

	headers := map[string][]string{}
	headers["Authorization"] = []string{"bad token"}
	recorder := performRequest(e, "GET", "/test", "", headers)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.Equal(t, "{\"error\":\"malformed token\"}\n", recorder.Body.String())
}

func Test_HandlingOfAuthentication_UnauthorizedUser(t *testing.T) {
	e := ginengine.CreateEngine()
	e.Register(controller.Controller{
		Method: "GET",
		Path:   "/test/:user_id",
		Middleware: []httpwrapper.Handler{
			func(context *httpwrapper.Context) {
				middleware.HandleAuthentication(context, internalauth.CreateTokenHandler())
			},
		},
		Body: func(ctx *httpwrapper.Context) {
			panic("Shouldn't reach here!")
		},
	})

	headers := map[string][]string{}
	// Token for user 1000
	headers["Authorization"] = []string{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEwMDB9.GVS-KC5nOCHybzzFIIH864u4KcGu-ZSd-96krqTUGWo"}
	recorder := performRequest(e, "GET", "/test/1", "", headers)

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	assert.Equal(t, "{\"error\":\"unauthorized\"}\n", recorder.Body.String())
}

func Test_HandlingOfAuthentication_Authorized_SameUser(t *testing.T) {
	e := ginengine.CreateEngine()
	e.Register(controller.Controller{
		Method: "GET",
		Path:   "/test/:user_id",
		Middleware: []httpwrapper.Handler{
			func(context *httpwrapper.Context) {
				middleware.HandleAuthentication(context, internalauth.CreateTokenHandler())
			},
		},
		Body: func(ctx *httpwrapper.Context) {
			ctx.WriteString(http.StatusOK, "Returned success")
		},
	})

	headers := map[string][]string{}
	// Token for user 1000
	headers["Authorization"] = []string{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEwMDB9.GVS-KC5nOCHybzzFIIH864u4KcGu-ZSd-96krqTUGWo"}
	recorder := performRequest(e, "GET", "/test/1000", "", headers)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "Returned success", recorder.Body.String())
}

func TestStrategy_Error_Halts(t *testing.T) {
	strategy := new(MockAuthenticationStrategy)
	strategy.On("Authenticate", mock.MatchedBy(func(token *auth.Token) bool {
		return token.UserId == uint(1000)
	}), "1001").Return(false, errors.New("whoops this fails")).Once()

	e := ginengine.CreateEngine()
	e.Register(controller.Controller{
		Method: "GET",
		Path:   "/test/:user_id",
		Middleware: []httpwrapper.Handler{
			func(context *httpwrapper.Context) {
				middleware.HandleAuthentication(context, internalauth.CreateTokenHandler(), strategy)
			},
		},
		Body: func(ctx *httpwrapper.Context) {
			ctx.WriteString(http.StatusOK, "Returned success")
		},
	})

	headers := map[string][]string{}
	// Token for user 1000
	headers["Authorization"] = []string{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEwMDB9.GVS-KC5nOCHybzzFIIH864u4KcGu-ZSd-96krqTUGWo"}
	recorder := performRequest(e, "GET", "/test/1001", "", headers)

	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	assert.Equal(t, "{\"error\":\"internal error\"}\n", recorder.Body.String())
	strategy.AssertExpectations(t)
}

func TestStrategy_Unauthorized_Halts(t *testing.T) {
	strategy := new(MockAuthenticationStrategy)
	strategy.On("Authenticate", mock.MatchedBy(func(token *auth.Token) bool {
		return token.UserId == uint(1000)
	}), "1001").Return(false, nil).Once()

	e := ginengine.CreateEngine()
	e.Register(controller.Controller{
		Method: "GET",
		Path:   "/test/:user_id",
		Middleware: []httpwrapper.Handler{
			func(context *httpwrapper.Context) {
				middleware.HandleAuthentication(context, internalauth.CreateTokenHandler(), strategy)
			},
		},
		Body: func(ctx *httpwrapper.Context) {
			ctx.WriteString(http.StatusOK, "Returned success")
		},
	})

	headers := map[string][]string{}
	// Token for user 1000
	headers["Authorization"] = []string{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEwMDB9.GVS-KC5nOCHybzzFIIH864u4KcGu-ZSd-96krqTUGWo"}
	recorder := performRequest(e, "GET", "/test/1001", "", headers)

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	assert.Equal(t, "{\"error\":\"unauthorized\"}\n", recorder.Body.String())
	strategy.AssertExpectations(t)
}

func TestStrategy_Authorized_Continues(t *testing.T) {
	strategy := new(MockAuthenticationStrategy)
	strategy.On("Authenticate", mock.MatchedBy(func(token *auth.Token) bool {
		return token.UserId == uint(1000)
	}), "1001").Return(true, nil).Once()

	e := ginengine.CreateEngine()
	e.Register(controller.Controller{
		Method: "GET",
		Path:   "/test/:user_id",
		Middleware: []httpwrapper.Handler{
			func(context *httpwrapper.Context) {
				middleware.HandleAuthentication(context, internalauth.CreateTokenHandler(), strategy)
			},
		},
		Body: func(ctx *httpwrapper.Context) {
			ctx.WriteString(http.StatusOK, "Returned success")
		},
	})

	headers := map[string][]string{}
	// Token for user 1000
	headers["Authorization"] = []string{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEwMDB9.GVS-KC5nOCHybzzFIIH864u4KcGu-ZSd-96krqTUGWo"}
	recorder := performRequest(e, "GET", "/test/1001", "", headers)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "Returned success", recorder.Body.String())
	strategy.AssertExpectations(t)
}
