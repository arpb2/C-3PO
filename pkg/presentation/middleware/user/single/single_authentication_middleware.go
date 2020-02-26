package single

import (
	"github.com/arpb2/C-3PO/pkg/domain/architecture/http"
	"github.com/arpb2/C-3PO/pkg/domain/session/repository"
	"github.com/arpb2/C-3PO/pkg/presentation/middleware/user"
)

func CreateMiddleware(tokenHandler repository.TokenRepository) http.Handler {
	return func(ctx *http.Context) {
		user.HandleAuthentication(ctx, tokenHandler)
	}
}
