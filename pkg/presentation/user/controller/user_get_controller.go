package controller

import (
	"fmt"

	"github.com/arpb2/C-3PO/pkg/presentation/user/command"

	"github.com/arpb2/C-3PO/pkg/domain/controller"
	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/domain/pipeline"
	userservice "github.com/arpb2/C-3PO/pkg/domain/service/user"
	gopipeline "github.com/saantiaguilera/go-pipeline"
)

func CreateGetController(executor pipeline.HttpPipeline,
	authMiddleware http.Handler,
	userService userservice.Service) controller.Controller {
	return controller.Controller{
		Method: "GET",
		Path:   fmt.Sprintf("/users/:%s", controller.ParamUserId),
		Middleware: []http.Handler{
			authMiddleware,
		},
		Body: CreateGetBody(executor, userService),
	}
}

func CreateGetBody(exec pipeline.HttpPipeline, userService userservice.Service) http.Handler {
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
