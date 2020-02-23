package command_test

import (
	"fmt"
	controller2 "github.com/arpb2/C-3PO/pkg/domain/controller"
	"github.com/arpb2/C-3PO/pkg/presentation/user_level/command"
	"testing"

	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/infra/pipeline"
	http2 "github.com/arpb2/C-3PO/test/mock/http"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestFetchLevelIdCommand_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	cmd := command.CreateFetchLevelIdCommand()

	name := cmd.Name()

	assert.Equal(t, fmt.Sprintf("fetch_%s_command", controller2.ParamLevelId), name)
}

func TestFetchLevelIdCommand_GivenOneAndAContextWithoutAReader_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	cmd := command.CreateFetchLevelIdCommand()

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestFetchLevelIdCommand_GivenOneAndAReaderWithoutLevelIdParameter_WhenRunning_Then400(t *testing.T) {
	reader := new(http2.MockReader)
	reader.On("GetParameter", controller2.ParamLevelId).Return("", false)
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline.TagHttpReader, reader)
	cmd := command.CreateFetchLevelIdCommand()

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateBadRequestError(fmt.Sprintf("'%s' empty", controller2.ParamLevelId)), err)
	reader.AssertExpectations(t)
}

func TestFetchLevelIdCommand_GivenOneAndAReaderWithMalformedLevelId_WhenRunning_Then400(t *testing.T) {
	reader := new(http2.MockReader)
	reader.On("GetParameter", controller2.ParamLevelId).Return("-1", true)
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline.TagHttpReader, reader)
	cmd := command.CreateFetchLevelIdCommand()

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateBadRequestError(fmt.Sprintf("'%s' malformed, expecting a positive number", controller2.ParamLevelId)), err)
	reader.AssertExpectations(t)
}

func TestFetchLevelIdCommand_GivenOne_WhenRunning_ThenRawCodeIsAddedToContext(t *testing.T) {
	reader := new(http2.MockReader)
	reader.On("GetParameter", controller2.ParamLevelId).Return("1000", true)
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline.TagHttpReader, reader)
	cmd := command.CreateFetchLevelIdCommand()

	err := cmd.Run(ctx)
	val, exists := ctx.Get(command.TagLevelId)

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, uint(1000), val)
	reader.AssertExpectations(t)
}
