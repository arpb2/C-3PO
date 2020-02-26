package command_test

import (
	"fmt"
	"github.com/arpb2/C-3PO/pkg/presentation/level"
	"testing"

	httppipeline "github.com/arpb2/C-3PO/pkg/domain/architecture/pipeline"
	levelcommand "github.com/arpb2/C-3PO/pkg/presentation/level/command"

	"github.com/arpb2/C-3PO/pkg/presentation/user_level/command"

	"github.com/arpb2/C-3PO/pkg/domain/architecture/http"
	http2 "github.com/arpb2/C-3PO/test/mock/http"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestFetchLevelIdCommand_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	cmd := command.CreateFetchLevelIdCommand()

	name := cmd.Name()

	assert.Equal(t, fmt.Sprintf("fetch_%s_command", level.ParamLevelId), name)
}

func TestFetchLevelIdCommand_GivenOneAndAContextWithoutAReader_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	cmd := command.CreateFetchLevelIdCommand()

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestFetchLevelIdCommand_GivenOneAndAReaderWithoutLevelIdParameter_WhenRunning_Then400(t *testing.T) {
	reader := new(http2.MockReader)
	reader.On("GetParameter", level.ParamLevelId).Return("", false)
	ctx := gopipeline.CreateContext()
	ctx.Set(httppipeline.TagHttpReader, reader)
	cmd := command.CreateFetchLevelIdCommand()

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateBadRequestError(fmt.Sprintf("'%s' empty", level.ParamLevelId)), err)
	reader.AssertExpectations(t)
}

func TestFetchLevelIdCommand_GivenOneAndAReaderWithMalformedLevelId_WhenRunning_Then400(t *testing.T) {
	reader := new(http2.MockReader)
	reader.On("GetParameter", level.ParamLevelId).Return("-1", true)
	ctx := gopipeline.CreateContext()
	ctx.Set(httppipeline.TagHttpReader, reader)
	cmd := command.CreateFetchLevelIdCommand()

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateBadRequestError(fmt.Sprintf("'%s' malformed, expecting a positive number", level.ParamLevelId)), err)
	reader.AssertExpectations(t)
}

func TestFetchLevelIdCommand_GivenOne_WhenRunning_ThenRawCodeIsAddedToContext(t *testing.T) {
	reader := new(http2.MockReader)
	reader.On("GetParameter", level.ParamLevelId).Return("1000", true)
	ctx := gopipeline.CreateContext()
	ctx.Set(httppipeline.TagHttpReader, reader)
	cmd := command.CreateFetchLevelIdCommand()

	err := cmd.Run(ctx)
	val, exists := ctx.Get(levelcommand.TagLevelId)

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, uint(1000), val)
	reader.AssertExpectations(t)
}
