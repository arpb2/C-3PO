package auth

import (
	"fmt"
	"github.com/arpb2/C-3PO/src/api/auth"
	"github.com/arpb2/C-3PO/src/api/auth/jwt"
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var TokenHandler auth.TokenHandler = jwt.TokenHandler{
	Secret: jwt.FetchJwtSecret(),
}

type AuthenticationStrategy func(token *auth.Token, userId string) (authorized bool, err error)

func handleAuthentication(ctx *gin.Context, strategies ...AuthenticationStrategy) {
	authToken := ctx.GetHeader("Authorization")

	if authToken == "" {
		controller.Halt(ctx, http.StatusUnauthorized, "no 'Authorization' header provided")
		return
	}

	token, err := TokenHandler.Retrieve(authToken)

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
		authenticated, err := strategy(token, requestedUserId)

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
	}(requestedUserId, ctx.Request.URL.String())

	controller.Halt(ctx, http.StatusUnauthorized, "unauthorized")
}
