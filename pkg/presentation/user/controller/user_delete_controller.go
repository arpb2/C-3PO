package controller

import (
	"fmt"

	"github.com/arpb2/C-3PO/pkg/presentation/user"

	"github.com/arpb2/C-3PO/pkg/domain/user/repository"

	"github.com/arpb2/C-3PO/pkg/presentation/user/command"

	"github.com/arpb2/C-3PO/pkg/domain/architecture/controller"
	"github.com/arpb2/C-3PO/pkg/domain/architecture/http"
	"github.com/arpb2/C-3PO/pkg/domain/architecture/pipeline"
	gopipeline "github.com/saantiaguilera/go-pipeline"
)

func CreateDeleteController(exec pipeline.HttpPipeline, authMiddleware http.Handler, userRepository repository.UserRepository) controller.Controller {
	return controller.Controller{
		Method: "DELETE",
		Path:   fmt.Sprintf("/users/:%s", user.ParamUserId),
		Middleware: []http.Handler{
			authMiddleware,
		},
		Body: CreateDeleteBody(exec, userRepository),
	}
}

func CreateDeleteBody(exec pipeline.HttpPipeline, userRepository repository.UserRepository) http.Handler {
	fetchUserIdCommand := command.CreateFetchUserIdCommand()
	repositoryCommand := command.CreateDeleteUserCommand(userRepository)
	renderCommand := command.CreateRenderEmptyCommand()

	graph := gopipeline.CreateSequentialStage(
		fetchUserIdCommand,
		repositoryCommand,
		renderCommand,
	)

	return func(ctx *http.Context) {
		exec.Run(ctx, graph)
	}
}
