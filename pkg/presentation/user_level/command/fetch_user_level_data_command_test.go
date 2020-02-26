package command_test

import (
	"testing"

	pipeline2 "github.com/arpb2/C-3PO/pkg/domain/architecture/pipeline"
	model2 "github.com/arpb2/C-3PO/pkg/domain/user_level/model"

	"github.com/arpb2/C-3PO/pkg/presentation/user_level/command"

	"github.com/arpb2/C-3PO/pkg/domain/architecture/http"
	http2 "github.com/arpb2/C-3PO/test/mock/http"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestFetchCodeCommand_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	cmd := command.CreateFetchCodeCommand()

	name := cmd.Name()

	assert.Equal(t, "fetch_user_level_data_command", name)
}

func TestFetchCodeCommand_GivenOneAndAContextWithoutAReader_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	cmd := command.CreateFetchCodeCommand()

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestFetchCodeCommand_GivenOneAndAReaderWithoutCodePart_WhenRunning_Then400(t *testing.T) {
	reader := new(http2.MockReader)
	reader.On("GetFormData", "code").Return("", false)
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline2.TagHttpReader, reader)
	cmd := command.CreateFetchCodeCommand()

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateBadRequestError("'code' part not found"), err)
	reader.AssertExpectations(t)
}

func TestFetchCodeCommand_GivenOneAndAReaderWithoutWorkspacePart_WhenRunning_Then400(t *testing.T) {
	reader := new(http2.MockReader)
	reader.On("GetFormData", "code").Return("code", true)
	reader.On("GetFormData", "workspace").Return("", false)
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline2.TagHttpReader, reader)
	cmd := command.CreateFetchCodeCommand()

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateBadRequestError("'workspace' part not found"), err)
	reader.AssertExpectations(t)
}

func TestFetchCodeCommand_GivenOne_WhenRunning_ThenRawCodeIsAddedToContext(t *testing.T) {
	reader := new(http2.MockReader)
	reader.On("GetFormData", "code").Return("test raw code", true)
	reader.On("GetFormData", "workspace").Return("test workspace", true)
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline2.TagHttpReader, reader)
	cmd := command.CreateFetchCodeCommand()

	err := cmd.Run(ctx)
	val, exists := ctx.Get(command.TagUserLevelData)

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, model2.UserLevelData{
		Code:      "test raw code",
		Workspace: "test workspace",
	}, val)
	reader.AssertExpectations(t)
}
