package gin_engine_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/arpb2/C-3PO/api/controller"
	"github.com/arpb2/C-3PO/api/http_wrapper"
	gin_engine "github.com/arpb2/C-3PO/pkg/engine/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetPortReturnsSpecificIfUsingEnvVar(t *testing.T) {
	_ = os.Setenv("PORT", "1234")
	assert.Equal(t, "1234", gin_engine.GetPort())
	_ = os.Unsetenv("PORT")
}

func TestGetPortReturns8080(t *testing.T) {
	assert.Equal(t, "8080", gin_engine.GetPort())
}

func TestRegisterMiddleware(t *testing.T) {
	e := gin_engine.New()
	e.Register(controller.Controller{
		Method: "GET",
		Path:   "/test",
		Middleware: []http_wrapper.Handler{func(c *http_wrapper.Context) {
			c.WriteString(http.StatusOK, "Test response")
			c.AbortTransaction()
		}},
		Body: func(c *http_wrapper.Context) {
			panic("shouldn't reach here.")
		},
	})

	req, _ := http.NewRequest("GET", "/test", strings.NewReader(""))
	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Test response", w.Body.String())
}
