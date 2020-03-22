package user_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/arpb2/C-3PO/pkg/data/usecase/user"
	user2 "github.com/arpb2/C-3PO/pkg/domain/model/user"

	"github.com/arpb2/C-3PO/pkg/domain/http"
	pipeline2 "github.com/arpb2/C-3PO/pkg/domain/pipeline"
	http2 "github.com/arpb2/C-3PO/test/mock/http"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestFetchCodeUseCase_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	useCase := user.CreateFetchCodeUseCase()

	name := useCase.Name()

	assert.Equal(t, "fetch_user_level_data_usecase", name)
}

func TestFetchCodeUseCase_GivenOneAndAContextWithoutAReader_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	useCase := user.CreateFetchCodeUseCase()

	err := useCase.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestFetchCodeUseCase_GivenOneAndAReaderThatFails_WhenRunning_Then400(t *testing.T) {
	reader := new(http2.MockReader)
	reader.On("ReadBody", mock.Anything).Return(errors.New("some error"))
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline2.TagHttpReader, reader)
	useCase := user.CreateFetchCodeUseCase()

	err := useCase.Run(ctx)

	assert.Equal(t, http.CreateBadRequestError("error reading user level json"), err)
	reader.AssertExpectations(t)
}

func TestFetchCodeUseCase_GivenOne_WhenRunning_ThenRawCodeIsAddedToContext(t *testing.T) {
	reader := new(http2.MockReader)
	reader.On("ReadBody", mock.Anything).Run(func(args mock.Arguments) {
		ul := args.Get(0).(*user2.LevelData)
		ul.Code = "test raw code"
		ul.Workspace = "test workspace"
	}).Return(nil)
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline2.TagHttpReader, reader)
	useCase := user.CreateFetchCodeUseCase()

	err := useCase.Run(ctx)
	val, exists := ctx.Get(user.TagUserLevelData)

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, user2.LevelData{
		Code:      "test raw code",
		Workspace: "test workspace",
	}, val)
	reader.AssertExpectations(t)
}
