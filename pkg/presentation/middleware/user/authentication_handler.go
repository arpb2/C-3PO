package user

import (
	"fmt"
	"strconv"

	"github.com/arpb2/C-3PO/pkg/domain/session/repository"

	middleware2 "github.com/arpb2/C-3PO/pkg/domain/session/middleware"
	"github.com/arpb2/C-3PO/pkg/presentation/middleware"

	"github.com/arpb2/C-3PO/pkg/domain/architecture/http"
)

func HandleAuthentication(ctx *http.Context, tokenHandler repository.TokenRepository, strategies ...middleware2.AuthenticationStrategy) {
	authToken := ctx.GetHeader(middleware.HeaderAuthorization)

	if authToken == "" {
		ctx.AbortTransactionWithError(http.CreateUnauthorizedError())
		return
	}

	token, err := tokenHandler.Retrieve(authToken)

	if err != nil {
		ctx.AbortTransactionWithError(err)
		return
	}

	requestedUserId := ctx.GetParameter("user_id")

	if strconv.FormatUint(uint64(token.UserId), 10) == requestedUserId {
		ctx.NextHandler()
		return
	}

	for _, strategy := range strategies {
		authenticated, err := strategy.Authenticate(token, requestedUserId)

		// If any of our strategies has an error we will instantly fail the authentication process.
		// TODO: In a future consider silently dismissing this, logging it somewhere but giving the user a 401.
		if err != nil {
			ctx.AbortTransactionWithError(err)
			return
		}

		// If at least one of the strategies considers us authenticated, then we can continue.
		if authenticated {
			ctx.NextHandler()
			return
		}
	}

	if len(requestedUserId) == 0 {
		fmt.Printf("Got an unauthorized because of no existing parameter 'user_id' in request "+
			"'%s'. Maybe you are malforming the Controller?", ctx.GetUrl())
	}

	ctx.AbortTransactionWithError(http.CreateUnauthorizedError())
}
