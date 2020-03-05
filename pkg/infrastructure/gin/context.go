package gin

import (
	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/gin-gonic/gin"
)

func CreateContext(context *gin.Context) *http.Context {
	return &http.Context{
		Reader:     CreateReader(context),
		Writer:     CreateWriter(context),
		Middleware: CreateMiddleware(context),
	}
}
