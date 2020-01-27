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

func CreatePutController(exec pipeline.HttpPipeline, authMiddleware http.Handler, codeService codeservice.Service) controller.Controller {
	return controller.Controller{
		Method: "PUT",
		Path:   "/users/:user_id/codes/:code_id",
		Middleware: []http.Handler{
			authMiddleware,
		},
		Body: CreatePutBody(exec, codeService),
	}
}

func CreatePutBody(exec pipeline.HttpPipeline, codeService codeservice.Service) http.Handler {
	return func(ctx *http.Context) {
		fetchUserIdCommand := usercommand.CreateFetchUserIdCommand(ctx)
		fetchCodeCommand := codecommand.CreateFetchCodeCommand(ctx)
		fetchCodeIdCommand := codecommand.CreateFetchCodeIdCommand(ctx)
		serviceCommand := codecommand.CreateReplaceCodeCommand(ctx, codeService,
			fetchCodeIdCommand.OutputStream, fetchUserIdCommand.OutputStream, fetchCodeCommand.OutputStream)
		renderCommand := codecommand.CreateRenderCodeCommand(ctx, serviceCommand.OutputStream)

		graph := sequential.CreateSequentialGroup(
			concurrent.CreateConcurrentStage(
				fetchUserIdCommand,
				fetchCodeCommand,
				fetchCodeIdCommand,
			),
			sequential.CreateSequentialStage(
				serviceCommand,
				renderCommand,
			),
		)

		exec.Run(ctx, graph)
	}
}
