package code

import (
	"github.com/arpb2/C-3PO/api/controller"
	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/api/pipeline"
	codeservice "github.com/arpb2/C-3PO/api/service/code"
	codecommand "github.com/arpb2/C-3PO/pkg/command/code"
	usercommand "github.com/arpb2/C-3PO/pkg/command/user"
	"github.com/saantiaguilera/go-pipeline/pkg/stage/concurrent"
	"github.com/saantiaguilera/go-pipeline/pkg/stage/sequential"
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
	return func(ctx *http.Context) {
		fetchUserIdCommand := usercommand.CreateFetchUserIdCommand(ctx)
		fetchCodeCommand := codecommand.CreateFetchCodeCommand(ctx)
		serviceCommand := codecommand.CreateCreateCodeCommand(ctx, codeService, fetchUserIdCommand.OutputStream, fetchCodeCommand.OutputStream)
		renderCommand := codecommand.CreateRenderCodeCommand(ctx, serviceCommand.OutputStream)

		graph := sequential.CreateSequentialGroup(
			concurrent.CreateConcurrentStage(
				fetchUserIdCommand,
				fetchCodeCommand,
			),
			sequential.CreateSequentialStage(
				serviceCommand,
				renderCommand,
			),
		)

		exec.Run(ctx, graph)
	}
}