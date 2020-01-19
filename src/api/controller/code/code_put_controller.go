package code

import (
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

var PutController = controller.Controller{
	Method: "PUT",
	Path:   "/users/:user_id/codes/:code_id",
	Middleware: []gin.HandlerFunc{
		AuthenticationMiddleware,
	},
	Body:   codePut,
}

func codePut(ctx *gin.Context) {
	userId, halt := FetchUserId(ctx)
	if halt {
		return
	}

	codeId, halt := FetchCodeId(ctx)
	if halt {
		return
	}

	code, halt := FetchCode(ctx)
	if halt {
		return
	}

	err := Service.Replace(userId, codeId, code)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": *code,
		"user_id": userId,
		"code_id": codeId,
	})
}
