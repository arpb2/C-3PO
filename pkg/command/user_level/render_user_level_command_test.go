package user_level_test

import (
	"testing"

	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/api/model"
	"github.com/arpb2/C-3PO/pkg/command/user_level"
	"github.com/arpb2/C-3PO/pkg/pipeline"
	http2 "github.com/arpb2/C-3PO/test/mock/http"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestRenderUserLevelCommand_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	cmd := user_level.CreateRenderUserLevelCommand()

	name := cmd.Name()

	assert.Equal(t, "render_user_level_command", name)
}

func TestRenderUserLevelCommand_GivenOneAndAContextWithoutWriter_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(user_level.TagUserLevel, model.UserLevel{})
	cmd := user_level.CreateRenderUserLevelCommand()

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestRenderUserLevelCommand_GivenOneAndAContextWithoutCode_WhenRunning_Then500(t *testing.T) {
	writer := new(http2.MockWriter)
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline.TagHttpWriter, writer)
	cmd := user_level.CreateRenderUserLevelCommand()

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestRenderUserLevelCommand_GivenOne_WhenRunning_ThenRendersCode(t *testing.T) {
	expectedVal := model.UserLevel{}
	writer := new(http2.MockWriter)
	writer.On("WriteJson", 200, expectedVal)
	ctx := gopipeline.CreateContext()
	ctx.Set(user_level.TagUserLevel, expectedVal)
	ctx.Set(pipeline.TagHttpWriter, writer)
	cmd := user_level.CreateRenderUserLevelCommand()

	err := cmd.Run(ctx)

	assert.Nil(t, err)
	writer.AssertExpectations(t)
}
