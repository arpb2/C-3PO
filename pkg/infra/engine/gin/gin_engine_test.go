package gin_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/arpb2/C-3PO/pkg/domain/controller"
	httpwrapper "github.com/arpb2/C-3PO/pkg/domain/http"
	ginengine "github.com/arpb2/C-3PO/pkg/infra/engine/gin"
	"github.com/stretchr/testify/assert"
)

func TestRegisterMiddleware(t *testing.T) {
	e := ginengine.CreateEngine("8080")
	e.Register(controller.Controller{
		Method: "GET",
		Path:   "/test",
		Middleware: []httpwrapper.Handler{func(c *httpwrapper.Context) {
			c.WriteString(http.StatusOK, "Test response")
			c.AbortTransaction()
		}},
		Body: func(c *httpwrapper.Context) {
			panic("shouldn't reach here.")
		},
	})

	req, _ := http.NewRequest("GET", "/test", strings.NewReader(""))
	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Test response", w.Body.String())
}
