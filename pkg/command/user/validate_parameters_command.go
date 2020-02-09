package user

import (
	"github.com/arpb2/C-3PO/api/http"
	httppipeline "github.com/arpb2/C-3PO/pkg/pipeline"
	uservalidation "github.com/arpb2/C-3PO/pkg/validation/user"
	"github.com/saantiaguilera/go-pipeline"
)

type validateParametersCommand struct {
	validations []uservalidation.Validation
}

func (c *validateParametersCommand) Name() string {
	return "validate_user_parameters_command"
}

func (c *validateParametersCommand) Run(ctx pipeline.Context) error {
	ctxAware := httppipeline.CreateContextAware(ctx)

	authenticatedUser, err := ctxAware.GetAuthenticatedUser(TagAuthenticatedUser)

	if err != nil {
		return err
	}

	for _, requirement := range c.validations {
		if err := requirement(&authenticatedUser); err != nil {
			return http.CreateBadRequestError(err.Error())
		}
	}

	return nil
}

func (c *validateParametersCommand) Fallback(err error) error {
	return err
}

func CreateValidateParametersCommand(validations []uservalidation.Validation) pipeline.Step {
	return &validateParametersCommand{
		validations: validations,
	}
}
