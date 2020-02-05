package user

import (
	"net/http"

	"github.com/arpb2/C-3PO/pkg/command"
	"github.com/saantiaguilera/go-pipeline"

	httpwrapper "github.com/arpb2/C-3PO/api/http"
)

type renderEmptyCommand struct{}

func (c *renderEmptyCommand) Name() string {
	return "render_empty_command"
}

func (c *renderEmptyCommand) Run(ctx pipeline.Context) error {
	httpWriter, existsWriter := ctx.Get(command.TagHttpWriter)

	if !existsWriter {
		return httpwrapper.CreateInternalError()
	}

	httpWriter.(httpwrapper.Writer).WriteStatus(http.StatusOK)
	return nil
}

func CreateRenderEmptyCommand() pipeline.Step {
	return &renderEmptyCommand{}
}
