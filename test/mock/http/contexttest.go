package http

import (
	"net/http/httptest"

	"github.com/arpb2/C-3PO/pkg/domain/http"
	internal "github.com/arpb2/C-3PO/pkg/infra/http/gin"
	"github.com/gin-gonic/gin"
)

func CreateTestContext() (*http.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)

	return internal.CreateContext(c), recorder
}
