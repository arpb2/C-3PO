package user_command

import (
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/controller/user/user_validation"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/model"
)

type ValidateParametersCommand struct {
	Context      *http_wrapper.Context
	Stream       chan *model.AuthenticatedUser
	Validations  []user_validation.Validation
}

func (c *ValidateParametersCommand) Name() string {
	return "validate_user_parameters_command"
}

func (c *ValidateParametersCommand) Run() error {
	user := <-c.Stream
	for _, requirement := range c.Validations {
		if err := requirement(user); err != nil {
			return http_wrapper.CreateBadRequestError(err.Error())
		}
	}

	c.Stream <- user
	return nil
}

func (c *ValidateParametersCommand) Fallback(err error) error {
	return controller.HaltError(c.Context, err)
}

func CreateValidateParametersCommand(ctx *http_wrapper.Context,
									 userInput chan *model.AuthenticatedUser,
									 validations []user_validation.Validation) *ValidateParametersCommand {
	return &ValidateParametersCommand{
		Context:      ctx,
		Stream:       userInput,
		Validations:  validations,
	}
}
