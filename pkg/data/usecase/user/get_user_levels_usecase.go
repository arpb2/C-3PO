package user

import (
	user2 "github.com/arpb2/C-3PO/pkg/data/repository/user"
	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/saantiaguilera/go-pipeline"
)

type getUserLevelsUseCase struct {
	repository user2.LevelRepository
}

func (c *getUserLevelsUseCase) Name() string {
	return "get_user_levels_usecase"
}

func (c *getUserLevelsUseCase) Run(ctx pipeline.Context) error {
	userId, existsUserId := ctx.GetUInt(TagUserId)

	if !existsUserId {
		return http.CreateInternalError()
	}

	userLevels, err := c.repository.GetUserLevels(userId)

	if err != nil {
		return err
	}

	ctx.Set(TagUserLevels, userLevels)
	return nil
}

func CreateGetUserLevelsUseCase(repository user2.LevelRepository) pipeline.Step {
	return &getUserLevelsUseCase{
		repository: repository,
	}
}
