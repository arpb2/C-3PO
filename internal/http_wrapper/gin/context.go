package gin_wrapper

import (
	"github.com/arpb2/C-3PO/api/http_wrapper"
	"github.com/gin-gonic/gin"
)

func CreateContext(context *gin.Context) *http_wrapper.Context {
	return &http_wrapper.Context{
		Reader:     CreateReader(context),
		Writer:     CreateWriter(context),
		Middleware: CreateMiddleware(context),
	}
}
