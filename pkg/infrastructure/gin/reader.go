package gin

import (
	"github.com/arpb2/C-3PO/pkg/domain/http"
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

func (g ginReader) GetMethod() string {
	return g.Request.Method
}

func (g ginReader) GetUrl() string {
	return g.Request.URL.String()
}

func CreateReader(ctx *gin.Context) http.Reader {
	return &ginReader{ctx}
}
