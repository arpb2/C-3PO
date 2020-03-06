package user

import (
	"github.com/arpb2/C-3PO/pkg/data/repository/user"
	"github.com/arpb2/C-3PO/pkg/domain/http"
	httppipeline "github.com/arpb2/C-3PO/pkg/domain/pipeline"
	"github.com/saantiaguilera/go-pipeline"
)

type updateUserUseCase struct {
	repository user.Repository
}

func (c *updateUserUseCase) Name() string {
	return "update_user_usecase"
}

func (c *updateUserUseCase) Run(ctx pipeline.Context) error {
	ctxAware := httppipeline.CreateContextAware(ctx)

	authenticatedUser, errAuthenticatedUser := ctxAware.GetAuthenticatedUser(TagAuthenticatedUser)
	userId, existsUserId := ctx.GetUInt(TagUserId)

	if errAuthenticatedUser != nil || !existsUserId {
		return http.CreateInternalError()
	}

	authenticatedUser.Id = userId

	user, err := c.repository.UpdateUser(authenticatedUser)

	if err != nil {
		return err
	}

	ctx.Set(TagUser, user)
	return nil
}

func CreateUpdateUserUseCase(repository user.Repository) pipeline.Step {
	return &updateUserUseCase{
		repository: repository,
	}
}
