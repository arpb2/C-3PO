package user

import (
	"fmt"
	"strconv"

	httppipeline "github.com/arpb2/C-3PO/pkg/domain/pipeline"

	"github.com/saantiaguilera/go-pipeline"

	"github.com/arpb2/C-3PO/pkg/domain/http"
)

type fetchUserIdUseCase struct {
	UserIdParam string
}

func (c *fetchUserIdUseCase) Name() string {
	return fmt.Sprintf("fetch_%s_usecase", c.UserIdParam)
}

func (c *fetchUserIdUseCase) Run(ctx pipeline.Context) error {
	ctxAware := httppipeline.CreateContextAware(ctx)

	httpReader, err := ctxAware.GetReader()

	if err != nil {
		return err
	}

	userId := httpReader.GetParameter(c.UserIdParam)

	if userId == "" {
		return http.CreateBadRequestError(fmt.Sprintf("'%s' empty", c.UserIdParam))
	}

	userIdUint, err := strconv.ParseUint(userId, 10, 64)

	if err != nil {
		return http.CreateBadRequestError(fmt.Sprintf("'%s' malformed, expecting a positive number", c.UserIdParam))
	}

	ctx.Set(TagUserId, uint(userIdUint))
	return nil
}

func CreateFetchUserIdUseCase(userIdParam string) pipeline.Step {
	return &fetchUserIdUseCase{
		UserIdParam: userIdParam,
	}
}
