package user

import (
	"github.com/arpb2/C-3PO/pkg/data/usecase/user/validation"
	"github.com/arpb2/C-3PO/pkg/domain/http"
	httppipeline "github.com/arpb2/C-3PO/pkg/domain/pipeline"
	"github.com/saantiaguilera/go-pipeline"
)

type validateParametersUseCase struct {
	validations []validation.Validation
}

func (c *validateParametersUseCase) Name() string {
	return "validate_user_parameters_usecase"
}

func (c *validateParametersUseCase) Run(ctx pipeline.Context) error {
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

func CreateValidateUserParametersUseCase(validations []validation.Validation) pipeline.Step {
	return &validateParametersUseCase{
		validations: validations,
	}
}
