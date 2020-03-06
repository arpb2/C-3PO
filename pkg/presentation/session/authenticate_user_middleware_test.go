package session_test

import (
	"testing"

	session2 "github.com/arpb2/C-3PO/pkg/data/repository/session"
	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/presentation/session"
	"github.com/arpb2/C-3PO/test/mock/auth"
	http2 "github.com/arpb2/C-3PO/test/mock/http"
)

func TestUserGivenAnAlreadyAuthenticatedContext_WhenRun_ThenCallsNextAndStops(t *testing.T) {
	ctxMiddle := new(http2.MockMiddleware)
	ctxMiddle.On("GetValue", "session_authenticated").Return(true).Once()
	ctxMiddle.On("NextHandler").Once()
	ctx := &http.Context{
		Middleware: ctxMiddle,
	}
	middle := session.CreateAuthenticateUserMiddleware(
		"user_id",
		nil,
	)

	middle(ctx)

	ctxMiddle.AssertExpectations(t)
}

func TestUserGivenASameUserId_WhenRun_ThenContextIsSetAndHandlerNexts(t *testing.T) {
	tokenHandler := new(auth.MockTokenHandler)
	tokenHandler.On("Retrieve", "test").Return(&session2.Token{
		UserId: 1000,
	}, nil)
	ctxReader := new(http2.MockReader)
	ctxReader.On("GetHeader", "Authorization").Return("test").Once()
	ctxReader.On("GetParameter", "user_id").Return("1000").Once()
	ctxMiddle := new(http2.MockMiddleware)
	ctxMiddle.On("GetValue", "session_authenticated").Return(false).Once()
	ctxMiddle.On("SetValue", "session_authenticated", true).Once()
	ctxMiddle.On("NextHandler").Once()
	ctx := &http.Context{
		Reader:     ctxReader,
		Middleware: ctxMiddle,
	}
	middle := session.CreateAuthenticateUserMiddleware(
		"user_id",
		tokenHandler,
	)

	middle(ctx)

	ctxMiddle.AssertExpectations(t)
	ctxReader.AssertExpectations(t)
	tokenHandler.AssertExpectations(t)
}

func TestUserGivenADifferentUserId_WhenRun_ThenHaltsWithUnauthorized(t *testing.T) {
	tokenHandler := new(auth.MockTokenHandler)
	tokenHandler.On("Retrieve", "test").Return(&session2.Token{
		UserId: 1000,
	}, nil)
	ctxReader := new(http2.MockReader)
	ctxReader.On("GetHeader", "Authorization").Return("test").Once()
	ctxReader.On("GetParameter", "user_id").Return("1").Once()
	ctxMiddle := new(http2.MockMiddleware)
	ctxMiddle.On("GetValue", "session_authenticated").Return(false).Once()
	ctxMiddle.On("AbortTransactionWithError", http.CreateUnauthorizedError()).Once()
	ctx := &http.Context{
		Reader:     ctxReader,
		Middleware: ctxMiddle,
	}
	middle := session.CreateAuthenticateUserMiddleware(
		"user_id",
		tokenHandler,
	)

	middle(ctx)

	ctxMiddle.AssertExpectations(t)
	ctxReader.AssertExpectations(t)
	tokenHandler.AssertExpectations(t)
}
