package controller

import (
	"fmt"
	"github.com/arpb2/C-3PO/pkg/presentation/user/command"
	command2 "github.com/arpb2/C-3PO/pkg/presentation/user_level/command"

	"github.com/arpb2/C-3PO/pkg/domain/controller"
	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/domain/pipeline"
	userlevelservice "github.com/arpb2/C-3PO/pkg/domain/service/user_level"
	gopipeline "github.com/saantiaguilera/go-pipeline"
)

func CreatePutController(exec pipeline.HttpPipeline, authMiddleware http.Handler, userLevelService userlevelservice.Service) controller.Controller {
	return controller.Controller{
		Method: "PUT",
		Path:   fmt.Sprintf("/users/:%s/levels/:%s", controller.ParamUserId, controller.ParamLevelId),
		Middleware: []http.Handler{
			authMiddleware,
		},
		Body: CreatePutBody(exec, userLevelService),
	}
}

func CreatePutBody(exec pipeline.HttpPipeline, userLevelService userlevelservice.Service) http.Handler {
	fetchUserIdCommand := command.CreateFetchUserIdCommand()
	fetchCodeCommand := command2.CreateFetchCodeCommand()
	fetchLevelIdCommand := command2.CreateFetchLevelIdCommand()
	serviceCommand := command2.CreateWriteUserLevelCommand(userLevelService)
	renderCommand := command2.CreateRenderUserLevelCommand()

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
