package controller

import (
	"fmt"

	controller2 "github.com/arpb2/C-3PO/pkg/domain/user/controller"
	"github.com/arpb2/C-3PO/pkg/domain/user/service"

	"github.com/arpb2/C-3PO/pkg/presentation/user/command"

	"github.com/arpb2/C-3PO/pkg/domain/infrastructure/controller"
	"github.com/arpb2/C-3PO/pkg/domain/infrastructure/http"
	"github.com/arpb2/C-3PO/pkg/domain/infrastructure/pipeline"
	gopipeline "github.com/saantiaguilera/go-pipeline"
)

func CreateDeleteController(exec pipeline.HttpPipeline, authMiddleware http.Handler, userService service.Service) controller.Controller {
	return controller.Controller{
		Method: "DELETE",
		Path:   fmt.Sprintf("/users/:%s", controller2.ParamUserId),
		Middleware: []http.Handler{
			authMiddleware,
		},
		Body: CreateDeleteBody(exec, userService),
	}
}

func CreateDeleteBody(exec pipeline.HttpPipeline, userService service.Service) http.Handler {
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
