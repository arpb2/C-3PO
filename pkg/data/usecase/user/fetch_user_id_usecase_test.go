package user_test

import (
	"testing"

	"github.com/arpb2/C-3PO/pkg/data/usecase/user"

	pipeline2 "github.com/arpb2/C-3PO/pkg/domain/pipeline"

	"github.com/arpb2/C-3PO/pkg/domain/http"
	http2 "github.com/arpb2/C-3PO/test/mock/http"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestFetchUserIdUseCase_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	useCase := user.CreateFetchUserIdUseCase("user_id")

	name := useCase.Name()

	assert.Equal(t, "fetch_user_id_usecase", name)
}

func TestFetchUserIdUseCase_GivenOneAndAContextWithoutAReader_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	useCase := user.CreateFetchUserIdUseCase("user_id")

	err := useCase.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestFetchUserIdUseCase_GivenOneAndAReaderWithoutUserIdParameter_WhenRunning_Then400(t *testing.T) {
	reader := new(http2.MockReader)
	reader.On("GetParameter", "user_id").Return("", false)
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline2.TagHttpReader, reader)
	useCase := user.CreateFetchUserIdUseCase("user_id")

	err := useCase.Run(ctx)

	assert.Equal(t, http.CreateBadRequestError("'user_id' empty"), err)
	reader.AssertExpectations(t)
}

func TestFetchUserIdUseCase_GivenOneAndAReaderWithMalformedUserId_WhenRunning_Then400(t *testing.T) {
	reader := new(http2.MockReader)
	reader.On("GetParameter", "user_id").Return("-1", true)
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline2.TagHttpReader, reader)
	useCase := user.CreateFetchUserIdUseCase("user_id")

	err := useCase.Run(ctx)

	assert.Equal(t, http.CreateBadRequestError("'user_id' malformed, expecting a positive number"), err)
	reader.AssertExpectations(t)
}

func TestFetchUserIdUseCase_GivenOne_WhenRunning_ThenRawUserIsAddedToContext(t *testing.T) {
	reader := new(http2.MockReader)
	reader.On("GetParameter", "user_id").Return("1000", true)
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline2.TagHttpReader, reader)
	useCase := user.CreateFetchUserIdUseCase("user_id")

	err := useCase.Run(ctx)
	val, exists := ctx.Get(user.TagUserId)

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, uint(1000), val)
	reader.AssertExpectations(t)
}
