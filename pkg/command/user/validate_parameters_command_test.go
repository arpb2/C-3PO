package user_test

import (
	"errors"
	"testing"

	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/api/model"
	"github.com/arpb2/C-3PO/pkg/command/user"
	uservalidation "github.com/arpb2/C-3PO/pkg/validation/user"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestValidateParametersCommand_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	cmd := user.CreateValidateUserParametersCommand(nil)

	name := cmd.Name()

	assert.Equal(t, "validate_user_parameters_command", name)
}

func TestValidateParametersCommand_GivenOneAndAContextWithoutAuthenticatedUser_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	cmd := user.CreateValidateUserParametersCommand(nil)

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestValidateParametersCommand_GivenOneWithAValidatorThatDoesntMetRequirements_WhenRunning_Then400(t *testing.T) {
	expectedMessage := "something wrong"
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagAuthenticatedUser, model.AuthenticatedUser{})
	cmd := user.CreateValidateUserParametersCommand([]uservalidation.Validation{
		func(user *model.AuthenticatedUser) error {
			return nil
		},
		func(user *model.AuthenticatedUser) error {
			return errors.New(expectedMessage)
		},
	})

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateBadRequestError(expectedMessage), err)
}

func TestValidateParametersCommand_GivenOneWithOkValidators_WhenRunning_ThenNoErrors(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(user.TagAuthenticatedUser, model.AuthenticatedUser{})
	cmd := user.CreateValidateUserParametersCommand([]uservalidation.Validation{
		func(user *model.AuthenticatedUser) error {
			return nil
		},
		func(user *model.AuthenticatedUser) error {
			return nil
		},
	})

	err := cmd.Run(ctx)

	assert.Nil(t, err)
}
