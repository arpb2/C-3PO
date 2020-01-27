package user

import (
	"github.com/arpb2/C-3PO/api/http"
	"github.com/arpb2/C-3PO/api/model"
	uservalidation "github.com/arpb2/C-3PO/pkg/validation/user"
)

type validateParametersCommand struct {
	context     *http.Context
	inputStream chan *model.AuthenticatedUser
	validations []uservalidation.Validation

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
			return http.CreateBadRequestError(err.Error())
		}
	}

	c.OutputStream <- user
	return nil
}

func (c *validateParametersCommand) Fallback(err error) error {
	return err
}

func CreateValidateParametersCommand(ctx *http.Context,
	userInput chan *model.AuthenticatedUser,
	validations []uservalidation.Validation) *validateParametersCommand {
	return &validateParametersCommand{
		context:      ctx,
		inputStream:  userInput,
		OutputStream: make(chan *model.AuthenticatedUser, 1),
		validations:  validations,
	}
}
