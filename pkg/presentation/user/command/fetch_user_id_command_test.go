package command_test

import (
	"testing"

	pipeline2 "github.com/arpb2/C-3PO/pkg/domain/infrastructure/pipeline"

	"github.com/arpb2/C-3PO/pkg/presentation/user/command"

	"github.com/arpb2/C-3PO/pkg/domain/infrastructure/http"
	http2 "github.com/arpb2/C-3PO/test/mock/http"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestFetchUserIdCommand_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	cmd := command.CreateFetchUserIdCommand()

	name := cmd.Name()

	assert.Equal(t, "fetch_user_id_command", name)
}

func TestFetchUserIdCommand_GivenOneAndAContextWithoutAReader_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	cmd := command.CreateFetchUserIdCommand()

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestFetchUserIdCommand_GivenOneAndAReaderWithoutUserIdParameter_WhenRunning_Then400(t *testing.T) {
	reader := new(http2.MockReader)
	reader.On("GetParameter", "user_id").Return("", false)
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline2.TagHttpReader, reader)
	cmd := command.CreateFetchUserIdCommand()

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateBadRequestError("'user_id' empty"), err)
	reader.AssertExpectations(t)
}

func TestFetchUserIdCommand_GivenOneAndAReaderWithMalformedUserId_WhenRunning_Then400(t *testing.T) {
	reader := new(http2.MockReader)
	reader.On("GetParameter", "user_id").Return("-1", true)
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline2.TagHttpReader, reader)
	cmd := command.CreateFetchUserIdCommand()

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateBadRequestError("'user_id' malformed, expecting a positive number"), err)
	reader.AssertExpectations(t)
}

func TestFetchUserIdCommand_GivenOne_WhenRunning_ThenRawUserIsAddedToContext(t *testing.T) {
	reader := new(http2.MockReader)
	reader.On("GetParameter", "user_id").Return("1000", true)
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline2.TagHttpReader, reader)
	cmd := command.CreateFetchUserIdCommand()

	err := cmd.Run(ctx)
	val, exists := ctx.Get(command.TagUserId)

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, uint(1000), val)
	reader.AssertExpectations(t)
}
