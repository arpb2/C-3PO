package user_controller

import (
	"github.com/arpb2/C-3PO/api/controller"
	"github.com/arpb2/C-3PO/api/http_wrapper"
	"github.com/arpb2/C-3PO/api/pipeline"
	user_service "github.com/arpb2/C-3PO/api/service/user"
	user_command "github.com/arpb2/C-3PO/internal/command/user"
	"github.com/saantiaguilera/go-pipeline/pkg/stage/sequential"
)

func CreateDeleteController(exec pipeline.HttpPipeline, authMiddleware http_wrapper.Handler, userService user_service.Service) controller.Controller {
	return controller.Controller{
		Method: "DELETE",
		Path:   "/users/:user_id",
		Middleware: []http_wrapper.Handler{
			authMiddleware,
		},
		Body: CreateDeleteBody(exec, userService),
	}
}

func CreateDeleteBody(exec pipeline.HttpPipeline, userService user_service.Service) http_wrapper.Handler {
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
