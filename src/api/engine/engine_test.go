package engine_test

import (
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/engine"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestGetPortReturnsSpecificIfUsingEnvVar(t *testing.T) {
	_ = os.Setenv("PORT", "1234")
	assert.Equal(t, "1234", engine.GetPort())
	_ = os.Unsetenv("PORT")
}

func TestGetPortReturns8080(t *testing.T) {
	assert.Equal(t, "8080", engine.GetPort())
}

func TestRegisterMiddleware(t *testing.T) {
	e := engine.CreateBasicServerEngine()
	e.Register(controller.Controller{
		Method:     "GET",
		Path:       "/test",
		Middleware: []gin.HandlerFunc{func(c *gin.Context) {
			c.String(http.StatusOK, "Test response")
			c.Abort()
		}},
		Body:       func(c *gin.Context) {
			panic("shouldn't reach here.")
		},
	})

	req, _ := http.NewRequest("GET", "/test", strings.NewReader(""))
	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Test response", w.Body.String())
}