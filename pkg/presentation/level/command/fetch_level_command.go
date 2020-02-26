package command

import (
	"github.com/arpb2/C-3PO/pkg/domain/infrastructure/http"
	httppipeline "github.com/arpb2/C-3PO/pkg/domain/infrastructure/pipeline"
	model2 "github.com/arpb2/C-3PO/pkg/domain/level/model"
	"github.com/saantiaguilera/go-pipeline"
)

type fetchLevelCommand struct{}

func (c *fetchLevelCommand) Name() string {
	return "fetch_level_command"
}

func (c *fetchLevelCommand) Run(ctx pipeline.Context) error {
	ctxAware := httppipeline.CreateContextAware(ctx)

	httpReader, err := ctxAware.GetReader()

	if err != nil {
		return err
	}

	var level model2.Level
	if err := httpReader.ReadBody(&level); err != nil {
		return http.CreateBadRequestError("malformed body")
	}

	ctx.Set(TagLevel, level)
	return nil
}

func CreateFetchLevelCommand() pipeline.Step {
	return &fetchLevelCommand{}
}
