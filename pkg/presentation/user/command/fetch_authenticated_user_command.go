package command

import (
	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/domain/model"
	httppipeline "github.com/arpb2/C-3PO/pkg/infra/pipeline"
	"github.com/saantiaguilera/go-pipeline"
)

type fetchAuthenticatedUserCommand struct{}

func (c *fetchAuthenticatedUserCommand) Name() string {
	return "fetch_authenticated_user_command"
}

func (c *fetchAuthenticatedUserCommand) Run(ctx pipeline.Context) error {
	ctxAware := httppipeline.CreateContextAware(ctx)

	httpReader, err := ctxAware.GetReader()

	if err != nil {
		return err
	}

	var authenticatedUser model.AuthenticatedUser
	if err := httpReader.ReadBody(&authenticatedUser); err != nil {
		return http.CreateBadRequestError("malformed body")
	}

	ctx.Set(TagAuthenticatedUser, authenticatedUser)
	return nil
}

func CreateFetchAuthenticatedUserCommand() pipeline.Step {
	return &fetchAuthenticatedUserCommand{}
}
