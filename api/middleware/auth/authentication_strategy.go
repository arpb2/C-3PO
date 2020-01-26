package auth_middleware

import "github.com/arpb2/C-3PO/api/auth"

type AuthenticationStrategy interface {
	Authenticate(token *auth.Token, userId string) (authorized bool, err error)
}
