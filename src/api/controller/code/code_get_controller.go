package code

import (
	"fmt"
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/service"
	"net/http"
)

func CreateGetController(authMiddleware http_wrapper.Handler, codeService service.CodeService) controller.Controller {
	return controller.Controller{
		Method:     	"GET",
		Path:       	"/users/:user_id/codes/:code_id",
		Middleware: 	[]http_wrapper.Handler{
			authMiddleware,
		},
		Body:			CreateGetBody(codeService),
	}
}

func CreateGetBody(codeService service.CodeService) http_wrapper.Handler {
	return func(ctx *http_wrapper.Context) {
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
			controller.Halt(ctx, http.StatusInternalServerError, "internal error")
			return
		}

		if code == nil {
			controller.Halt(ctx, http.StatusNotFound, fmt.Sprintf("no code associated to the user_id '%d' and code_id '%d'", userId, codeId))
			return
		}

		ctx.JSON(http.StatusOK, http_wrapper.Json{
			"code": *code,
			"user_id": userId,
			"code_id": codeId,
		})
	}
}
