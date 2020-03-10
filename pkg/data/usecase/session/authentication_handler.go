package session

import (
	"fmt"
	"strconv"

	"github.com/arpb2/C-3PO/pkg/data/repository/session"

	"github.com/arpb2/C-3PO/pkg/domain/http"
)

type AuthenticationStrategy interface {
	Authenticate(token *session.Token, userId string) (authorized bool, err error)
}

func HandleTokenizedAuthentication(authHeader, userId string, tokenHandler session.TokenRepository, strategies ...AuthenticationStrategy) error {
	if len(authHeader) == 0 {
		return http.CreateUnauthorizedError()
	}

	token, err := tokenHandler.Retrieve(authHeader)

	if err != nil {
		return err
	}

	if strconv.FormatUint(uint64(token.UserId), 10) == userId {
		return nil
	}

	for _, strategy := range strategies {
		authenticated, err := strategy.Authenticate(token, userId)

		// If any of our strategies has an error we will instantly fail the authentication process.
		if err != nil {
			fmt.Println(err)
			return http.CreateUnauthorizedError()
		}

		// If at least one of the strategies considers us authenticated, then we can continue.
		if authenticated {
			return nil
		}
	}

	return http.CreateUnauthorizedError()
}
