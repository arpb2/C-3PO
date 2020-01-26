package code_controller

import (
	"github.com/arpb2/C-3PO/api/controller"
	"github.com/arpb2/C-3PO/api/http_wrapper"
	"github.com/arpb2/C-3PO/api/pipeline"
	code_service "github.com/arpb2/C-3PO/api/service/code"
	code_command "github.com/arpb2/C-3PO/internal/command/code"
	user_command "github.com/arpb2/C-3PO/internal/command/user"
	"github.com/saantiaguilera/go-pipeline/pkg/stage/concurrent"
	"github.com/saantiaguilera/go-pipeline/pkg/stage/sequential"
)

func CreateGetController(exec pipeline.HttpPipeline, authMiddleware http_wrapper.Handler, codeService code_service.Service) controller.Controller {
	return controller.Controller{
		Method: "GET",
		Path:   "/users/:user_id/codes/:code_id",
		Middleware: []http_wrapper.Handler{
			authMiddleware,
		},
		Body: CreateGetBody(exec, codeService),
	}
}

func CreateGetBody(exec pipeline.HttpPipeline, codeService code_service.Service) http_wrapper.Handler {
	return func(ctx *http_wrapper.Context) {
		fetchUserIdCommand := user_command.CreateFetchUserIdCommand(ctx)
		fetchCodeIdCommand := code_command.CreateFetchCodeIdCommand(ctx)
		serviceCommand := code_command.CreateGetCodeCommand(ctx, codeService, fetchUserIdCommand.OutputStream, fetchCodeIdCommand.OutputStream)
		renderCommand := code_command.CreateRenderCodeCommand(ctx, serviceCommand.OutputStream)

		graph := sequential.CreateSequentialGroup(
			concurrent.CreateConcurrentStage(
				fetchUserIdCommand,
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
