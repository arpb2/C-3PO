package controller

import (
	"fmt"

	"github.com/arpb2/C-3PO/pkg/domain/user/service"

	"github.com/arpb2/C-3PO/pkg/presentation/user/command"

	"github.com/arpb2/C-3PO/pkg/domain/architecture/controller"
	"github.com/arpb2/C-3PO/pkg/domain/architecture/http"
	"github.com/arpb2/C-3PO/pkg/domain/architecture/pipeline"
	gopipeline "github.com/saantiaguilera/go-pipeline"
)

func CreateGetController(executor pipeline.HttpPipeline,
	authMiddleware http.Handler,
	userService service.Service) controller.Controller {
	return controller.Controller{
		Method: "GET",
		Path:   fmt.Sprintf("/users/:%s", ParamUserId),
		Middleware: []http.Handler{
			authMiddleware,
		},
		Body: CreateGetBody(executor, userService),
	}
}

func CreateGetBody(exec pipeline.HttpPipeline, userService service.Service) http.Handler {
	fetchUserIdCommand := command.CreateFetchUserIdCommand()
	serviceCommand := command.CreateGetUserCommand(userService)
	renderCommand := command.CreateRenderUserCommand()

	graph := gopipeline.CreateSequentialStage(
		fetchUserIdCommand,
		serviceCommand,
		renderCommand,
	)

	return func(ctx *http.Context) {
		exec.Run(ctx, graph)
	}
}
