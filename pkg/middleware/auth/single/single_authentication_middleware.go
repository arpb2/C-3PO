package single

import (
	"github.com/arpb2/C-3PO/api/auth"
	"github.com/arpb2/C-3PO/api/http"
	middlewareauth "github.com/arpb2/C-3PO/pkg/middleware/auth"
)

func CreateMiddleware(tokenHandler auth.TokenHandler) http.Handler {
	return func(ctx *http.Context) {
		middlewareauth.HandleAuthentication(ctx, tokenHandler)
	}
}
