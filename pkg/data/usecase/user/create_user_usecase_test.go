package user_test

import (
	"errors"
	"testing"

	"github.com/arpb2/C-3PO/pkg/data/usecase/user"
	user2 "github.com/arpb2/C-3PO/pkg/domain/model/user"

	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/test/mock/repository"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserUseCase_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	useCase := user.CreateCreateUserUseCase(nil)

	name := useCase.Name()

	assert.Equal(t, "create_user_usecase", name)
}

func TestCreateUserUseCase_GivenOneAndAContextWithoutAuthenticatedUser_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	useCase := user.CreateCreateUserUseCase(nil)

	err := useCase.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestCreateUserUseCase_GivenOneAndAFailingRepository_WhenRunning_ThenRepositoryError(t *testing.T) {
	expectedVal := user2.AuthenticatedUser{}
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagAuthenticatedUser, expectedVal)
	expectedErr := errors.New("some error")
	s := new(repository.MockUserRepository)
	s.On("CreateUser", expectedVal).Return(expectedVal.User, expectedErr)
	useCase := user.CreateCreateUserUseCase(s)

	err := useCase.Run(ctx)

	assert.Equal(t, expectedErr, err)
	s.AssertExpectations(t)
}

func TestCreateUserUseCase_GivenOne_WhenRunning_ThenContextHasUserAndReturnsNoError(t *testing.T) {
	expectedVal := user2.AuthenticatedUser{}
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagAuthenticatedUser, expectedVal)
	s := new(repository.MockUserRepository)
	s.On("CreateUser", expectedVal).Return(expectedVal.User, nil)
	useCase := user.CreateCreateUserUseCase(s)

	err := useCase.Run(ctx)
	val, exists := ctx.Get(user.TagUser)

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, expectedVal.User, val)
	s.AssertExpectations(t)
}
