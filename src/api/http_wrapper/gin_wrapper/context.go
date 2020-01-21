package gin_wrapper

import (
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
)

func CreateContext(context *gin.Context) *http_wrapper.Context {
	return &http_wrapper.Context{
		Reader:     CreateReader(context),
		Writer:     context,
		Middleware: context,
	}
}

func CreateTestContext() (*http_wrapper.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)

	return CreateContext(c), recorder
}