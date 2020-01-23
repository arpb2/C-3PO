package user_command

import (
	"github.com/arpb2/C-3PO/src/api/command"
	"github.com/arpb2/C-3PO/src/api/controller/user/user_validation"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/model"
)

type validateParametersCommand struct {
	context      *http_wrapper.Context
	inputStream  chan *model.AuthenticatedUser
	validations  []user_validation.Validation

	OutputStream chan *model.AuthenticatedUser

	user         *model.AuthenticatedUser
}

func (c *validateParametersCommand) Name() string {
	return "validate_user_parameters_command"
}

func (c *validateParametersCommand) Prepare() bool {
	c.user = <-c.inputStream

	if c.user == nil {
		close(c.OutputStream)
	}

	return c.user != nil
}

func (c *validateParametersCommand) Run() error {
	defer close(c.OutputStream)

	for _, requirement := range c.validations {
		if err := requirement(c.user); err != nil {
			return command.HaltExternalError(c.context, http_wrapper.CreateBadRequestError(err.Error()))
		}
	}

	c.OutputStream <- c.user
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
