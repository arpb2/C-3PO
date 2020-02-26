package single

import (
	"github.com/arpb2/C-3PO/pkg/domain/infrastructure/http"
	"github.com/arpb2/C-3PO/pkg/domain/session/token"
	"github.com/arpb2/C-3PO/pkg/presentation/middleware/user"
)

func CreateMiddleware(tokenHandler token.Handler) http.Handler {
	return func(ctx *http.Context) {
		user.HandleAuthentication(ctx, tokenHandler)
	}
}
