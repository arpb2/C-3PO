package gin

import (
	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/gin-gonic/gin"
)

type ginWriter struct {
	*gin.Context
}

func (g ginWriter) WriteJson(code int, obj interface{}) {
	if !g.IsAborted() {
		g.JSON(code, obj)
	}
}

func (g ginWriter) WriteString(code int, format string, values ...interface{}) {
	if !g.IsAborted() {
		g.String(code, format, values...)
	}
}

func (g ginWriter) WriteStatus(code int) {
	if !g.IsAborted() {
		g.Status(code)
	}
}

func CreateWriter(ctx *gin.Context) http.Writer {
	return &ginWriter{ctx}
}
