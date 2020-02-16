package user_level

import (
	"fmt"

	userlevelservice "github.com/arpb2/C-3PO/api/service/user_level"
	userlevelcommand "github.com/arpb2/C-3PO/pkg/command/user_level"

	controller2 "github.com/arpb2/C-3PO/pkg/controller"

	"github.com/arpb2/C-3PO/api/controller"
	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/api/pipeline"
	usercommand "github.com/arpb2/C-3PO/pkg/command/user"
	gopipeline "github.com/saantiaguilera/go-pipeline"
)

func CreateGetController(exec pipeline.HttpPipeline, authMiddleware http.Handler, userLevelService userlevelservice.Service) controller.Controller {
	return controller.Controller{
		Method: "GET",
		Path:   fmt.Sprintf("/users/:%s/levels/:%s", controller2.ParamUserId, controller2.ParamLevelId),
		Middleware: []http.Handler{
			authMiddleware,
		},
		Body: CreateGetBody(exec, userLevelService),
	}
}

func CreateGetBody(exec pipeline.HttpPipeline, userLevelService userlevelservice.Service) http.Handler {
	fetchUserIdCommand := usercommand.CreateFetchUserIdCommand()
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
