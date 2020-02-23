package command_test

import (
	"errors"
	"testing"

	"github.com/arpb2/C-3PO/pkg/presentation/user/command"

	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/domain/model"
	"github.com/arpb2/C-3PO/pkg/infra/pipeline"
	http2 "github.com/arpb2/C-3PO/test/mock/http"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetchAuthenticatedUserCommand_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	cmd := command.CreateFetchAuthenticatedUserCommand()

	name := cmd.Name()

	assert.Equal(t, "fetch_authenticated_user_command", name)
}

func TestFetchAuthenticatedUserCommand_GivenOneAndAContextWithoutAReader_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	cmd := command.CreateFetchAuthenticatedUserCommand()

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestFetchAuthenticatedUserCommand_GivenOneAndAReaderWithoutBody_WhenRunning_Then400(t *testing.T) {
	reader := new(http2.MockReader)
	reader.On("ReadBody", mock.Anything).Return(errors.New("some error"))
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline.TagHttpReader, reader)
	cmd := command.CreateFetchAuthenticatedUserCommand()

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateBadRequestError("malformed body"), err)
	reader.AssertExpectations(t)
}

func TestFetchAuthenticatedUserCommand_GivenOne_WhenRunning_ThenAuthenticatedUserIsAddedToContext(t *testing.T) {
	reader := new(http2.MockReader)
	reader.On("ReadBody", mock.Anything).Return(nil)
	ctx := gopipeline.CreateContext()
	ctx.Set(pipeline.TagHttpReader, reader)
	cmd := command.CreateFetchAuthenticatedUserCommand()

	err := cmd.Run(ctx)
	val, exists := ctx.Get(command.TagAuthenticatedUser)

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, model.AuthenticatedUser{}, val)
	reader.AssertExpectations(t)
}
