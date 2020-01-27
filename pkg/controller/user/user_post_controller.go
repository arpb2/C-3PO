package user_controller

import (
	"github.com/arpb2/C-3PO/api/controller"
	"github.com/arpb2/C-3PO/api/http_wrapper"
	"github.com/arpb2/C-3PO/api/pipeline"
	user_service "github.com/arpb2/C-3PO/api/service/user"
	user_command "github.com/arpb2/C-3PO/pkg/command/user"
	user_validation "github.com/arpb2/C-3PO/pkg/validation/user"
	"github.com/saantiaguilera/go-pipeline/pkg/stage/sequential"
)

func CreatePostController(exec pipeline.HttpPipeline, validations []user_validation.Validation, userService user_service.Service) controller.Controller {
	return controller.Controller{
		Method: "POST",
		Path:   "/users",
		Body:   CreatePostBody(exec, validations, userService),
	}
}

func CreatePostBody(exec pipeline.HttpPipeline, validations []user_validation.Validation, userService user_service.Service) http_wrapper.Handler {
	return func(ctx *http_wrapper.Context) {
		fetchUserCommand := user_command.CreateFetchAuthenticatedUserCommand(ctx)
		validateCommand := user_command.CreateValidateParametersCommand(ctx, fetchUserCommand.OutputStream, validations)
		serviceCommand := user_command.CreateCreateUserCommand(ctx, userService, validateCommand.OutputStream)
		renderCommand := user_command.CreateRenderUserCommand(ctx, serviceCommand.OutputStream)

		graph := sequential.CreateSequentialStage(
			fetchUserCommand,
			validateCommand,
			serviceCommand,
			renderCommand,
		)

		exec.Run(ctx, graph)
	}
}
