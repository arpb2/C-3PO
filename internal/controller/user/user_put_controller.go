package user_controller

import (
	controller2 "github.com/arpb2/C-3PO/api/controller"
	"github.com/arpb2/C-3PO/api/http_wrapper"
	"github.com/arpb2/C-3PO/api/pipeline"
	user_service "github.com/arpb2/C-3PO/api/service/user"
	user_command "github.com/arpb2/C-3PO/internal/command/user"
	user_validation "github.com/arpb2/C-3PO/internal/validation/user"
	"github.com/saantiaguilera/go-pipeline/pkg/stage/concurrent"
	"github.com/saantiaguilera/go-pipeline/pkg/stage/sequential"
)

func CreatePutController(exec pipeline.HttpPipeline, validations []user_validation.Validation, authMiddleware http_wrapper.Handler, userService user_service.Service) controller2.Controller {
	return controller2.Controller{
		Method: "PUT",
		Path:   "/users/:user_id",
		Middleware: []http_wrapper.Handler{
			authMiddleware,
		},
		Body: CreatePutBody(exec, validations, userService),
	}
}

func CreatePutBody(exec pipeline.HttpPipeline, validations []user_validation.Validation, userService user_service.Service) http_wrapper.Handler {
	return func(ctx *http_wrapper.Context) {
		fetchUserIdCommand := user_command.CreateFetchUserIdCommand(ctx)
		fetchUserCommand := user_command.CreateFetchAuthenticatedUserCommand(ctx)
		validateCommand := user_command.CreateValidateParametersCommand(ctx, fetchUserCommand.OutputStream, validations)
		serviceCommand := user_command.CreateUpdateUserCommand(ctx, userService, fetchUserIdCommand.OutputStream, validateCommand.OutputStream)
		renderCommand := user_command.CreateRenderUserCommand(ctx, serviceCommand.OutputStream)

		graph := sequential.CreateSequentialGroup(
			concurrent.CreateConcurrentGroup(
				sequential.CreateSequentialStage(
					fetchUserIdCommand,
				),
				sequential.CreateSequentialStage(
					fetchUserCommand,
					validateCommand,
				),
			),
			sequential.CreateSequentialStage(
				serviceCommand,
				renderCommand,
			),
		)

		exec.Run(ctx, graph)
	}
}
