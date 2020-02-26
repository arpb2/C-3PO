package controller

import (
	"fmt"
	controller2 "github.com/arpb2/C-3PO/pkg/presentation/level/controller"
	controller4 "github.com/arpb2/C-3PO/pkg/presentation/user/controller"

	"github.com/arpb2/C-3PO/pkg/domain/user_level/service"

	"github.com/arpb2/C-3PO/pkg/presentation/user/command"
	userlevelcommand "github.com/arpb2/C-3PO/pkg/presentation/user_level/command"

	"github.com/arpb2/C-3PO/pkg/domain/architecture/controller"
	"github.com/arpb2/C-3PO/pkg/domain/architecture/http"
	"github.com/arpb2/C-3PO/pkg/domain/architecture/pipeline"
	gopipeline "github.com/saantiaguilera/go-pipeline"
)

func CreateGetController(exec pipeline.HttpPipeline, authMiddleware http.Handler, userLevelService service.Service) controller.Controller {
	return controller.Controller{
		Method: "GET",
		Path:   fmt.Sprintf("/users/:%s/levels/:%s", controller4.ParamUserId, controller2.ParamLevelId),
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
