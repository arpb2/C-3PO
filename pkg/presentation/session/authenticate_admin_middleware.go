package session

import (
	"github.com/arpb2/C-3PO/pkg/data/usecase/session"
	"github.com/arpb2/C-3PO/pkg/domain/http"
)

func CreateAuthenticateAdminMiddleware(secret []byte) http.Handler {
	useCase := session.CreateAdminAuthenticationUseCase(secret)

	return func(ctx *http.Context) {
		if ctx.GetValue(authenticated) == true {
			ctx.NextHandler()
			return
		}

		authToken := ctx.GetHeader("Authorization")

		err := useCase(authToken)

		if err != nil {
			ctx.AbortTransactionWithError(err)
		} else {
			ctx.SetValue(authenticated, true)
			ctx.NextHandler()
		}
	}
}
