package user

import (
	"github.com/arpb2/C-3PO/pkg/data/repository/user"
	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/saantiaguilera/go-pipeline"
)

type deleteUserUseCase struct {
	repository user.Repository
}

func (c *deleteUserUseCase) Name() string {
	return "delete_user_usecase"
}

func (c *deleteUserUseCase) Run(ctx pipeline.Context) error {
	userId, exists := ctx.GetUInt(TagUserId)

	if !exists {
		return http.CreateInternalError()
	}

	return c.repository.DeleteUser(userId)
}

func CreateDeleteUserUseCase(repository user.Repository) pipeline.Step {
	return &deleteUserUseCase{
		repository: repository,
	}
}
