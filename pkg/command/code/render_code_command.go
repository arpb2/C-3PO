package code

import (
	"net/http"

	"github.com/arpb2/C-3PO/pkg/command"
	"github.com/saantiaguilera/go-pipeline"

	httpwrapper "github.com/arpb2/C-3PO/api/http"
)

type renderCodeCommand struct{}

func (c *renderCodeCommand) Name() string {
	return "render_code_command"
}

func (c *renderCodeCommand) Run(ctx pipeline.Context) error {
	httpWriter, existsWriter := ctx.Get(command.TagHttpWriter)
	code, existsCode := ctx.Get(TagCode)

	if !existsWriter || !existsCode {
		return httpwrapper.CreateInternalError()
	}

	httpWriter.(httpwrapper.Writer).WriteJson(http.StatusOK, code)
	return nil
}

func CreateRenderCodeCommand() pipeline.Step {
	return &renderCodeCommand{}
}
