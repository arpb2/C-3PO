package single_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/arpb2/C-3PO/pkg/presentation/middleware"

	"github.com/arpb2/C-3PO/pkg/data/jwt"
	"github.com/arpb2/C-3PO/pkg/domain/controller"
	httpwrapper "github.com/arpb2/C-3PO/pkg/domain/http"
	ginengine "github.com/arpb2/C-3PO/pkg/infra/engine/gin"
	"github.com/arpb2/C-3PO/pkg/presentation/middleware/user/single"
	"github.com/stretchr/testify/assert"
)

var SingleTokenHandler = jwt.CreateTokenHandler([]byte("52bfd2de0a2e69dff4517518590ac32a46bd76606ec22a258f99584a6e70aca2"))

func performRequest(r http.Handler, method, path, body string, headers map[string][]string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, strings.NewReader(body))
	req.Header = headers

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func Test_Single_HandlingOfAuthentication_NoHeader(t *testing.T) {
	e := ginengine.CreateEngine("8080")
	e.Register(controller.Controller{
		Method: "GET",
		Path:   "/test",
		Middleware: []httpwrapper.Handler{
			single.CreateMiddleware(SingleTokenHandler),
		},
		Body: func(ctx *httpwrapper.Context) {
			panic("Shouldn't reach here!")
		},
	})

	recorder := performRequest(e, "GET", "/test", "", map[string][]string{})

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	assert.Equal(t, "{\"error\":\"unauthorized\"}\n", recorder.Body.String())
}

func Test_Single_HandlingOfAuthentication_BadHeader(t *testing.T) {
	e := ginengine.CreateEngine("8080")
	e.Register(controller.Controller{
		Method: "GET",
		Path:   "/test",
		Middleware: []httpwrapper.Handler{
			single.CreateMiddleware(SingleTokenHandler),
		},
		Body: func(ctx *httpwrapper.Context) {
			panic("Shouldn't reach here!")
		},
	})

	headers := map[string][]string{}
	headers[middleware.HeaderAuthorization] = []string{"bad token"}
	recorder := performRequest(e, "GET", "/test", "", headers)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.Equal(t, "{\"error\":\"malformed token\"}\n", recorder.Body.String())
}

func Test_Single_HandlingOfAuthentication_UnauthorizedUser(t *testing.T) {
	e := ginengine.CreateEngine("8080")
	e.Register(controller.Controller{
		Method: "GET",
		Path:   "/test/:user_id",
		Middleware: []httpwrapper.Handler{
			single.CreateMiddleware(SingleTokenHandler),
		},
		Body: func(ctx *httpwrapper.Context) {
			panic("Shouldn't reach here!")
		},
	})

	headers := map[string][]string{}
	// Token for user 1000
	headers[middleware.HeaderAuthorization] = []string{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEwMDB9.GVS-KC5nOCHybzzFIIH864u4KcGu-ZSd-96krqTUGWo"}
	recorder := performRequest(e, "GET", "/test/1", "", headers)

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	assert.Equal(t, "{\"error\":\"unauthorized\"}\n", recorder.Body.String())
}

func Test_Single_HandlingOfAuthentication_Authorized_SameUser(t *testing.T) {
	e := ginengine.CreateEngine("8080")
	e.Register(controller.Controller{
		Method: "GET",
		Path:   "/test/:user_id",
		Middleware: []httpwrapper.Handler{
			single.CreateMiddleware(SingleTokenHandler),
		},
		Body: func(ctx *httpwrapper.Context) {
			ctx.WriteString(http.StatusOK, "Returned success")
		},
	})

	headers := map[string][]string{}
	// Token for user 1000
	headers[middleware.HeaderAuthorization] = []string{"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEwMDB9.GVS-KC5nOCHybzzFIIH864u4KcGu-ZSd-96krqTUGWo"}
	recorder := performRequest(e, "GET", "/test/1000", "", headers)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "Returned success", recorder.Body.String())
}
