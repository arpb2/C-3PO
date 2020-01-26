package auth

import (
	"fmt"
	"github.com/arpb2/C-3PO/api/auth"
	"github.com/arpb2/C-3PO/api/http_wrapper"
	auth_middleware "github.com/arpb2/C-3PO/api/middleware/auth"
	"strconv"
)

func HandleAuthentication(ctx *http_wrapper.Context, tokenHandler auth.TokenHandler, strategies ...auth_middleware.AuthenticationStrategy) {
	authToken := ctx.GetHeader("Authorization")

	if authToken == "" {
		ctx.AbortTransactionWithError(http_wrapper.CreateUnauthorizedError())
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

	go func(userId string, requestUrl string) {
		if userId == "" {
			fmt.Printf("Got an unauthorized because of no existing parameter 'user_id' in request "+
				"'%s'. Maybe you are malforming the Controller?", requestUrl)
		}
	}(requestedUserId, ctx.GetUrl())

	ctx.AbortTransactionWithError(http_wrapper.CreateUnauthorizedError())
}
