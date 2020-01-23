package code

import (
	"github.com/arpb2/C-3PO/src/api/command/code_command"
	"github.com/arpb2/C-3PO/src/api/command/user_command"
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/executor"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/service"
)

func CreatePostController(exec executor.HttpExecutor, authMiddleware http_wrapper.Handler, codeService service.CodeService) controller.Controller {
	return controller.Controller{
		Method: "POST",
		Path:   "/users/:user_id/codes",
		Middleware: []http_wrapper.Handler{
			authMiddleware,
		},
		Body:   CreatePostBody(exec, codeService),
	}
}

func CreatePostBody(exec executor.HttpExecutor, codeService service.CodeService) http_wrapper.Handler {
	return func(ctx *http_wrapper.Context) {
		fetchUserIdCommand := user_command.CreateFetchUserIdCommand(ctx)
		fetchCodeCommand := code_command.CreateFetchCodeCommand(ctx)
		serviceCommand := code_command.CreateCreateCodeCommand(ctx, codeService, fetchUserIdCommand.OutputStream, fetchCodeCommand.OutputStream)
		renderCommand := code_command.CreateRenderCodeCommand(ctx, serviceCommand.OutputStream)

		commands := []executor.Command{
			fetchUserIdCommand,
			fetchCodeCommand,
			serviceCommand,
			renderCommand,
		}

		exec.BatchRun(ctx, commands)
	}
}
