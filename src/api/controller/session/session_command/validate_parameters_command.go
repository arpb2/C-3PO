package session_command

import (
	"github.com/arpb2/C-3PO/src/api/controller"
	"github.com/arpb2/C-3PO/src/api/controller/session/session_validation"
	"github.com/arpb2/C-3PO/src/api/http_wrapper"
	"github.com/arpb2/C-3PO/src/api/model"
	"golang.org/x/xerrors"
	"net/http"
)

type ValidateParametersCommand struct {
	Context      *http_wrapper.Context
	Stream       chan *model.AuthenticatedUser
	Validations  []session_validation.Validation
}

func (c *ValidateParametersCommand) Name() string {
	return "validate_parameters_command"
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
	var httpError http_wrapper.HttpError
	if xerrors.As(err, &httpError) {
		if httpError.Code == http.StatusInternalServerError {
			return err
		} else {
			controller.Halt(c.Context, httpError.Code, httpError.Error())
			return nil
		}
	}

	return err
}

func CreateValidateParametersCommand(ctx *http_wrapper.Context,
									 userInput chan *model.AuthenticatedUser,
									 validations []session_validation.Validation) *ValidateParametersCommand {
	return &ValidateParametersCommand{
		Context:      ctx,
		Stream:       userInput,
		Validations:  validations,
	}
}
