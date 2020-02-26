package command_test

import (
	"fmt"
	"testing"

	pipeline2 "github.com/arpb2/C-3PO/pkg/domain/architecture/pipeline"
	controller3 "github.com/arpb2/C-3PO/pkg/domain/level/controller"

	"github.com/arpb2/C-3PO/pkg/presentation/level/command"

	"github.com/arpb2/C-3PO/pkg/domain/architecture/http"
	http2 "github.com/arpb2/C-3PO/test/mock/http"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestFetchLevelIdCommand_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	cmd := command.CreateFetchLevelIdCommand()

	name := cmd.Name()

	assert.Equal(t, "fetch_level_id_command", name)
}

func TestFetchLevelIdCommand_GivenOneAndAContextWithoutAReader_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	cmd := command.CreateFetchLevelIdCommand()

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestFetchLevelIdCommand_GivenOneAndAReaderWithoutLevelIdParameter_WhenRunning_Then400(t *testing.T) {
	reader := new(http2.MockReader)
	reader.On("GetParameter", controller3.ParamLevelId).Return("", false)
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline2.TagHttpReader, reader)
	cmd := command.CreateFetchLevelIdCommand()

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateBadRequestError(fmt.Sprintf("'%s' empty", controller3.ParamLevelId)), err)
	reader.AssertExpectations(t)
}

func TestFetchLevelIdCommand_GivenOneAndAReaderWithMalformedLevelId_WhenRunning_Then400(t *testing.T) {
	reader := new(http2.MockReader)
	reader.On("GetParameter", controller3.ParamLevelId).Return("-1", true)
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline2.TagHttpReader, reader)
	cmd := command.CreateFetchLevelIdCommand()

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateBadRequestError(fmt.Sprintf("'%s' malformed, expecting a positive number", controller3.ParamLevelId)), err)
	reader.AssertExpectations(t)
}

func TestFetchLevelIdCommand_GivenOne_WhenRunning_ThenRawLevelIsAddedToContext(t *testing.T) {
	reader := new(http2.MockReader)
	reader.On("GetParameter", controller3.ParamLevelId).Return("1000", true)
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline2.TagHttpReader, reader)
	cmd := command.CreateFetchLevelIdCommand()

	err := cmd.Run(ctx)
	val, exists := ctx.Get(command.TagLevelId)

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, uint(1000), val)
	reader.AssertExpectations(t)
}
