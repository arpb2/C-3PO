package code

import (
	"github.com/arpb2/C-3PO/src/api/command/code_command"
	"github.com/arpb2/C-3PO/src/api/command/user_command"
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/executor"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/service"
)

func CreatePutController(exec executor.HttpExecutor, authMiddleware http_wrapper.Handler, codeService service.CodeService) controller.Controller {
	return controller.Controller{
		Method: "PUT",
		Path:   "/users/:user_id/codes/:code_id",
		Middleware: []http_wrapper.Handler{
			authMiddleware,
		},
		Body:   CreatePutBody(exec, codeService),
	}
}

func CreatePutBody(exec executor.HttpExecutor, codeService service.CodeService) http_wrapper.Handler {
	return func(ctx *http_wrapper.Context) {
		fetchUserIdCommand := user_command.CreateFetchUserIdCommand(ctx)
		fetchCodeCommand := code_command.CreateFetchCodeCommand(ctx)
		fetchCodeIdCommand := code_command.CreateFetchCodeIdCommand(ctx)
		serviceCommand := code_command.CreateReplaceCodeCommand(ctx, codeService,
			fetchCodeIdCommand.OutputStream, fetchUserIdCommand.OutputStream, fetchCodeCommand.OutputStream)
		renderCommand := code_command.CreateRenderCodeCommand(ctx, serviceCommand.OutputStream)

		commands := []executor.Command{
			fetchUserIdCommand,
			fetchCodeIdCommand,
			fetchCodeCommand,
			serviceCommand,
			renderCommand,
		}

		exec.BatchRun(ctx, commands)
	}
}
