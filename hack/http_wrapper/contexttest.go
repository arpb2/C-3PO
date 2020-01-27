package http_wrapper

import (
	"github.com/arpb2/C-3PO/api/http_wrapper"
	internal "github.com/arpb2/C-3PO/pkg/http_wrapper/gin"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
)

func CreateTestContext() (*http_wrapper.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)

	return internal.CreateContext(c), recorder
}
