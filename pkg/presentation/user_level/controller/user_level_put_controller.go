package controller

import (
	"fmt"
	"github.com/arpb2/C-3PO/pkg/domain/user_level/service"
	"github.com/arpb2/C-3PO/pkg/presentation/level"
	"github.com/arpb2/C-3PO/pkg/presentation/user"

	"github.com/arpb2/C-3PO/pkg/presentation/user/command"
	userlevelcommand "github.com/arpb2/C-3PO/pkg/presentation/user_level/command"

	"github.com/arpb2/C-3PO/pkg/domain/architecture/controller"
	"github.com/arpb2/C-3PO/pkg/domain/architecture/http"
	"github.com/arpb2/C-3PO/pkg/domain/architecture/pipeline"
	gopipeline "github.com/saantiaguilera/go-pipeline"
)

func CreatePutController(exec pipeline.HttpPipeline, authMiddleware http.Handler, userLevelService service.Service) controller.Controller {
	return controller.Controller{
		Method: "PUT",
		Path:   fmt.Sprintf("/users/:%s/levels/:%s", user.ParamUserId, level.ParamLevelId),
		Middleware: []http.Handler{
			authMiddleware,
		},
		Body: CreatePutBody(exec, userLevelService),
	}
}

func CreatePutBody(exec pipeline.HttpPipeline, userLevelService service.Service) http.Handler {
	fetchUserIdCommand := command.CreateFetchUserIdCommand()
	fetchCodeCommand := userlevelcommand.CreateFetchCodeCommand()
	fetchLevelIdCommand := userlevelcommand.CreateFetchLevelIdCommand()
	serviceCommand := userlevelcommand.CreateWriteUserLevelCommand(userLevelService)
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
