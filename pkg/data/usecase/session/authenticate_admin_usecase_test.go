package session_test

import (
	"testing"

	"github.com/arpb2/C-3PO/pkg/data/usecase/session"
	"github.com/stretchr/testify/assert"

	"github.com/arpb2/C-3PO/pkg/domain/http"
)

func TestAuthMiddleware_GivenOneWithoutAToken_WhenAuthorizing_ThenItAbortsWithUnauthorized(t *testing.T) {
	m := session.CreateAdminAuthenticationUseCase([]byte("test"))

	err := m("")

	assert.Equal(t, http.CreateUnauthorizedError(), err)
}

func TestAuthMiddleware_GivenOneWithADifferentToken_WhenAuthorizing_ThenItAbortsWithUnauthorized(t *testing.T) {
	m := session.CreateAdminAuthenticationUseCase([]byte("test"))

	err := m("not_test")

	assert.Equal(t, http.CreateUnauthorizedError(), err)
}

func TestAuthMiddleware_GivenOneWithTheSameSecretAsToken_WhenAuthorizing_ThenItContinuesTransaction(t *testing.T) {
	m := session.CreateAdminAuthenticationUseCase([]byte("test"))

	err := m("test")

	assert.Nil(t, err)
}
