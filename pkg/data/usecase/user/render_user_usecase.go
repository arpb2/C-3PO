package user

import (
	"net/http"

	ctxaware "github.com/arpb2/C-3PO/pkg/domain/pipeline"

	"github.com/saantiaguilera/go-pipeline"

	httpwrapper "github.com/arpb2/C-3PO/pkg/domain/http"
)

type renderUserUseCase struct{}

func (c *renderUserUseCase) Name() string {
	return "render_user_usecase"
}

func (c *renderUserUseCase) Run(ctx pipeline.Context) error {
	ctxAware := ctxaware.CreateContextAware(ctx)

	httpWriter, errWriter := ctxAware.GetWriter()
	user, errUser := ctxAware.GetUser(TagUser)

	if errWriter != nil || errUser != nil {
		return httpwrapper.CreateInternalError()
	}

	httpWriter.WriteJson(http.StatusOK, user)
	return nil
}

func CreateRenderUserUseCase() pipeline.Step {
	return &renderUserUseCase{}
}
