package middleware

import (
	"github.com/arpb2/C-3PO/pkg/domain/session/token"
)

type AuthenticationStrategy interface {
	Authenticate(token *token.Token, userId string) (authorized bool, err error)
}
