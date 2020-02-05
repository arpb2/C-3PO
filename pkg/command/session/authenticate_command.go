package session

import (
	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/api/model"
	credentialservice "github.com/arpb2/C-3PO/api/service/credential"
	"github.com/arpb2/C-3PO/pkg/command/user"
	"github.com/saantiaguilera/go-pipeline"
)

type authenticateCommand struct {
	service credentialservice.Service
}

func (c *authenticateCommand) Name() string {
	return "authenticate_command"
}

func (c *authenticateCommand) Run(ctx pipeline.Context) error {
	value, exists := ctx.Get(user.TagAuthenticatedUser)

	if !exists {
		return http.CreateInternalError()
	}

	authenticatedUser := value.(model.AuthenticatedUser)

	userId, err := c.service.Retrieve(authenticatedUser.Email, authenticatedUser.Password)

	if err != nil {
		return err
	}

	ctx.Set(user.TagUserId, userId)
	return nil
}

func CreateAuthenticateCommand(service credentialservice.Service) pipeline.Step {
	return &authenticateCommand{
		service: service,
	}
}
