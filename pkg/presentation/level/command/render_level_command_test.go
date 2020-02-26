package command_test

import (
	"testing"

	pipeline2 "github.com/arpb2/C-3PO/pkg/domain/infrastructure/pipeline"
	model2 "github.com/arpb2/C-3PO/pkg/domain/level/model"

	"github.com/arpb2/C-3PO/pkg/presentation/level/command"

	"github.com/arpb2/C-3PO/pkg/domain/infrastructure/http"
	http2 "github.com/arpb2/C-3PO/test/mock/http"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestRenderLevelCommand_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	cmd := command.CreateRenderLevelCommand()

	name := cmd.Name()

	assert.Equal(t, "render_level_command", name)
}

func TestRenderLevelCommand_GivenOneAndAContextWithoutWriter_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(command.TagLevel, model2.Level{})
	cmd := command.CreateRenderLevelCommand()

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestRenderLevelCommand_GivenOneAndAContextWithoutLevel_WhenRunning_Then500(t *testing.T) {
	writer := new(http2.MockWriter)
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline2.TagHttpWriter, writer)
	cmd := command.CreateRenderLevelCommand()

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestRenderLevelCommand_GivenOne_WhenRunning_ThenRendersLevel(t *testing.T) {
	expectedVal := model2.Level{}
	writer := new(http2.MockWriter)
	writer.On("WriteJson", 200, expectedVal)
	ctx := gopipeline.CreateContext()
	ctx.Set(command.TagLevel, model2.Level{})
	ctx.Set(pipeline2.TagHttpWriter, writer)
	cmd := command.CreateRenderLevelCommand()

	err := cmd.Run(ctx)

	assert.Nil(t, err)
	writer.AssertExpectations(t)
}
