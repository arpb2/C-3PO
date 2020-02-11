package user_test

import (
	"testing"

	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/pkg/command/user"
	"github.com/arpb2/C-3PO/pkg/pipeline"
	http2 "github.com/arpb2/C-3PO/test/mock/http"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestFetchUserIdCommand_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	cmd := user.CreateFetchUserIdCommand()

	name := cmd.Name()

	assert.Equal(t, "fetch_user_id_command", name)
}

func TestFetchUserIdCommand_GivenOneAndAContextWithoutAReader_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	cmd := user.CreateFetchUserIdCommand()

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestFetchUserIdCommand_GivenOneAndAReaderWithoutUserIdParameter_WhenRunning_Then400(t *testing.T) {
	reader := new(http2.MockReader)
	reader.On("GetParameter", "user_id").Return("", false)
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline.TagHttpReader, reader)
	cmd := user.CreateFetchUserIdCommand()

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateBadRequestError("'user_id' empty"), err)
	reader.AssertExpectations(t)
}

func TestFetchUserIdCommand_GivenOneAndAReaderWithMalformedUserId_WhenRunning_Then400(t *testing.T) {
	reader := new(http2.MockReader)
	reader.On("GetParameter", "user_id").Return("-1", true)
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline.TagHttpReader, reader)
	cmd := user.CreateFetchUserIdCommand()

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateBadRequestError("'user_id' malformed, expecting a positive number"), err)
	reader.AssertExpectations(t)
}

func TestFetchUserIdCommand_GivenOne_WhenRunning_ThenRawUserIsAddedToContext(t *testing.T) {
	reader := new(http2.MockReader)
	reader.On("GetParameter", "user_id").Return("1000", true)
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline.TagHttpReader, reader)
	cmd := user.CreateFetchUserIdCommand()

	err := cmd.Run(ctx)
	val, exists := ctx.Get(user.TagUserId)

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, uint(1000), val)
	reader.AssertExpectations(t)
}
