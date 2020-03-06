package user

import (
	"github.com/arpb2/C-3PO/pkg/data/repository/user"
	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/saantiaguilera/go-pipeline"
)

type getUserUseCase struct {
	repository user.Repository
}

func (c *getUserUseCase) Name() string {
	return "get_user_usecase"
}

func (c *getUserUseCase) Run(ctx pipeline.Context) error {
	userId, existsUserId := ctx.GetUInt(TagUserId)

	if !existsUserId {
		return http.CreateInternalError()
	}

	user, err := c.repository.GetUser(userId)

	if err != nil {
		return err
	}

	ctx.Set(TagUser, user)
	return nil
}

func CreateGetUserUseCase(repository user.Repository) pipeline.Step {
	return &getUserUseCase{
		repository: repository,
	}
}
