package session_command

import (
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/controller/session/session_validation"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/model"
)

type ValidateParametersCommand struct {
	Context      *http_wrapper.Context
	InputStream  chan *model.AuthenticatedUser
	OutputStream chan *model.AuthenticatedUser
	Validations  []session_validation.Validation
}

func (c *ValidateParametersCommand) Name() string {
	return "validate_session_parameters_command"
}

func (c *ValidateParametersCommand) Run() error {
	defer close(c.OutputStream)
	user, openChan := <-c.InputStream

	if !openChan {
		return nil
	}

	for _, requirement := range c.Validations {
		if err := requirement(user); err != nil {
			return controller.HaltExternalError(c.Context, http_wrapper.CreateBadRequestError(err.Error()))
		}
	}

	c.OutputStream <- user
	return nil
}

func (c *ValidateParametersCommand) Fallback(err error) error {
	return err
}

func CreateValidateParametersCommand(ctx *http_wrapper.Context,
									 userInput chan *model.AuthenticatedUser,
									 validations []session_validation.Validation) *ValidateParametersCommand {
	return &ValidateParametersCommand{
		Context:      ctx,
		InputStream:  userInput,
		OutputStream: make(chan *model.AuthenticatedUser, 1),
		Validations:  validations,
	}
}
