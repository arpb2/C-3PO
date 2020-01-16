package code

import (
	"fmt"
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/middleware/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

var GetController = controller.Controller{
	Method: "GET",
	Path:   "/users/:user_id/codes/:code_id",
	Middleware: []gin.HandlerFunc{
		auth.UserOrTeacherAuthenticationMiddleware,
	},
	Body:   codeGet,
}

func codeGet(ctx *gin.Context) {
	userId, halt := FetchUserId(ctx)
	if halt {
		return
	}

	codeId, halt := FetchCodeId(ctx)
	if halt {
		return
	}

	code, err := Service.GetCode(userId, codeId)

	if err != nil {
		controller.Halt(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if code == nil {
		controller.Halt(ctx, http.StatusBadRequest, fmt.Sprintf("no code associated to the user_id '%s' and code_id '%s'", userId, codeId))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": *code,
		"user_id": userId,
		"code_id": codeId,
	})
}