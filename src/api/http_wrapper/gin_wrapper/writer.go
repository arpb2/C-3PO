package gin_wrapper

import (
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/gin-gonic/gin"
)

type ginWriter struct {
	*gin.Context
}

func (g ginWriter) WriteJson(code int, obj interface{}) {
	g.JSON(code, obj)
}

func (g ginWriter) WriteString(code int, format string, values ...interface{}) {
	g.String(code, format, values...)
}

func (g ginWriter) WriteStatus(code int) {
	g.Status(code)
}

func CreateWriter(ctx *gin.Context) http_wrapper.Writer {
	return &ginWriter{ctx}
}
