package code

import (
	"github.com/arpb2/C-3PO/api/controller"
	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/api/pipeline"
	codeservice "github.com/arpb2/C-3PO/api/service/code"
	codecommand "github.com/arpb2/C-3PO/pkg/command/code"
	usercommand "github.com/arpb2/C-3PO/pkg/command/user"
	gopipeline "github.com/saantiaguilera/go-pipeline"
)

func CreatePostController(exec pipeline.HttpPipeline, authMiddleware http.Handler, codeService codeservice.Service) controller.Controller {
	return controller.Controller{
		Method: "POST",
		Path:   "/users/:user_id/codes",
		Middleware: []http.Handler{
			authMiddleware,
		},
		Body: CreatePostBody(exec, codeService),
	}
}

func CreatePostBody(exec pipeline.HttpPipeline, codeService codeservice.Service) http.Handler {
	fetchUserIdCommand := usercommand.CreateFetchUserIdCommand()
	fetchCodeCommand := codecommand.CreateFetchCodeCommand()
	serviceCommand := codecommand.CreateCreateCodeCommand(codeService)
	renderCommand := codecommand.CreateRenderCodeCommand()

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
