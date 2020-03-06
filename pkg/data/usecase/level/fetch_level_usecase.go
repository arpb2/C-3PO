package level

import (
	"github.com/arpb2/C-3PO/pkg/domain/http"
	level2 "github.com/arpb2/C-3PO/pkg/domain/model/level"
	httppipeline "github.com/arpb2/C-3PO/pkg/domain/pipeline"
	"github.com/saantiaguilera/go-pipeline"
)

type fetchLevelUseCase struct{}

func (c *fetchLevelUseCase) Name() string {
	return "fetch_level_usecase"
}

func (c *fetchLevelUseCase) Run(ctx pipeline.Context) error {
	ctxAware := httppipeline.CreateContextAware(ctx)

	httpReader, err := ctxAware.GetReader()

	if err != nil {
		return err
	}

	var level level2.Level
	if err := httpReader.ReadBody(&level); err != nil {
		return http.CreateBadRequestError("malformed body")
	}

	ctx.Set(TagLevel, level)
	return nil
}

func CreateFetchLevelUseCase() pipeline.Step {
	return &fetchLevelUseCase{}
}
