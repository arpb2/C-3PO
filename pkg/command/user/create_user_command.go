package user

import (
	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/api/model"
	userservice "github.com/arpb2/C-3PO/api/service/user"
	"github.com/saantiaguilera/go-pipeline"
)

type createUserCommand struct {
	service userservice.Service
}

func (c *createUserCommand) Name() string {
	return "create_user_command"
}

func (c *createUserCommand) Run(ctx pipeline.Context) error {
	value, exists := ctx.Get(TagAuthenticatedUser)

	if !exists {
		return http.CreateInternalError()
	}

	authenticatedUser := value.(model.AuthenticatedUser)

	user, err := c.service.CreateUser(&authenticatedUser)

	if err != nil {
		return err
	}

	if user == nil {
		return http.CreateInternalError()
	}

	ctx.Set(TagUser, *user)
	return nil
}

func CreateCreateUserCommand(service userservice.Service) pipeline.Step {
	return &createUserCommand{
		service: service,
	}
}
