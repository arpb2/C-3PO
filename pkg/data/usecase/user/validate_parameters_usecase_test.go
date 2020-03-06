package user_test

import (
	"errors"
	"testing"

	"github.com/arpb2/C-3PO/pkg/data/usecase/user"
	"github.com/arpb2/C-3PO/pkg/data/usecase/user/validation"
	user2 "github.com/arpb2/C-3PO/pkg/domain/model/user"

	"github.com/arpb2/C-3PO/pkg/domain/http"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestValidateParametersUseCase_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	useCase := user.CreateValidateUserParametersUseCase(nil)

	name := useCase.Name()

	assert.Equal(t, "validate_user_parameters_usecase", name)
}

func TestValidateParametersUseCase_GivenOneAndAContextWithoutAuthenticatedUser_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	useCase := user.CreateValidateUserParametersUseCase(nil)

	err := useCase.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestValidateParametersUseCase_GivenOneWithAValidatorThatDoesntMetRequirements_WhenRunning_Then400(t *testing.T) {
	expectedMessage := "something wrong"
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagAuthenticatedUser, user2.AuthenticatedUser{})
	useCase := user.CreateValidateUserParametersUseCase([]validation.Validation{
		func(user *user2.AuthenticatedUser) error {
			return nil
		},
		func(user *user2.AuthenticatedUser) error {
			return errors.New(expectedMessage)
		},
	})

	err := useCase.Run(ctx)

	assert.Equal(t, http.CreateBadRequestError(expectedMessage), err)
}

func TestValidateParametersUseCase_GivenOneWithOkValidators_WhenRunning_ThenNoErrors(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagAuthenticatedUser, user2.AuthenticatedUser{})
	useCase := user.CreateValidateUserParametersUseCase([]validation.Validation{
		func(user *user2.AuthenticatedUser) error {
			return nil
		},
		func(user *user2.AuthenticatedUser) error {
			return nil
		},
	})

	err := useCase.Run(ctx)

	assert.Nil(t, err)
}
