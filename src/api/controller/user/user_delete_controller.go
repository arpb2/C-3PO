package user

import (
	"github.com/arpb2/C-3PO/src/api/command/user_command"
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/executor"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/service"
	"github.com/saantiaguilera/go-pipeline/pkg/stage/sequential"
)

func CreateDeleteController(exec executor.HttpPipeline, authMiddleware http_wrapper.Handler, userService service.UserService) controller.Controller {
	return controller.Controller{
		Method: "DELETE",
		Path:   "/users/:user_id",
		Middleware: []http_wrapper.Handler{
			authMiddleware,
		},
		Body:   CreateDeleteBody(exec, userService),
	}
}

func CreateDeleteBody(exec executor.HttpPipeline, userService service.UserService) http_wrapper.Handler {
	return func(ctx *http_wrapper.Context) {
		fetchUserIdCommand := user_command.CreateFetchUserIdCommand(ctx)
		serviceCommand := user_command.CreateDeleteUserCommand(ctx, userService, fetchUserIdCommand.OutputStream)
		renderCommand := user_command.CreateRenderEmptyCommand(ctx, serviceCommand.OutputStream)

		graph := sequential.CreateSequentialStage(
			fetchUserIdCommand,
			serviceCommand,
			renderCommand,
		)

		exec.Run(ctx, graph)
	}
}