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

func CreateDeleteController(exec pipeline.HttpPipeline, authMiddleware http.Handler, userService userservice.Service) controller.Controller {
	return controller.Controller{
		Method: "DELETE",
		Path:   fmt.Sprintf("/users/:%s", controller.ParamUserId),
		Middleware: []http.Handler{
			authMiddleware,
		},
		Body: CreateDeleteBody(exec, userService),
	}
}

func CreateDeleteBody(exec pipeline.HttpPipeline, userService userservice.Service) http.Handler {
	fetchUserIdCommand := command.CreateFetchUserIdCommand()
	serviceCommand := command.CreateDeleteUserCommand(userService)
	renderCommand := command.CreateRenderEmptyCommand()

	graph := gopipeline.CreateSequentialStage(
		fetchUserIdCommand,
		serviceCommand,
		renderCommand,
	)

	return func(ctx *http.Context) {
		exec.Run(ctx, graph)
	}
}
