package controller

import (
	"github.com/arpb2/C-3PO/pkg/domain/controller"
	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/domain/pipeline"
	userservice "github.com/arpb2/C-3PO/pkg/domain/service/user"
	"github.com/arpb2/C-3PO/pkg/presentation/user/command"
	"github.com/arpb2/C-3PO/pkg/presentation/user/validation"
	gopipeline "github.com/saantiaguilera/go-pipeline"
)

func CreatePostController(exec pipeline.HttpPipeline, userService userservice.Service, validations []validation.Validation) controller.Controller {
	return controller.Controller{
		Method: "POST",
		Path:   "/users",
		Body:   CreatePostBody(exec, userService, validations),
	}
}

func CreatePostBody(exec pipeline.HttpPipeline, userService userservice.Service, validations []validation.Validation) http.Handler {
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
