package user_level

import (
	"fmt"

	controller2 "github.com/arpb2/C-3PO/pkg/controller"

	"github.com/arpb2/C-3PO/api/controller"
	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/api/pipeline"
	userlevelservice "github.com/arpb2/C-3PO/api/service/user_level"
	usercommand "github.com/arpb2/C-3PO/pkg/command/user"
	userlevelcommand "github.com/arpb2/C-3PO/pkg/command/user_level"
	gopipeline "github.com/saantiaguilera/go-pipeline"
)

func CreatePutController(exec pipeline.HttpPipeline, authMiddleware http.Handler, userLevelService userlevelservice.Service) controller.Controller {
	return controller.Controller{
		Method: "PUT",
		Path:   fmt.Sprintf("/users/:%s/levels/:%s", controller2.ParamUserId, controller2.ParamLevelId),
		Middleware: []http.Handler{
			authMiddleware,
		},
		Body: CreatePutBody(exec, userLevelService),
	}
}

func CreatePutBody(exec pipeline.HttpPipeline, userLevelService userlevelservice.Service) http.Handler {
	fetchUserIdCommand := usercommand.CreateFetchUserIdCommand()
	fetchCodeCommand := userlevelcommand.CreateFetchCodeCommand()
	fetchLevelIdCommand := userlevelcommand.CreateFetchLevelIdCommand()
	serviceCommand := userlevelcommand.CreateReplaceUserLevelCommand(userLevelService)
	renderCommand := userlevelcommand.CreateRenderUserLevelCommand()

	graph := gopipeline.CreateSequentialGroup(
		gopipeline.CreateConcurrentStage(
			fetchUserIdCommand,
			fetchCodeCommand,
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
