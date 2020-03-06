package level

import (
	"github.com/arpb2/C-3PO/pkg/data/repository/level"
	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/saantiaguilera/go-pipeline"
)

type getLevelUseCase struct {
	repository level.Repository
}

func (c *getLevelUseCase) Name() string {
	return "get_level_usecase"
}

func (c *getLevelUseCase) Run(ctx pipeline.Context) error {
	levelId, existsUserId := ctx.GetUInt(TagLevelId)

	if !existsUserId {
		return http.CreateInternalError()
	}

	l, err := c.repository.GetLevel(levelId)

	if err != nil {
		return err
	}

	ctx.Set(TagLevel, l)
	return nil
}

func CreateGetLevelUseCase(repository level.Repository) pipeline.Step {
	return &getLevelUseCase{
		repository: repository,
	}
}
