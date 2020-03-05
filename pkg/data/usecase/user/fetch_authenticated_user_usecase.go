package user

import (
	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/domain/model/user"
	httppipeline "github.com/arpb2/C-3PO/pkg/domain/pipeline"
	"github.com/saantiaguilera/go-pipeline"
)

type fetchAuthenticatedUserUseCase struct{}

func (c *fetchAuthenticatedUserUseCase) Name() string {
	return "fetch_authenticated_user_usecase"
}

func (c *fetchAuthenticatedUserUseCase) Run(ctx pipeline.Context) error {
	ctxAware := httppipeline.CreateContextAware(ctx)

	httpReader, err := ctxAware.GetReader()

	if err != nil {
		return err
	}

	var authenticatedUser user.AuthenticatedUser
	if err := httpReader.ReadBody(&authenticatedUser); err != nil {
		return http.CreateBadRequestError("malformed body")
	}

	ctx.Set(TagAuthenticatedUser, authenticatedUser)
	return nil
}

func CreateFetchAuthenticatedUserUseCase() pipeline.Step {
	return &fetchAuthenticatedUserUseCase{}
}
