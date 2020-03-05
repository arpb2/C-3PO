package user

import (
	"github.com/arpb2/C-3PO/pkg/data/repository/user"
	httppipeline "github.com/arpb2/C-3PO/pkg/domain/pipeline"
	"github.com/saantiaguilera/go-pipeline"
)

type createUserUseCase struct {
	repository user.Repository
}

func (c *createUserUseCase) Name() string {
	return "create_user_usecase"
}

func (c *createUserUseCase) Run(ctx pipeline.Context) error {
	ctxAware := httppipeline.CreateContextAware(ctx)

	authenticatedUser, err := ctxAware.GetAuthenticatedUser(TagAuthenticatedUser)

	if err != nil {
		return err
	}

	user, err := c.repository.CreateUser(authenticatedUser)

	if err != nil {
		return err
	}

	ctx.Set(TagUser, user)
	return nil
}

func CreateCreateUserUseCase(repository user.Repository) pipeline.Step {
	return &createUserUseCase{
		repository: repository,
	}
}
