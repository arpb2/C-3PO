package code

import (
	"net/http"

	httppipeline "github.com/arpb2/C-3PO/pkg/pipeline"

	"github.com/saantiaguilera/go-pipeline"

	httpwrapper "github.com/arpb2/C-3PO/api/http"
)

type renderCodeCommand struct{}

func (c *renderCodeCommand) Name() string {
	return "render_code_command"
}

func (c *renderCodeCommand) Run(ctx pipeline.Context) error {
	ctxAware := httppipeline.CreateContextAware(ctx)

	httpWriter, errWriter := ctxAware.GetWriter()
	code, errCode := ctxAware.GetCode(TagCode)

	if errWriter != nil || errCode != nil {
		return httpwrapper.CreateInternalError()
	}

	httpWriter.WriteJson(http.StatusOK, code)
	return nil
}

func CreateRenderCodeCommand() pipeline.Step {
	return &renderCodeCommand{}
}
