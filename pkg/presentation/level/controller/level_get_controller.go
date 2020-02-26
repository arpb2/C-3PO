package controller

import (
	"fmt"

	"github.com/arpb2/C-3PO/pkg/presentation/level"

	"github.com/arpb2/C-3PO/pkg/domain/level/repository"

	"github.com/arpb2/C-3PO/pkg/domain/architecture/controller"
	"github.com/arpb2/C-3PO/pkg/domain/architecture/http"
	"github.com/arpb2/C-3PO/pkg/domain/architecture/pipeline"
	"github.com/arpb2/C-3PO/pkg/presentation/level/command"
	gopipeline "github.com/saantiaguilera/go-pipeline"
)

func CreateGetController(exec pipeline.HttpPipeline, levelRepository repository.LevelRepository) controller.Controller {
	return controller.Controller{
		Method: "GET",
		Path:   fmt.Sprintf("/levels/:%s", level.ParamLevelId),
		Body:   CreateGetBody(exec, levelRepository),
	}
}

func CreateGetBody(exec pipeline.HttpPipeline, levelRepository repository.LevelRepository) http.Handler {
	fetchLevelIdCommand := command.CreateFetchLevelIdCommand()
	repositoryCommand := command.CreateGetLevelCommand(levelRepository)
	renderCommand := command.CreateRenderLevelCommand()

	graph := gopipeline.CreateSequentialStage(
		fetchLevelIdCommand,
		repositoryCommand,
		renderCommand,
	)

	return func(ctx *http.Context) {
		exec.Run(ctx, graph)
	}
}
