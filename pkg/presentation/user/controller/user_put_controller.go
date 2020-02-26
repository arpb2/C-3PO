package controller

import (
	"fmt"

	"github.com/arpb2/C-3PO/pkg/presentation/user"

	"github.com/arpb2/C-3PO/pkg/domain/user/repository"

	"github.com/arpb2/C-3PO/pkg/presentation/user/command"
	"github.com/arpb2/C-3PO/pkg/presentation/user/validation"

	"github.com/arpb2/C-3PO/pkg/domain/architecture/controller"
	"github.com/arpb2/C-3PO/pkg/domain/architecture/http"
	"github.com/arpb2/C-3PO/pkg/domain/architecture/pipeline"
	gopipeline "github.com/saantiaguilera/go-pipeline"
)

func CreatePutController(exec pipeline.HttpPipeline, authMiddleware http.Handler, userRepository repository.UserRepository, validations []validation.Validation) controller.Controller {
	return controller.Controller{
		Method: "PUT",
		Path:   fmt.Sprintf("/users/:%s", user.ParamUserId),
		Middleware: []http.Handler{
			authMiddleware,
		},
		Body: CreatePutBody(exec, userRepository, validations),
	}
}

func CreatePutBody(exec pipeline.HttpPipeline, userRepository repository.UserRepository, validations []validation.Validation) http.Handler {
	fetchUserIdCommand := command.CreateFetchUserIdCommand()
	fetchUserCommand := command.CreateFetchAuthenticatedUserCommand()
	validateCommand := command.CreateValidateUserParametersCommand(validations)
	repositoryCommand := command.CreateUpdateUserCommand(userRepository)
	renderCommand := command.CreateRenderUserCommand()

	graph := gopipeline.CreateSequentialGroup(
		gopipeline.CreateConcurrentGroup(
			gopipeline.CreateSequentialStage(
				fetchUserIdCommand,
			),
			gopipeline.CreateSequentialStage(
				fetchUserCommand,
				validateCommand,
			),
		),
		gopipeline.CreateSequentialStage(
			repositoryCommand,
			renderCommand,
		),
	)

	return func(ctx *http.Context) {
		exec.Run(ctx, graph)
	}
}
