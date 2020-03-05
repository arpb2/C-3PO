package user

import (
	"fmt"

	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/domain/model/user"
	httppipeline "github.com/arpb2/C-3PO/pkg/domain/pipeline"
	"github.com/saantiaguilera/go-pipeline"
)

type fetchCodeUseCase struct {
	CodePartParam      string
	WorkspacePartParam string
}

func (c *fetchCodeUseCase) Name() string {
	return "fetch_user_level_data_usecase"
}

func (c *fetchCodeUseCase) Run(ctx pipeline.Context) error {
	ctxAware := httppipeline.CreateContextAware(ctx)

	httpReader, err := ctxAware.GetReader()

	if err != nil {
		return err
	}

	code, exists := httpReader.GetFormData(c.CodePartParam)

	if !exists {
		return http.CreateBadRequestError(fmt.Sprintf("'%s' part not found", c.CodePartParam))
	}

	workspace, exists := httpReader.GetFormData(c.WorkspacePartParam)

	if !exists {
		return http.CreateBadRequestError(fmt.Sprintf("'%s' part not found", c.WorkspacePartParam))
	}

	ctx.Set(TagUserLevelData, user.LevelData{
		Code:      code,
		Workspace: workspace,
	})
	return nil
}

func CreateFetchCodeUseCase(codePartParam, workspacePartParam string) pipeline.Step {
	return &fetchCodeUseCase{
		CodePartParam:      codePartParam,
		WorkspacePartParam: workspacePartParam,
	}
}
