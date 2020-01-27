package code_controller

import (
	"github.com/arpb2/C-3PO/api/controller"
	"github.com/arpb2/C-3PO/api/http_wrapper"
	"github.com/arpb2/C-3PO/api/pipeline"
	code_service "github.com/arpb2/C-3PO/api/service/code"
	code_command "github.com/arpb2/C-3PO/pkg/command/code"
	user_command "github.com/arpb2/C-3PO/pkg/command/user"
	"github.com/saantiaguilera/go-pipeline/pkg/stage/concurrent"
	"github.com/saantiaguilera/go-pipeline/pkg/stage/sequential"
)

func CreatePostController(exec pipeline.HttpPipeline, authMiddleware http_wrapper.Handler, codeService code_service.Service) controller.Controller {
	return controller.Controller{
		Method: "POST",
		Path:   "/users/:user_id/codes",
		Middleware: []http_wrapper.Handler{
			authMiddleware,
		},
		Body: CreatePostBody(exec, codeService),
	}
}

func CreatePostBody(exec pipeline.HttpPipeline, codeService code_service.Service) http_wrapper.Handler {
	return func(ctx *http_wrapper.Context) {
		fetchUserIdCommand := user_command.CreateFetchUserIdCommand(ctx)
		fetchCodeCommand := code_command.CreateFetchCodeCommand(ctx)
		serviceCommand := code_command.CreateCreateCodeCommand(ctx, codeService, fetchUserIdCommand.OutputStream, fetchCodeCommand.OutputStream)
		renderCommand := code_command.CreateRenderCodeCommand(ctx, serviceCommand.OutputStream)

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
