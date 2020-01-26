package user

import (
	"github.com/arpb2/C-3PO/src/api/command/user_command"
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/controller/user/user_validation"
	"github.com/arpb2/C-3PO/src/api/executor"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/service"
	"github.com/saantiaguilera/go-pipeline/pkg/stage/concurrent"
	"github.com/saantiaguilera/go-pipeline/pkg/stage/sequential"
)

func CreatePutController(exec executor.HttpPipeline, validations []user_validation.Validation, authMiddleware http_wrapper.Handler, userService service.UserService) controller.Controller {
	return controller.Controller{
		Method: "PUT",
		Path:   "/users/:user_id",
		Middleware: []http_wrapper.Handler{
			authMiddleware,
		},
		Body:   CreatePutBody(exec, validations, userService),
	}
}

func CreatePutBody(exec executor.HttpPipeline, validations []user_validation.Validation, userService service.UserService) http_wrapper.Handler {
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