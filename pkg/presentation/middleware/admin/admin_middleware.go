package admin

import (
	"bytes"

	"github.com/arpb2/C-3PO/pkg/presentation/middleware"

	"github.com/arpb2/C-3PO/pkg/domain/http"
)

func CreateMiddleware(secret []byte) http.Handler {
	return func(ctx *http.Context) {
		tkn := ctx.GetHeader(middleware.HeaderAuthorization)

		if len(tkn) == 0 || bytes.Compare([]byte(tkn), secret) != 0 {
			ctx.AbortTransactionWithError(http.CreateUnauthorizedError())
		} else {
			ctx.NextHandler()
		}
	}
}
