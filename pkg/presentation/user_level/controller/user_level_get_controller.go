package controller

import (
	"fmt"

	controller3 "github.com/arpb2/C-3PO/pkg/domain/level/controller"
	controller2 "github.com/arpb2/C-3PO/pkg/domain/user/controller"
	"github.com/arpb2/C-3PO/pkg/domain/user_level/service"

	"github.com/arpb2/C-3PO/pkg/presentation/user/command"
	userlevelcommand "github.com/arpb2/C-3PO/pkg/presentation/user_level/command"

	"github.com/arpb2/C-3PO/pkg/domain/infrastructure/controller"
	"github.com/arpb2/C-3PO/pkg/domain/infrastructure/http"
	"github.com/arpb2/C-3PO/pkg/domain/infrastructure/pipeline"
	gopipeline "github.com/saantiaguilera/go-pipeline"
)

func CreateGetController(exec pipeline.HttpPipeline, authMiddleware http.Handler, userLevelService service.Service) controller.Controller {
	return controller.Controller{
		Method: "GET",
		Path:   fmt.Sprintf("/users/:%s/levels/:%s", controller2.ParamUserId, controller3.ParamLevelId),
		Middleware: []http.Handler{
			authMiddleware,
		},
		Body: CreateGetBody(exec, userLevelService),
	}
}

func CreateGetBody(exec pipeline.HttpPipeline, userLevelService service.Service) http.Handler {
	fetchUserIdCommand := command.CreateFetchUserIdCommand()
	fetchLevelIdCommand := userlevelcommand.CreateFetchLevelIdCommand()
	serviceCommand := userlevelcommand.CreateGetUserLevelCommand(userLevelService)
	renderCommand := userlevelcommand.CreateRenderUserLevelCommand()

	graph := gopipeline.CreateSequentialGroup(
		gopipeline.CreateConcurrentStage(
			fetchUserIdCommand,
			fetchLevelIdCommand,
		),
		gopipeline.CreateSequentialStage(
			serviceCommand,
			renderCommand,
		),
	)

	return func(ctx *http.Context) {
		exec.Run(ctx, graph)
	}
}
