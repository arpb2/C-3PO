package single

import (
	"github.com/arpb2/C-3PO/pkg/domain/auth"
	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/presentation/middleware/user"
)

func CreateMiddleware(tokenHandler auth.TokenHandler) http.Handler {
	return func(ctx *http.Context) {
		user.HandleAuthentication(ctx, tokenHandler)
	}
}
