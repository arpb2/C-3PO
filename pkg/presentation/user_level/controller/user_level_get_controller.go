package controller

import (
	"fmt"

	"github.com/arpb2/C-3PO/pkg/domain/user_level/repository"
	"github.com/arpb2/C-3PO/pkg/presentation/level"
	"github.com/arpb2/C-3PO/pkg/presentation/user"

	"github.com/arpb2/C-3PO/pkg/presentation/user/command"
	userlevelcommand "github.com/arpb2/C-3PO/pkg/presentation/user_level/command"

	"github.com/arpb2/C-3PO/pkg/domain/architecture/controller"
	"github.com/arpb2/C-3PO/pkg/domain/architecture/http"
	"github.com/arpb2/C-3PO/pkg/domain/architecture/pipeline"
	gopipeline "github.com/saantiaguilera/go-pipeline"
)

func CreateGetController(exec pipeline.HttpPipeline, authMiddleware http.Handler, userLevelRepository repository.UserLevelRepository) controller.Controller {
	return controller.Controller{
		Method: "GET",
		Path:   fmt.Sprintf("/users/:%s/levels/:%s", user.ParamUserId, level.ParamLevelId),
		Middleware: []http.Handler{
			authMiddleware,
		},
		Body: CreateGetBody(exec, userLevelRepository),
	}
}

func CreateGetBody(exec pipeline.HttpPipeline, userLevelRepository repository.UserLevelRepository) http.Handler {
	fetchUserIdCommand := command.CreateFetchUserIdCommand()
	fetchLevelIdCommand := userlevelcommand.CreateFetchLevelIdCommand()
	repositoryCommand := userlevelcommand.CreateGetUserLevelCommand(userLevelRepository)
	renderCommand := userlevelcommand.CreateRenderUserLevelCommand()

	graph := gopipeline.CreateSequentialGroup(
		gopipeline.CreateConcurrentStage(
			fetchUserIdCommand,
			fetchLevelIdCommand,
		),
		gopipeline.CreateSequentialStage(
			repositoryCommand,
			renderCommand,
		),
	)

	return func(ctx *http.Context) {
		exec.Run(ctx, graph)
	}
}
