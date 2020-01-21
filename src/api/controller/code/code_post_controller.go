package code

import (
	"github.com/arpb2/C-3PO/src/api/auth/jwt"
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/middleware/auth/teacher_auth"
	"github.com/arpb2/C-3PO/src/api/service"
	"github.com/arpb2/C-3PO/src/api/service/code_service"
	"github.com/arpb2/C-3PO/src/api/service/teacher_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreatePostController() controller.Controller {
	return controller.Controller{
		Method: "POST",
		Path:   "/users/:user_id/codes",
		Middleware: []gin.HandlerFunc{
			teacher_auth.CreateMiddleware(
				jwt.CreateTokenHandler(),
				teacher_service.GetService(),
			),
		},
		Body:   CreatePostBody(code_service.GetService()),
	}
}

func CreatePostBody(codeService service.CodeService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, halt := FetchUserId(ctx)
		if halt {
			return
		}

		code, halt := FetchCode(ctx)
		if halt {
			return
		}

		codeId, err := codeService.CreateCode(userId, code)

		if err != nil {
			controller.Halt(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"code": *code,
			"user_id": userId,
			"code_id": codeId,
		})
	}
}
