package user

import (
	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/api/model"
	"github.com/arpb2/C-3PO/pkg/command"
	"github.com/saantiaguilera/go-pipeline"
)

type fetchAuthenticatedUserCommand struct{}

func (c *fetchAuthenticatedUserCommand) Name() string {
	return "fetch_user_command"
}

func (c *fetchAuthenticatedUserCommand) Run(ctx pipeline.Context) error {
	httpReader, exists := ctx.Get(command.TagHttpReader)

	if !exists {
		return http.CreateInternalError()
	}

	var authenticatedUser model.AuthenticatedUser

	if err := httpReader.(http.Reader).ReadBody(&authenticatedUser); err != nil {
		return http.CreateBadRequestError("malformed body")
	}

	ctx.Set(TagAuthenticatedUser, authenticatedUser)
	return nil
}

func CreateFetchAuthenticatedUserCommand() pipeline.Step {
	return &fetchAuthenticatedUserCommand{}
}
