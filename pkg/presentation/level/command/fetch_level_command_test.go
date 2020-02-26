package command_test

import (
	"errors"
	"testing"

	pipeline2 "github.com/arpb2/C-3PO/pkg/domain/infrastructure/pipeline"
	model2 "github.com/arpb2/C-3PO/pkg/domain/level/model"

	"github.com/arpb2/C-3PO/pkg/presentation/level/command"

	"github.com/arpb2/C-3PO/pkg/domain/infrastructure/http"
	http2 "github.com/arpb2/C-3PO/test/mock/http"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetchLevelCommand_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	cmd := command.CreateFetchLevelCommand()

	name := cmd.Name()

	assert.Equal(t, "fetch_level_command", name)
}

func TestFetchLevelCommand_GivenOneAndAContextWithoutAReader_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	cmd := command.CreateFetchLevelCommand()

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestFetchLevelCommand_GivenOneAndAReaderWithoutBody_WhenRunning_Then400(t *testing.T) {
	reader := new(http2.MockReader)
	reader.On("ReadBody", mock.Anything).Return(errors.New("some error"))
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline2.TagHttpReader, reader)
	cmd := command.CreateFetchLevelCommand()

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateBadRequestError("malformed body"), err)
	reader.AssertExpectations(t)
}

func TestFetchLevelCommand_GivenOne_WhenRunning_ThenLevelIsAddedToContext(t *testing.T) {
	reader := new(http2.MockReader)
	reader.On("ReadBody", mock.Anything).Return(nil)
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline2.TagHttpReader, reader)
	cmd := command.CreateFetchLevelCommand()

	err := cmd.Run(ctx)
	val, exists := ctx.Get(command.TagLevel)

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, model2.Level{}, val)
	reader.AssertExpectations(t)
}
