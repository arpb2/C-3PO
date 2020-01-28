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

func CreateGetController(exec pipeline.HttpPipeline, authMiddleware http.Handler, codeService codeservice.Service) controller.Controller {
	return controller.Controller{
		Method: "GET",
		Path:   "/users/:user_id/codes/:code_id",
		Middleware: []http.Handler{
			authMiddleware,
		},
		Body: CreateGetBody(exec, codeService),
	}
}

func CreateGetBody(exec pipeline.HttpPipeline, codeService codeservice.Service) http.Handler {
	return func(ctx *http.Context) {
		fetchUserIdCommand := usercommand.CreateFetchUserIdCommand(ctx)
		fetchCodeIdCommand := codecommand.CreateFetchCodeIdCommand(ctx)
		serviceCommand := codecommand.CreateGetCodeCommand(ctx, codeService, fetchUserIdCommand.OutputStream, fetchCodeIdCommand.OutputStream)
		renderCommand := codecommand.CreateRenderCodeCommand(ctx, serviceCommand.OutputStream)

		graph := gopipeline.CreateSequentialGroup(
			gopipeline.CreateConcurrentStage(
				fetchUserIdCommand,
				fetchCodeIdCommand,
			),
			gopipeline.CreateSequentialStage(
				serviceCommand,
				renderCommand,
			),
		)

		exec.Run(ctx, graph)
	}
}
