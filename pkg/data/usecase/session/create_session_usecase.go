package session

import (
	session2 "github.com/arpb2/C-3PO/pkg/data/repository/session"
	"github.com/arpb2/C-3PO/pkg/data/usecase/user"
	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/domain/model/session"
	"github.com/saantiaguilera/go-pipeline"
)

type createSessionUseCase struct {
	tokenHandler session2.TokenRepository
}

func (c *createSessionUseCase) Name() string {
	return "create_session_usecase"
}

func (c *createSessionUseCase) Run(ctx pipeline.Context) error {
	userId, exists := ctx.GetUInt(user.TagUserId)

	if !exists {
		return http.CreateInternalError()
	}

	token, err := c.tokenHandler.Create(&session2.Token{
		UserId: userId,
	})

	if err != nil {
		return err
	}

	ctx.Set(TagSession, session.Session{
		UserId: userId,
		Token:  token,
	})
	return nil
}

func CreateCreateSessionUseCase(tokenHandler session2.TokenRepository) pipeline.Step {
	return &createSessionUseCase{
		tokenHandler: tokenHandler,
	}
}
