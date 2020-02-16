package code_test

import (
	"fmt"
	"testing"

	"github.com/arpb2/C-3PO/pkg/controller"

	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/pkg/command/code"
	"github.com/arpb2/C-3PO/pkg/pipeline"
	http2 "github.com/arpb2/C-3PO/test/mock/http"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestFetchCodeIdCommand_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	cmd := code.CreateFetchCodeIdCommand()

	name := cmd.Name()

	assert.Equal(t, fmt.Sprintf("fetch_%s_command", controller.ParamCodeId), name)
}

func TestFetchCodeIdCommand_GivenOneAndAContextWithoutAReader_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	cmd := code.CreateFetchCodeIdCommand()

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestFetchCodeIdCommand_GivenOneAndAReaderWithoutCodeIdParameter_WhenRunning_Then400(t *testing.T) {
	reader := new(http2.MockReader)
	reader.On("GetParameter", controller.ParamCodeId).Return("", false)
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline.TagHttpReader, reader)
	cmd := code.CreateFetchCodeIdCommand()

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateBadRequestError(fmt.Sprintf("'%s' empty", controller.ParamCodeId)), err)
	reader.AssertExpectations(t)
}

func TestFetchCodeIdCommand_GivenOneAndAReaderWithMalformedCodeId_WhenRunning_Then400(t *testing.T) {
	reader := new(http2.MockReader)
	reader.On("GetParameter", controller.ParamCodeId).Return("-1", true)
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline.TagHttpReader, reader)
	cmd := code.CreateFetchCodeIdCommand()

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateBadRequestError(fmt.Sprintf("'%s' malformed, expecting a positive number", controller.ParamCodeId)), err)
	reader.AssertExpectations(t)
}

func TestFetchCodeIdCommand_GivenOne_WhenRunning_ThenRawCodeIsAddedToContext(t *testing.T) {
	reader := new(http2.MockReader)
	reader.On("GetParameter", controller.ParamCodeId).Return("1000", true)
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline.TagHttpReader, reader)
	cmd := code.CreateFetchCodeIdCommand()

	err := cmd.Run(ctx)
	val, exists := ctx.Get(code.TagCodeId)

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, uint(1000), val)
	reader.AssertExpectations(t)
}
