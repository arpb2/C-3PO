package code

import (
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/controller/code/code_command"
	"github.com/arpb2/C-3PO/src/api/controller/user/user_command"
	"github.com/arpb2/C-3PO/src/api/executor"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/service"
	"net/http"
)

func CreatePostController(exec executor.Executor, authMiddleware http_wrapper.Handler, codeService service.CodeService) controller.Controller {
	return controller.Controller{
		Method: "POST",
		Path:   "/users/:user_id/codes",
		Middleware: []http_wrapper.Handler{
			authMiddleware,
		},
		Body:   CreatePostBody(exec, codeService),
	}
}

func CreatePostBody(exec executor.Executor, codeService service.CodeService) http_wrapper.Handler {
	return func(ctx *http_wrapper.Context) {
		fetchUserIdCommand := user_command.CreateFetchUserIdCommand(ctx)
		fetchCodeCommand := code_command.CreateFetchCodeCommand(ctx)
		serviceCommand := code_command.CreateCreateCodeCommand(ctx, codeService, fetchUserIdCommand.OutputStream, fetchCodeCommand.OutputStream)

		commands := []executor.Command{
			fetchUserIdCommand,
			fetchCodeCommand,
			serviceCommand,
		}

		if err := controller.BatchRun(exec, commands, ctx); err == nil {
			code := <-serviceCommand.OutputStream
			ctx.WriteJson(http.StatusOK, *code)
		}
	}
}
