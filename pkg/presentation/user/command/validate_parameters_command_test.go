package command_test

import (
	"errors"
	"testing"

	model2 "github.com/arpb2/C-3PO/pkg/domain/user/model"

	"github.com/arpb2/C-3PO/pkg/presentation/user/command"
	"github.com/arpb2/C-3PO/pkg/presentation/user/validation"

	"github.com/arpb2/C-3PO/pkg/domain/architecture/http"
	gopipeline "github.com/saantiaguilera/go-pipeline"
	"github.com/stretchr/testify/assert"
)

func TestValidateParametersCommand_GivenOne_WhenCallingName_ThenItsTheExpected(t *testing.T) {
	cmd := command.CreateValidateUserParametersCommand(nil)

	name := cmd.Name()

	assert.Equal(t, "validate_user_parameters_command", name)
}

func TestValidateParametersCommand_GivenOneAndAContextWithoutAuthenticatedUser_WhenRunning_Then500(t *testing.T) {
	ctx := gopipeline.CreateContext()
	cmd := command.CreateValidateUserParametersCommand(nil)

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateInternalError(), err)
}

func TestValidateParametersCommand_GivenOneWithAValidatorThatDoesntMetRequirements_WhenRunning_Then400(t *testing.T) {
	expectedMessage := "something wrong"
	ctx := gopipeline.CreateContext()
	ctx.Set(command.TagAuthenticatedUser, model2.AuthenticatedUser{})
	cmd := command.CreateValidateUserParametersCommand([]validation.Validation{
		func(user *model2.AuthenticatedUser) error {
			return nil
		},
		func(user *model2.AuthenticatedUser) error {
			return errors.New(expectedMessage)
		},
	})

	err := cmd.Run(ctx)

	assert.Equal(t, http.CreateBadRequestError(expectedMessage), err)
}

func TestValidateParametersCommand_GivenOneWithOkValidators_WhenRunning_ThenNoErrors(t *testing.T) {
	ctx := gopipeline.CreateContext()
	ctx.Set(command.TagAuthenticatedUser, model2.AuthenticatedUser{})
	cmd := command.CreateValidateUserParametersCommand([]validation.Validation{
		func(user *model2.AuthenticatedUser) error {
			return nil
		},
		func(user *model2.AuthenticatedUser) error {
			return nil
		},
	})

	err := cmd.Run(ctx)

	assert.Nil(t, err)
}
