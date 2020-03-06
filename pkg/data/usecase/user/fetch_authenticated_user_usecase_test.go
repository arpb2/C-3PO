package user_test

import (
	"errors"
	"testing"

	"github.com/arpb2/C-3PO/pkg/data/usecase/user"
	user2 "github.com/arpb2/C-3PO/pkg/domain/model/user"

	"github.com/arpb2/C-3PO/pkg/domain/http"
	pipeline2 "github.com/arpb2/C-3PO/pkg/domain/pipeline"
	http2 "github.com/arpb2/C-3PO/test/mock/http"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetchAuthenticatedUserUseCase_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	useCase := user.CreateFetchAuthenticatedUserUseCase()

	name := useCase.Name()

	assert.Equal(t, "fetch_authenticated_user_usecase", name)
}

func TestFetchAuthenticatedUserUseCase_GivenOneAndAContextWithoutAReader_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	useCase := user.CreateFetchAuthenticatedUserUseCase()

	err := useCase.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestFetchAuthenticatedUserUseCase_GivenOneAndAReaderWithoutBody_WhenRunning_Then400(t *testing.T) {
	reader := new(http2.MockReader)
	reader.On("ReadBody", mock.Anything).Return(errors.New("some error"))
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline2.TagHttpReader, reader)
	useCase := user.CreateFetchAuthenticatedUserUseCase()

	err := useCase.Run(ctx)

	assert.Equal(t, http.CreateBadRequestError("malformed body"), err)
	reader.AssertExpectations(t)
}

func TestFetchAuthenticatedUserUseCase_GivenOne_WhenRunning_ThenAuthenticatedUserIsAddedToContext(t *testing.T) {
	reader := new(http2.MockReader)
	reader.On("ReadBody", mock.Anything).Return(nil)
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline2.TagHttpReader, reader)
	useCase := user.CreateFetchAuthenticatedUserUseCase()

	err := useCase.Run(ctx)
	val, exists := ctx.Get(user.TagAuthenticatedUser)

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, user2.AuthenticatedUser{}, val)
	reader.AssertExpectations(t)
}
