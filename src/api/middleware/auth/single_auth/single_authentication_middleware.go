package single_auth

import (
	"github.com/arpb2/C-3PO/src/api/auth"
	middleware_auth "github.com/arpb2/C-3PO/src/api/middleware/auth"
	"github.com/gin-gonic/gin"
)

func CreateMiddleware(tokenHandler auth.TokenHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		middleware_auth.HandleAuthentication(ctx, tokenHandler)
	}
}
