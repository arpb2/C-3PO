package middleware

import (
	"github.com/arpb2/C-3PO/pkg/domain/session/repository"
)

type AuthenticationStrategy interface {
	Authenticate(token *repository.Token, userId string) (authorized bool, err error)
}
