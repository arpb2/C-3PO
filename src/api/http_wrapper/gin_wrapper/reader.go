package gin_wrapper

import (
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/gin-gonic/gin"
)

type ginReader struct {
	*gin.Context
}

func (g ginReader) Url() string {
	return g.Request.URL.String()
}

func CreateReader(ctx *gin.Context) http_wrapper.Reader {
	return &ginReader{ctx}
}
