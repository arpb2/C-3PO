package command

import (
	"strconv"

	httppipeline "github.com/arpb2/C-3PO/pkg/infra/pipeline"

	"github.com/saantiaguilera/go-pipeline"

	"github.com/arpb2/C-3PO/pkg/domain/http"
)

type fetchUserIdCommand struct{}

func (c *fetchUserIdCommand) Name() string {
	return "fetch_user_id_command"
}

func (c *fetchUserIdCommand) Run(ctx pipeline.Context) error {
	ctxAware := httppipeline.CreateContextAware(ctx)

	httpReader, err := ctxAware.GetReader()

	if err != nil {
		return err
	}

	userId := httpReader.GetParameter("user_id")

	if userId == "" {
		return http.CreateBadRequestError("'user_id' empty")
	}

	userIdUint, err := strconv.ParseUint(userId, 10, 64)

	if err != nil {
		return http.CreateBadRequestError("'user_id' malformed, expecting a positive number")
	}

	ctx.Set(TagUserId, uint(userIdUint))
	return nil
}

func CreateFetchUserIdCommand() pipeline.Step {
	return &fetchUserIdCommand{}
}
