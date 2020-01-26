package user_command

import (
	"github.com/arpb2/C-3PO/api/http_wrapper"
	"github.com/arpb2/C-3PO/api/model"
	"github.com/arpb2/C-3PO/internal/validation/user"
)

type validateParametersCommand struct {
	context     *http_wrapper.Context
	inputStream chan *model.AuthenticatedUser
	validations []user_validation.Validation

	OutputStream chan *model.AuthenticatedUser
}

func (c *validateParametersCommand) Name() string {
	return "validate_user_parameters_command"
}

func (c *validateParametersCommand) Run() error {
	defer close(c.OutputStream)
	user := <-c.inputStream

	for _, requirement := range c.validations {
		if err := requirement(user); err != nil {
			return http_wrapper.CreateBadRequestError(err.Error())
		}
	}

	c.OutputStream <- user
	return nil
}

func (c *validateParametersCommand) Fallback(err error) error {
	return err
}

func CreateValidateParametersCommand(ctx *http_wrapper.Context,
	userInput chan *model.AuthenticatedUser,
	validations []user_validation.Validation) *validateParametersCommand {
	return &validateParametersCommand{
		context:      ctx,
		inputStream:  userInput,
		OutputStream: make(chan *model.AuthenticatedUser, 1),
		validations:  validations,
	}
}
