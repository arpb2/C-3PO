package session

import (
	"net/http"

	"github.com/arpb2/C-3PO/pkg/command"
	"github.com/saantiaguilera/go-pipeline"

	httpwrapper "github.com/arpb2/C-3PO/api/http"
)

type renderSessionCommand struct{}

func (c *renderSessionCommand) Name() string {
	return "render_session_command"
}

func (c *renderSessionCommand) Run(ctx pipeline.Context) error {
	httpWriter, existsWriter := ctx.Get(command.TagHttpWriter)
	session, existsSession := ctx.Get(TagSession)

	if !existsWriter || !existsSession {
		return httpwrapper.CreateInternalError()
	}

	httpWriter.(httpwrapper.Writer).WriteJson(http.StatusOK, session)
	return nil
}

func CreateRenderSessionCommand() pipeline.Step {
	return &renderSessionCommand{}
}
