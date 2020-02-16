package user_level

import (
	"github.com/arpb2/C-3PO/api/controller"
	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/api/pipeline"
	userlevelservice "github.com/arpb2/C-3PO/api/service/user_level"
	usercommand "github.com/arpb2/C-3PO/pkg/command/user"
	userlevelcommand "github.com/arpb2/C-3PO/pkg/command/user_level"
	gopipeline "github.com/saantiaguilera/go-pipeline"
)

func CreatePostController(exec pipeline.HttpPipeline, authMiddleware http.Handler, userLevelService userlevelservice.Service) controller.Controller {
	return controller.Controller{
		Method: "POST",
		Path:   "/users/:user_id/levels",
		Middleware: []http.Handler{
			authMiddleware,
		},
		Body: CreatePostBody(exec, userLevelService),
	}
}

func CreatePostBody(exec pipeline.HttpPipeline, userLevelService userlevelservice.Service) http.Handler {
	fetchUserIdCommand := usercommand.CreateFetchUserIdCommand()
	fetchCodeCommand := userlevelcommand.CreateFetchCodeCommand()
	serviceCommand := userlevelcommand.CreateCreateUserLevelCommand(userLevelService)
	renderCommand := userlevelcommand.CreateRenderUserLevelCommand()

	graph := gopipeline.CreateSequentialGroup(
		gopipeline.CreateConcurrentStage(
			fetchUserIdCommand,
			fetchCodeCommand,
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
