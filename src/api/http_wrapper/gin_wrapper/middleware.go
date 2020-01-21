package gin_wrapper

import (
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/gin-gonic/gin"
)

type ginMiddleware struct {
	*gin.Context
}

func (g ginMiddleware) NextHandler() {
	g.Next()
}

func (g ginMiddleware) AbortTransaction() {
	g.Abort()
}

func (g ginMiddleware) AbortTransactionWithStatus(code int, jsonObj interface{}) {
	g.AbortWithStatusJSON(code, jsonObj)
}

func CreateMiddleware(ctx *gin.Context) http_wrapper.Middleware {
	return &ginMiddleware{ctx}
}
