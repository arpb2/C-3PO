package session

import (
	"github.com/arpb2/C-3PO/pkg/domain/http"
)

func CreateAuthenticateDebugMiddleware() http.Handler {
	return func(ctx *http.Context) {
		if ctx.GetValue(authenticated) == true {
			ctx.NextHandler()
			return
		}

		authToken := ctx.GetHeader("Authorization")

		if authToken == "DEBUG" {
			ctx.SetValue(authenticated, true)
			ctx.NextHandler()
		}
	}
}
