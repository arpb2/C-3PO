package auth_test

import (
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/engine"
	"github.com/arpb2/C-3PO/src/api/middleware/auth"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_Single_HandlingOfAuthentication_NoHeader(t *testing.T) {
	e := engine.CreateBasicServerEngine()
	e.Register(controller.Controller{
		Method:        "GET",
		Path:          "/test",
		Middleware:    []gin.HandlerFunc{
			auth.SingleAuthenticationMiddleware,
		},
		Body:          func(ctx *gin.Context) {
			panic("Shouldn't reach here!")
		},
	})

	recorder := performRequest(e, "GET", "/test", "", map[string][]string{})

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	assert.Equal(t, "{\"error\":\"no 'Authorization' header provided\"}\n", recorder.Body.String())
}

func Test_Single_HandlingOfAuthentication_BadHeader(t *testing.T) {
	e := engine.CreateBasicServerEngine()
	e.Register(controller.Controller{
		Method:        "GET",
		Path:          "/test",
		Middleware:    []gin.HandlerFunc{
			auth.SingleAuthenticationMiddleware,
		},
		Body:          func(ctx *gin.Context) {
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
	e := engine.CreateBasicServerEngine()
	e.Register(controller.Controller{
		Method:        "GET",
		Path:          "/test/:user_id",
		Middleware:    []gin.HandlerFunc{
			auth.SingleAuthenticationMiddleware,
		},
		Body:          func(ctx *gin.Context) {
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
	e := engine.CreateBasicServerEngine()
	e.Register(controller.Controller{
		Method:        "GET",
		Path:          "/test/:user_id",
		Middleware:    []gin.HandlerFunc{
			auth.SingleAuthenticationMiddleware,
		},
		Body:          func(ctx *gin.Context) {
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