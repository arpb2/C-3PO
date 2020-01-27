package user

import (
	"github.com/arpb2/C-3PO/api/controller"
	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/api/pipeline"
	userservice "github.com/arpb2/C-3PO/api/service/user"
	usercommand "github.com/arpb2/C-3PO/pkg/command/user"
	"github.com/saantiaguilera/go-pipeline/pkg/stage/sequential"
)

func CreateDeleteController(exec pipeline.HttpPipeline, authMiddleware http.Handler, userService userservice.Service) controller.Controller {
	return controller.Controller{
		Method: "DELETE",
		Path:   "/users/:user_id",
		Middleware: []http.Handler{
			authMiddleware,
		},
		Body: CreateDeleteBody(exec, userService),
	}
}

func CreateDeleteBody(exec pipeline.HttpPipeline, userService userservice.Service) http.Handler {
	return func(ctx *http.Context) {
		fetchUserIdCommand := usercommand.CreateFetchUserIdCommand(ctx)
		serviceCommand := usercommand.CreateDeleteUserCommand(ctx, userService, fetchUserIdCommand.OutputStream)
		renderCommand := usercommand.CreateRenderEmptyCommand(ctx, serviceCommand.OutputStream)

		graph := sequential.CreateSequentialStage(
			fetchUserIdCommand,
			serviceCommand,
			renderCommand,
		)

		exec.Run(ctx, graph)
	}
}
