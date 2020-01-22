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

func CreateGetController(exec executor.Executor, authMiddleware http_wrapper.Handler, codeService service.CodeService) controller.Controller {
	return controller.Controller{
		Method:     	"GET",
		Path:       	"/users/:user_id/codes/:code_id",
		Middleware: 	[]http_wrapper.Handler{
			authMiddleware,
		},
		Body:			CreateGetBody(exec, codeService),
	}
}

func CreateGetBody(exec executor.Executor, codeService service.CodeService) http_wrapper.Handler {
	return func(ctx *http_wrapper.Context) {
		fetchUserIdCommand := user_command.CreateFetchUserIdCommand(ctx)
		fetchCodeIdCommand := code_command.CreateFetchCodeIdCommand(ctx)
		serviceCommand := code_command.CreateGetCodeCommand(ctx, codeService, fetchUserIdCommand.OutputStream, fetchCodeIdCommand.OutputStream)

		commands := []executor.Command{
			fetchUserIdCommand,
			fetchCodeIdCommand,
			serviceCommand,
		}

		if err := controller.BatchRun(exec, commands, ctx); err == nil {
			code := <-serviceCommand.OutputStream
			ctx.WriteJson(http.StatusOK, *code)
		}
	}
}
