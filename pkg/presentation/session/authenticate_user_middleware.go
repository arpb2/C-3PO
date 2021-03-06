package session

import (
	session2 "github.com/arpb2/C-3PO/pkg/data/repository/session"
	"github.com/arpb2/C-3PO/pkg/data/usecase/session"
	"github.com/arpb2/C-3PO/pkg/domain/http"
)

func CreateAuthenticateUserMiddleware(userParam string, tokenHandler session2.TokenRepository) http.Handler {
	useCase := session.CreateUserAuthenticationUseCase(tokenHandler)

	return func(ctx *http.Context) {
		if ctx.GetValue(authenticated) == true {
			ctx.NextHandler()
			return
		}

		authToken := ctx.GetHeader("Authorization")
		userId := ctx.GetParameter(userParam)

		err := useCase(authToken, userId)

		if err != nil {
			ctx.AbortTransactionWithError(err)
		} else {
			ctx.SetValue(authenticated, true)
			ctx.NextHandler()
		}
	}
}
