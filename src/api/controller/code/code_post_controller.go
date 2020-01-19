package code

import (
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/middleware/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

var PostController = controller.Controller{
	Method: "POST",
	Path:   "/users/:user_id/codes",
	Middleware: []gin.HandlerFunc{
		auth.UserOrTeacherAuthenticationMiddleware,
	},
	Body:   codePost,
}

func codePost(ctx *gin.Context) {
	userId, halt := FetchUserId(ctx)
	if halt {
		return
	}

	code, halt := FetchCode(ctx)
	if halt {
		return
	}

	codeId, err := Service.CreateCode(userId, code)

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
