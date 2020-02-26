package http

import (
	"net/http/httptest"

	internal "github.com/arpb2/C-3PO/cmd/c3po/infrastructure/http/gin"
	"github.com/arpb2/C-3PO/pkg/domain/infrastructure/http"
	"github.com/gin-gonic/gin"
)

func CreateTestContext() (*http.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)

	return internal.CreateContext(c), recorder
}
