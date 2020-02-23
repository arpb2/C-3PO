package command

import (
	"github.com/arpb2/C-3PO/pkg/domain/http"
	httppipeline "github.com/arpb2/C-3PO/pkg/infra/pipeline"
	"github.com/arpb2/C-3PO/pkg/presentation/user/validation"
	"github.com/saantiaguilera/go-pipeline"
)

type validateParametersCommand struct {
	validations []validation.Validation
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

func CreateValidateUserParametersCommand(validations []validation.Validation) pipeline.Step {
	return &validateParametersCommand{
		validations: validations,
	}
}
