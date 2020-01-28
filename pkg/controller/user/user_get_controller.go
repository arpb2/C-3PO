package user

import (
	"github.com/arpb2/C-3PO/api/controller"
	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/api/pipeline"
	userservice "github.com/arpb2/C-3PO/api/service/user"
	usercommand "github.com/arpb2/C-3PO/pkg/command/user"
	gopipeline "github.com/saantiaguilera/go-pipeline"
)

func CreateGetController(executor pipeline.HttpPipeline,
	authMiddleware http.Handler,
	userService userservice.Service) controller.Controller {
	return controller.Controller{
		Method: "GET",
		Path:   "/users/:user_id",
		Middleware: []http.Handler{
			authMiddleware,
		},
		Body: CreateGetBody(executor, userService),
	}
}

func CreateGetBody(exec pipeline.HttpPipeline, userService userservice.Service) http.Handler {
	return func(ctx *http.Context) {
		fetchUserIdCommand := usercommand.CreateFetchUserIdCommand(ctx)
		serviceCommand := usercommand.CreateGetUserCommand(ctx, userService, fetchUserIdCommand.OutputStream)
		renderCommand := usercommand.CreateRenderUserCommand(ctx, serviceCommand.OutputStream)

		graph := gopipeline.CreateSequentialStage(
			fetchUserIdCommand,
			serviceCommand,
			renderCommand,
		)

		exec.Run(ctx, graph)
	}
}
