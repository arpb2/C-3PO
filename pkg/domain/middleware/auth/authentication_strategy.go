package auth

import "github.com/arpb2/C-3PO/pkg/domain/auth"

type AuthenticationStrategy interface {
	Authenticate(token *auth.Token, userId string) (authorized bool, err error)
}
