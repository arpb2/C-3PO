package session_command

import (
	"github.com/arpb2/C-3PO/src/api/controller/session/session_validation"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/model"
)

type validateParametersCommand struct {
	context      *http_wrapper.Context
	inputStream  chan *model.AuthenticatedUser
	validations  []session_validation.Validation

	OutputStream chan *model.AuthenticatedUser
}

func (c *validateParametersCommand) Name() string {
	return "validate_session_parameters_command"
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

func CreateValidateParametersCommand(ctx *http_wrapper.Context,
									 userInput chan *model.AuthenticatedUser,
									 validations []session_validation.Validation) *validateParametersCommand {
	return &validateParametersCommand{
		context:      ctx,
		inputStream:  userInput,
		OutputStream: make(chan *model.AuthenticatedUser, 1),
		validations:  validations,
	}
}
