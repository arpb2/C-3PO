package token_task

import (
	"github.com/arpb2/C-3PO/src/api/auth"
)

type CreateTokenTask func(userId uint, handler auth.TokenHandler) (token *string, err *auth.TokenError)

func CreateTokenTaskImpl(userId uint, handler auth.TokenHandler) (token *string, err *auth.TokenError) {
	token, tokenErr := handler.Create(auth.Token{
		UserId: userId,
	})

	if tokenErr != nil {
		return nil, tokenErr
	}

	return token, nil
}
