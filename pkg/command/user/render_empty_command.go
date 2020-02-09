package user

import (
	"net/http"

	httppipeline "github.com/arpb2/C-3PO/pkg/pipeline"

	"github.com/saantiaguilera/go-pipeline"
)

type renderEmptyCommand struct{}

func (c *renderEmptyCommand) Name() string {
	return "render_empty_command"
}

func (c *renderEmptyCommand) Run(ctx pipeline.Context) error {
	ctxAware := httppipeline.CreateContextAware(ctx)

	httpWriter, err := ctxAware.GetWriter()

	if err != nil {
		return err
	}

	httpWriter.WriteStatus(http.StatusOK)
	return nil
}

func CreateRenderEmptyCommand() pipeline.Step {
	return &renderEmptyCommand{}
}
