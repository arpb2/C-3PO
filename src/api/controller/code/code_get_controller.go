package code

import (
	"github.com/arpb2/C-3PO/src/api/command/code_command"
	"github.com/arpb2/C-3PO/src/api/command/user_command"
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/executor"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/service"
	"github.com/saantiaguilera/go-pipeline/pkg/stage/concurrent"
	"github.com/saantiaguilera/go-pipeline/pkg/stage/sequential"
)

func CreateGetController(exec executor.HttpPipeline, authMiddleware http_wrapper.Handler, codeService service.CodeService) controller.Controller {
	return controller.Controller{
		Method:     	"GET",
		Path:       	"/users/:user_id/codes/:code_id",
		Middleware: 	[]http_wrapper.Handler{
			authMiddleware,
		},
		Body:			CreateGetBody(exec, codeService),
	}
}

func CreateGetBody(exec executor.HttpPipeline, codeService service.CodeService) http_wrapper.Handler {
	return func(ctx *http_wrapper.Context) {
		fetchUserIdCommand := user_command.CreateFetchUserIdCommand(ctx)
		fetchCodeIdCommand := code_command.CreateFetchCodeIdCommand(ctx)
		serviceCommand := code_command.CreateGetCodeCommand(ctx, codeService, fetchUserIdCommand.OutputStream, fetchCodeIdCommand.OutputStream)
		renderCommand := code_command.CreateRenderCodeCommand(ctx, serviceCommand.OutputStream)

		graph := sequential.CreateSequentialGroup(
			concurrent.CreateConcurrentStage(
				fetchUserIdCommand,
				fetchCodeIdCommand,
			),
			sequential.CreateSequentialStage(
				serviceCommand,
				renderCommand,
			),
		)

		exec.Run(ctx, graph)
	}
}
