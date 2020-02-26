package command

import (
	"github.com/arpb2/C-3PO/pkg/domain/architecture/http"
	httppipeline "github.com/arpb2/C-3PO/pkg/domain/architecture/pipeline"
	model2 "github.com/arpb2/C-3PO/pkg/domain/user_level/model"
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

	ctx.Set(TagUserLevelData, model2.UserLevelData{
		Code:      code,
		Workspace: workspace,
	})
	return nil
}

func CreateFetchCodeCommand() pipeline.Step {
	return &fetchCodeCommand{}
}
