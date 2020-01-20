package single_auth_test

import (
	"github.com/arpb2/C-3PO/src/api/auth/jwt"
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/engine/gin_engine"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/middleware/auth/single_auth"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var SingleTokenHandler = jwt.CreateTokenHandler()

func performRequest(r http.Handler, method, path, body string, headers map[string][]string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, strings.NewReader(body))
	req.Header = headers

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func Test_Single_HandlingOfAuthentication_NoHeader(t *testing.T) {
	e := gin_engine.New()
	e.Register(controller.Controller{
		Method:        "GET",
		Path:          "/test",
		Middleware:    []http_wrapper.Handler{
			single_auth.CreateMiddleware(SingleTokenHandler),
		},
		Body:          func(ctx *http_wrapper.Context) {
			panic("Shouldn't reach here!")
		},
	})

	recorder := performRequest(e, "GET", "/test", "", map[string][]string{})

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	assert.Equal(t, "{\"error\":\"no 'Authorization' header provided\"}\n", recorder.Body.String())
}

func Test_Single_HandlingOfAuthentication_BadHeader(t *testing.T) {
	e := gin_engine.New()
	e.Register(controller.Controller{
		Method:        "GET",
		Path:          "/test",
		Middleware:    []http_wrapper.Handler{
			single_auth.CreateMiddleware(SingleTokenHandler),
		},
		Body:          func(ctx *http_wrapper.Context) {
			panic("Shouldn't reach here!")
		},
	})

	headers := map[string][]string{}
	headers["Authorization"] = []string{"bad token"}
	recorder := performRequest(e, "GET", "/test", "", headers)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.Equal(t, "{\"error\":\"malformed token\"}\n", recorder.Body.String())
}

func Test_Single_HandlingOfAuthentication_UnauthorizedUser(t *testing.T) {
	e := gin_engine.New()
	e.Register(controller.Controller{
		Method:        "GET",
		Path:          "/test/:user_id",
		Middleware:    []http_wrapper.Handler{
			single_auth.CreateMiddleware(SingleTokenHandler),
		},
		Body:          func(ctx *http_wrapper.Context) {
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

func Test_Single_HandlingOfAuthentication_Authorized_SameUser(t *testing.T) {
	e := gin_engine.New()
	e.Register(controller.Controller{
		Method:        "GET",
		Path:          "/test/:user_id",
		Middleware:    []http_wrapper.Handler{
			single_auth.CreateMiddleware(SingleTokenHandler),
		},
		Body:          func(ctx *http_wrapper.Context) {
			ctx.String(http.StatusOK, "Returned success")
		},
	})

	headers := map[string][]string{}
	// Token for user 1000
	headers["Authorization"] = []string{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEwMDB9.GVS-KC5nOCHybzzFIIH864u4KcGu-ZSd-96krqTUGWo"}
	recorder := performRequest(e, "GET", "/test/1000", "", headers)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "Returned success", recorder.Body.String())
}