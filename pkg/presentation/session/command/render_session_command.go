package command

import (
	"net/http"

	httppipeline "github.com/arpb2/C-3PO/pkg/infra/pipeline"

	"github.com/saantiaguilera/go-pipeline"

	httpwrapper "github.com/arpb2/C-3PO/pkg/domain/http"
)

type renderSessionCommand struct{}

func (c *renderSessionCommand) Name() string {
	return "render_session_command"
}

func (c *renderSessionCommand) Run(ctx pipeline.Context) error {
	ctxAware := httppipeline.CreateContextAware(ctx)

	httpWriter, errWriter := ctxAware.GetWriter()
	session, errSession := ctxAware.GetSession(TagSession)

	if errWriter != nil || errSession != nil {
		return httpwrapper.CreateInternalError()
	}

	httpWriter.WriteJson(http.StatusOK, session)
	return nil
}

func CreateRenderSessionCommand() pipeline.Step {
	return &renderSessionCommand{}
}
