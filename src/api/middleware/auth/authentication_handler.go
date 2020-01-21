package auth

import (
	"fmt"
	"github.com/arpb2/C-3PO/src/api/auth"
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"net/http"
	"strconv"
)

type AuthenticationStrategy interface {
	Authenticate(token *auth.Token, userId string) (authorized bool, err error)
}

func HandleAuthentication(ctx *http_wrapper.Context, tokenHandler auth.TokenHandler, strategies ...AuthenticationStrategy) {
	authToken := ctx.GetHeader("Authorization")

	if authToken == "" {
		controller.Halt(ctx, http.StatusUnauthorized, "no 'Authorization' header provided")
		return
	}

	token, err := tokenHandler.Retrieve(authToken)

	if err != nil {
		controller.Halt(ctx, err.Status, err.Error.Error())
		return
	}

	requestedUserId := ctx.Param("user_id")

	if strconv.FormatUint(uint64(token.UserId), 10) == requestedUserId {
		ctx.Next()
		return
	}

	for _, strategy := range strategies {
		authenticated, err := strategy.Authenticate(token, requestedUserId)

		// If any of our strategies has an error we will instantly fail the authentication process.
		// TODO: In a future consider silently dismissing this, logging it somewhere but giving the user a 401.
		if err != nil {
			controller.Halt(ctx, http.StatusInternalServerError, "internal error")
			return
		}

		// If at least one of the strategies considers us authenticated, then we can continue.
		if authenticated {
			ctx.Next()
			return
		}
	}

	go func(userId string, requestUrl string) {
		if userId == "" {
			fmt.Printf("Got an unauthorized because of no existing parameter 'user_id' in request " +
				"'%s'. Maybe you are malforming the Controller?", requestUrl)
		}
	}(requestedUserId, ctx.Url())

	controller.Halt(ctx, http.StatusUnauthorized, "unauthorized")
}
