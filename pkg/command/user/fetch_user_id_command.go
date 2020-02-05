package user

import (
	"strconv"

	"github.com/arpb2/C-3PO/pkg/command"
	"github.com/saantiaguilera/go-pipeline"

	"github.com/arpb2/C-3PO/api/http"
)

type fetchUserIdCommand struct{}

func (c *fetchUserIdCommand) Name() string {
	return "fetch_user_id_command"
}

func (c *fetchUserIdCommand) Run(ctx pipeline.Context) error {
	httpReader, exists := ctx.Get(command.TagHttpReader)

	if !exists {
		return http.CreateInternalError()
	}

	userId := httpReader.(http.Reader).GetParameter("user_id")

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
