package code

import (
	"fmt"
	"github.com/arpb2/C-3PO/src/api/engine"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func HandleAuthentication(ctx *gin.Context) {
	authToken := ctx.GetHeader("Authentication")

	if authToken == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "no 'Authentication' header provided",
		})
		ctx.Abort()
		return
	}

	token, err := engine.DefaultTokenHandler.Retrieve(authToken)

	if err != nil {
		ctx.JSON(err.Status, gin.H{
			"error": err.Error.Error(),
		})
		ctx.Abort()
		return
	}

	if strconv.FormatUint(uint64(token.UserId), 10) != ctx.Param("user_id") {
		go func(userId string, requestUrl string) {
			if userId == "" {
				fmt.Printf("Got an unauthorized because of no existing parameter 'user_id' in request " +
					"'%s'. Maybe you are malforming the Controller?", requestUrl)
			}
		}(ctx.Param("user_id"), ctx.Request.URL.String())

		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		ctx.Abort()
		return
	}

	ctx.Next()
}
