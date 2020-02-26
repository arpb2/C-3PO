package admin_test

import (
	"testing"

	"github.com/arpb2/C-3PO/pkg/presentation/middleware"

	"github.com/arpb2/C-3PO/pkg/domain/infrastructure/http"
	"github.com/arpb2/C-3PO/pkg/presentation/middleware/admin"
	httpmock "github.com/arpb2/C-3PO/test/mock/http"
)

func TestAuthMiddleware_GivenOneWithoutAToken_WhenAuthorizing_ThenItAbortsWithUnauthorized(t *testing.T) {
	reader := new(httpmock.MockReader)
	reader.On("GetHeader", middleware.HeaderAuthorization).Return("").Once()
	middlew := new(httpmock.MockMiddleware)
	middlew.On("AbortTransactionWithError", http.CreateUnauthorizedError()).Once()
	ctx := &http.Context{
		Reader:     reader,
		Middleware: middlew,
	}
	m := admin.CreateMiddleware([]byte("test"))

	m(ctx)

	reader.AssertExpectations(t)
	middlew.AssertExpectations(t)
}

func TestAuthMiddleware_GivenOneWithADifferentToken_WhenAuthorizing_ThenItAbortsWithUnauthorized(t *testing.T) {
	reader := new(httpmock.MockReader)
	reader.On("GetHeader", middleware.HeaderAuthorization).Return("not test").Once()
	middlew := new(httpmock.MockMiddleware)
	middlew.On("AbortTransactionWithError", http.CreateUnauthorizedError()).Once()
	ctx := &http.Context{
		Reader:     reader,
		Middleware: middlew,
	}
	m := admin.CreateMiddleware([]byte("test"))

	m(ctx)

	reader.AssertExpectations(t)
	middlew.AssertExpectations(t)
}

func TestAuthMiddleware_GivenOneWithTheSameSecretAsToken_WhenAuthorizing_ThenItContinuesTransaction(t *testing.T) {
	reader := new(httpmock.MockReader)
	reader.On("GetHeader", middleware.HeaderAuthorization).Return("test").Once()
	middlew := new(httpmock.MockMiddleware)
	middlew.On("NextHandler").Once()
	ctx := &http.Context{
		Reader:     reader,
		Middleware: middlew,
	}
	m := admin.CreateMiddleware([]byte("test"))

	m(ctx)

	reader.AssertExpectations(t)
	middlew.AssertExpectations(t)
}
