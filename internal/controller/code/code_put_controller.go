package code_controller

import (
	controller2 "github.com/arpb2/C-3PO/api/controller"
	"github.com/arpb2/C-3PO/api/http_wrapper"
	"github.com/arpb2/C-3PO/api/pipeline"
	code_service "github.com/arpb2/C-3PO/api/service/code"
	code_command "github.com/arpb2/C-3PO/internal/command/code"
	user_command "github.com/arpb2/C-3PO/internal/command/user"
	"github.com/saantiaguilera/go-pipeline/pkg/stage/concurrent"
	"github.com/saantiaguilera/go-pipeline/pkg/stage/sequential"
)

func CreatePutController(exec pipeline.HttpPipeline, authMiddleware http_wrapper.Handler, codeService code_service.Service) controller2.Controller {
	return controller2.Controller{
		Method: "PUT",
		Path:   "/users/:user_id/codes/:code_id",
		Middleware: []http_wrapper.Handler{
			authMiddleware,
		},
		Body: CreatePutBody(exec, codeService),
	}
}

func CreatePutBody(exec pipeline.HttpPipeline, codeService code_service.Service) http_wrapper.Handler {
	return func(ctx *http_wrapper.Context) {
		fetchUserIdCommand := user_command.CreateFetchUserIdCommand(ctx)
		fetchCodeCommand := code_command.CreateFetchCodeCommand(ctx)
		fetchCodeIdCommand := code_command.CreateFetchCodeIdCommand(ctx)
		serviceCommand := code_command.CreateReplaceCodeCommand(ctx, codeService,
			fetchCodeIdCommand.OutputStream, fetchUserIdCommand.OutputStream, fetchCodeCommand.OutputStream)
		renderCommand := code_command.CreateRenderCodeCommand(ctx, serviceCommand.OutputStream)

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
