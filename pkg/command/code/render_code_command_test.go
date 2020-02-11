package code_test

import (
	"testing"

	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/api/model"
	"github.com/arpb2/C-3PO/pkg/command/code"
	"github.com/arpb2/C-3PO/pkg/pipeline"
	http2 "github.com/arpb2/C-3PO/test/mock/http"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestRenderCodeCommand_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	cmd := code.CreateRenderCodeCommand()

	name := cmd.Name()

	assert.Equal(t, "render_code_command", name)
}

func TestRenderCodeCommand_GivenOneAndAContextWithoutWriter_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(code.TagCode, model.Code{})
	cmd := code.CreateRenderCodeCommand()

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestRenderCodeCommand_GivenOneAndAContextWithoutCode_WhenRunning_Then500(t *testing.T) {
	writer := new(http2.MockWriter)
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline.TagHttpWriter, writer)
	cmd := code.CreateRenderCodeCommand()

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestRenderCodeCommand_GivenOne_WhenRunning_ThenRendersCode(t *testing.T) {
	exptectedVal := model.Code{}
	writer := new(http2.MockWriter)
	writer.On("WriteJson", 200, exptectedVal)
	ctx := gopipeline.CreateContext()
	ctx.Set(code.TagCode, exptectedVal)
	ctx.Set(pipeline.TagHttpWriter, writer)
	cmd := code.CreateRenderCodeCommand()

	err := cmd.Run(ctx)

	assert.Nil(t, err)
	writer.AssertExpectations(t)
}
