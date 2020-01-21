package single_auth

import (
	"github.com/arpb2/C-3PO/src/api/auth"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	middleware_auth "github.com/arpb2/C-3PO/src/api/middleware/auth"
)

func CreateMiddleware(tokenHandler auth.TokenHandler) http_wrapper.Handler {
	return func(ctx *http_wrapper.Context) {
		middleware_auth.HandleAuthentication(ctx, tokenHandler)
	}
}