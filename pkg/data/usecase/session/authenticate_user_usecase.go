package session

import (
	"github.com/arpb2/C-3PO/pkg/data/repository/session"
)

func CreateUserAuthenticationUseCase(
	tokenHandler session.TokenRepository,
) func(string, string) error {
	return func(authToken, userId string) error {
		return HandleTokenizedAuthentication(authToken, userId, tokenHandler)
	}
}
