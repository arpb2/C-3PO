package user_test

import (
	"errors"
	"testing"

	"github.com/arpb2/C-3PO/pkg/data/usecase/user"

	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/test/mock/repository"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestDeleteUserUseCase_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	useCase := user.CreateDeleteUserUseCase(nil)

	name := useCase.Name()

	assert.Equal(t, "delete_user_usecase", name)
}

func TestDeleteUserUseCase_GivenOneAndAContextWithoutUserId_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	useCase := user.CreateDeleteUserUseCase(nil)

	err := useCase.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestDeleteUserUseCase_GivenOneAndAFailingRepository_WhenRunning_ThenRepositoryError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagUserId, uint(1000))
	expectedErr := errors.New("some error")
	s := new(repository.MockUserRepository)
	s.On("DeleteUser", uint(1000)).Return(expectedErr)
	useCase := user.CreateDeleteUserUseCase(s)

	err := useCase.Run(ctx)

	assert.Equal(t, expectedErr, err)
	s.AssertExpectations(t)
}

func TestDeleteUserUseCase_GivenOne_WhenRunning_ThenReturnsNoError(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagUserId, uint(1000))
	s := new(repository.MockUserRepository)
	s.On("DeleteUser", uint(1000)).Return(nil)
	useCase := user.CreateDeleteUserUseCase(s)

	err := useCase.Run(ctx)

	assert.Nil(t, err)
	s.AssertExpectations(t)
}
