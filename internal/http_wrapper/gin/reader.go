package gin_wrapper

import (
	"github.com/arpb2/C-3PO/api/http_wrapper"
	"github.com/gin-gonic/gin"
)

type ginReader struct {
	*gin.Context
}

func (g ginReader) GetParameter(key string) string {
	return g.Param(key)
}

func (g ginReader) GetFormData(key string) (string, bool) {
	return g.GetPostForm(key)
}

func (g ginReader) ReadBody(obj interface{}) error {
	return g.ShouldBindJSON(obj)
}

func (g ginReader) GetUrl() string {
	return g.Request.URL.String()
}

func CreateReader(ctx *gin.Context) http_wrapper.Reader {
	return &ginReader{ctx}
}
