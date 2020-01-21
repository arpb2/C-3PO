package code

import (
	"fmt"
	"github.com/arpb2/C-3PO/src/api/auth/jwt"
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/middleware/auth/teacher_auth"
	"github.com/arpb2/C-3PO/src/api/service"
	"github.com/arpb2/C-3PO/src/api/service/code_service"
	"github.com/arpb2/C-3PO/src/api/service/teacher_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateGetController() controller.Controller {
	return controller.Controller{
		Method:     	"GET",
		Path:       	"/users/:user_id/codes/:code_id",
		Middleware: 	[]gin.HandlerFunc{
			teacher_auth.CreateMiddleware(
				jwt.CreateTokenHandler(),
				teacher_service.GetService(),
			),
		},
		Body:			CreateGetBody(code_service.GetService()),
	}
}

func CreateGetBody(codeService service.CodeService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId, halt := FetchUserId(ctx)
		if halt {
			return
		}

		codeId, halt := FetchCodeId(ctx)
		if halt {
			return
		}

		code, err := codeService.GetCode(userId, codeId)

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
}
