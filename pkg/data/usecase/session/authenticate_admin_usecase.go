package session

import (
	"bytes"

	"github.com/arpb2/C-3PO/pkg/domain/http"
)

func CreateAdminAuthenticationUseCase(
	secret []byte,
) func(string) error {
	return func(tkn string) error {
		if len(tkn) == 0 || bytes.Compare([]byte(tkn), secret) != 0 {
			return http.CreateUnauthorizedError()
		}
		return nil
	}
}
