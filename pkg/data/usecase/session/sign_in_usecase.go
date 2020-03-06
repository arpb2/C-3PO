package session

import (
	"github.com/arpb2/C-3PO/pkg/data/repository/session"
	"github.com/arpb2/C-3PO/pkg/data/usecase/user"
	httppipeline "github.com/arpb2/C-3PO/pkg/domain/pipeline"
	"github.com/saantiaguilera/go-pipeline"
)

type authenticateUseCase struct {
	repository session.CredentialRepository
}

func (c *authenticateUseCase) Name() string {
	return "authenticate_usecase"
}

func (c *authenticateUseCase) Run(ctx pipeline.Context) error {
	ctxAware := httppipeline.CreateContextAware(ctx)

	authenticatedUser, err := ctxAware.GetAuthenticatedUser(user.TagAuthenticatedUser)

	if err != nil {
		return err
	}

	userId, err := c.repository.GetUserId(authenticatedUser.Email, authenticatedUser.Password)

	if err != nil {
		return err
	}

	ctx.Set(user.TagUserId, userId)
	return nil
}

func CreateAuthenticateUseCase(repository session.CredentialRepository) pipeline.Step {
	return &authenticateUseCase{
		repository: repository,
	}
}
