package controller

import (
	"github.com/arpb2/C-3PO/pkg/domain/architecture/controller"
	"github.com/arpb2/C-3PO/pkg/domain/architecture/http"
	"github.com/arpb2/C-3PO/pkg/domain/architecture/pipeline"
	"github.com/arpb2/C-3PO/pkg/domain/user/service"
	"github.com/arpb2/C-3PO/pkg/presentation/user/command"
	"github.com/arpb2/C-3PO/pkg/presentation/user/validation"
	gopipeline "github.com/saantiaguilera/go-pipeline"
)

func CreatePostController(exec pipeline.HttpPipeline, userService service.Service, validations []validation.Validation) controller.Controller {
	return controller.Controller{
		Method: "POST",
		Path:   "/users",
		Body:   CreatePostBody(exec, userService, validations),
	}
}

func CreatePostBody(exec pipeline.HttpPipeline, userService service.Service, validations []validation.Validation) http.Handler {
	fetchUserCommand := command.CreateFetchAuthenticatedUserCommand()
	validateCommand := command.CreateValidateUserParametersCommand(validations)
	createUserCommand := command.CreateCreateUserCommand(userService)
	renderCommand := command.CreateRenderUserCommand()

	graph := gopipeline.CreateSequentialStage(
		fetchUserCommand,
		validateCommand,
		createUserCommand,
		renderCommand,
	)

	return func(ctx *http.Context) {
		exec.Run(ctx, graph)
	}
}
