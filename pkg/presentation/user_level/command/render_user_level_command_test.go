package command_test

import (
	"testing"

	pipeline2 "github.com/arpb2/C-3PO/pkg/domain/infrastructure/pipeline"
	model2 "github.com/arpb2/C-3PO/pkg/domain/user_level/model"

	"github.com/arpb2/C-3PO/pkg/presentation/user_level/command"

	"github.com/arpb2/C-3PO/pkg/domain/infrastructure/http"
	http2 "github.com/arpb2/C-3PO/test/mock/http"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestRenderUserLevelCommand_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	cmd := command.CreateRenderUserLevelCommand()

	name := cmd.Name()

	assert.Equal(t, "render_user_level_command", name)
}

func TestRenderUserLevelCommand_GivenOneAndAContextWithoutWriter_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(command.TagUserLevel, model2.UserLevel{})
	cmd := command.CreateRenderUserLevelCommand()

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestRenderUserLevelCommand_GivenOneAndAContextWithoutCode_WhenRunning_Then500(t *testing.T) {
	writer := new(http2.MockWriter)
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline2.TagHttpWriter, writer)
	cmd := command.CreateRenderUserLevelCommand()

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestRenderUserLevelCommand_GivenOne_WhenRunning_ThenRendersCode(t *testing.T) {
	expectedVal := model2.UserLevel{}
	writer := new(http2.MockWriter)
	writer.On("WriteJson", 200, expectedVal)
	ctx := gopipeline.CreateContext()
	ctx.Set(command.TagUserLevel, expectedVal)
	ctx.Set(pipeline2.TagHttpWriter, writer)
	cmd := command.CreateRenderUserLevelCommand()

	err := cmd.Run(ctx)

	assert.Nil(t, err)
	writer.AssertExpectations(t)
}
