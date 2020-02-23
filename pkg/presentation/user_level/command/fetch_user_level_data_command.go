package command

import (
	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/domain/model"
	httppipeline "github.com/arpb2/C-3PO/pkg/infra/pipeline"
	"github.com/saantiaguilera/go-pipeline"
)

type fetchCodeCommand struct{}

func (c *fetchCodeCommand) Name() string {
	return "fetch_user_level_data_command"
}

func (c *fetchCodeCommand) Run(ctx pipeline.Context) error {
	ctxAware := httppipeline.CreateContextAware(ctx)

	httpReader, err := ctxAware.GetReader()

	if err != nil {
		return err
	}

	code, exists := httpReader.GetFormData("code")

	if !exists {
		return http.CreateBadRequestError("'code' part not found")
	}

	workspace, exists := httpReader.GetFormData("workspace")

	if !exists {
		return http.CreateBadRequestError("'workspace' part not found")
	}

	ctx.Set(TagUserLevelData, model.UserLevelData{
		Code:      code,
		Workspace: workspace,
	})
	return nil
}

func CreateFetchCodeCommand() pipeline.Step {
	return &fetchCodeCommand{}
}
