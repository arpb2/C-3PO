package code

import (
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/service"
	"net/http"
)

func CreatePutController(authMiddleware http_wrapper.Handler, codeService service.CodeService) controller.Controller {
	return controller.Controller{
		Method: "PUT",
		Path:   "/users/:user_id/codes/:code_id",
		Middleware: []http_wrapper.Handler{
			authMiddleware,
		},
		Body:   CreatePutBody(codeService),
	}
}

func CreatePutBody(codeService service.CodeService) http_wrapper.Handler {
	return func(ctx *http_wrapper.Context) {
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

		err := codeService.ReplaceCode(userId, codeId, code)

		if err != nil {
			controller.Halt(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, http_wrapper.Json{
			"code":    *code,
			"user_id": userId,
			"code_id": codeId,
		})
	}
}
