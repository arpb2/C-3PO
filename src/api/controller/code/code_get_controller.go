package code

import (
	"fmt"
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

var GetController = controller.Controller{
	Method: "GET",
	Path:   "/users/:user_id/codes/:code_id",
	Middleware: []gin.HandlerFunc{
		AuthenticationMiddleware,
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

	code, err := Service.Read(userId, codeId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if code == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("no code associated to the user_id '%s' and code_id '%s'", userId, codeId),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": *code,
		"user_id": userId,
		"code_id": codeId,
	})
}