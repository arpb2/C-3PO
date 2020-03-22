package user

import (
	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/domain/model/user"
	httppipeline "github.com/arpb2/C-3PO/pkg/domain/pipeline"
	"github.com/saantiaguilera/go-pipeline"
)

type fetchCodeUseCase struct{}

func (c *fetchCodeUseCase) Name() string {
	return "fetch_user_level_data_usecase"
}

func (c *fetchCodeUseCase) Run(ctx pipeline.Context) error {
	ctxAware := httppipeline.CreateContextAware(ctx)

	httpReader, err := ctxAware.GetReader()

	if err != nil {
		return err
	}

	var ulData user.LevelData
	err = httpReader.ReadBody(&ulData)

	if err != nil {
		return http.CreateBadRequestError("error reading user level json")
	}

	ctx.Set(TagUserLevelData, ulData)
	return nil
}

func CreateFetchCodeUseCase() pipeline.Step {
	return &fetchCodeUseCase{}
}
