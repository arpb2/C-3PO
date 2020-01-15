package auth

import (
	"fmt"
	"github.com/arpb2/C-3PO/src/api/engine"
	"github.com/arpb2/C-3PO/src/api/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var TeacherService service.TeacherService // TODO Set.

func handleAuthentication(ctx *gin.Context, allowTeacher bool) {
	authToken := ctx.GetHeader("Authorization")

	if authToken == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "no 'Authorization' header provided",
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

	userIdStr := strconv.FormatUint(uint64(token.UserId), 10)
	requestedUserId := ctx.Param("user_id")

	if userIdStr == requestedUserId {
		ctx.Next()
		return
	}

	if allowTeacher {
		students, serviceErr := TeacherService.GetStudents(userIdStr)

		if serviceErr != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal error",
			})
			ctx.Abort()
			return
		}

		if students != nil {
			for _, student := range *students {
				if strconv.FormatUint(uint64(student), 10) == requestedUserId {
					ctx.Next()
					return
				}
			}
		}
	}

	go func(userId string, requestUrl string) {
		if userId == "" {
			fmt.Printf("Got an unauthorized because of no existing parameter 'user_id' in request " +
				"'%s'. Maybe you are malforming the Controller?", requestUrl)
		}
	}(requestedUserId, ctx.Request.URL.String())

	ctx.JSON(http.StatusUnauthorized, gin.H{
		"error": "unauthorized",
	})
	ctx.Abort()
}
