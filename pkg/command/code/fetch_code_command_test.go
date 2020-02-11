package code_test

import (
	"testing"

	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/pkg/command/code"
	"github.com/arpb2/C-3PO/pkg/pipeline"
	http2 "github.com/arpb2/C-3PO/test/mock/http"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestFetchCodeCommand_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	cmd := code.CreateFetchCodeCommand()

	name := cmd.Name()

	assert.Equal(t, "fetch_code_command", name)
}

func TestFetchCodeCommand_GivenOneAndAContextWithoutAReader_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	cmd := code.CreateFetchCodeCommand()

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestFetchCodeCommand_GivenOneAndAReaderWithoutCodePart_WhenRunning_Then400(t *testing.T) {
	reader := new(http2.MockReader)
	reader.On("GetFormData", "code").Return("", false)
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline.TagHttpReader, reader)
	cmd := code.CreateFetchCodeCommand()

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateBadRequestError("'code' part not found"), err)
	reader.AssertExpectations(t)
}

func TestFetchCodeCommand_GivenOne_WhenRunning_ThenRawCodeIsAddedToContext(t *testing.T) {
	reader := new(http2.MockReader)
	reader.On("GetFormData", "code").Return("test raw code", true)
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline.TagHttpReader, reader)
	cmd := code.CreateFetchCodeCommand()

	err := cmd.Run(ctx)
	val, exists := ctx.Get(code.TagCodeRaw)

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, "test raw code", val)
	reader.AssertExpectations(t)
}
