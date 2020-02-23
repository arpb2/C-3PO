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

func CreateGetController(exec pipeline.HttpPipeline, authMiddleware http.Handler, userLevelService userlevelservice.Service) controller.Controller {
	return controller.Controller{
		Method: "GET",
		Path:   fmt.Sprintf("/users/:%s/levels/:%s", controller.ParamUserId, controller.ParamLevelId),
		Middleware: []http.Handler{
			authMiddleware,
		},
		Body: CreateGetBody(exec, userLevelService),
	}
}

func CreateGetBody(exec pipeline.HttpPipeline, userLevelService userlevelservice.Service) http.Handler {
	fetchUserIdCommand := command.CreateFetchUserIdCommand()
	fetchLevelIdCommand := command2.CreateFetchLevelIdCommand()
	serviceCommand := command2.CreateGetUserLevelCommand(userLevelService)
	renderCommand := command2.CreateRenderUserLevelCommand()

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
