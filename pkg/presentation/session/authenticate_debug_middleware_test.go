package session_test

import (
	"testing"

	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/arpb2/C-3PO/pkg/presentation/session"
	http2 "github.com/arpb2/C-3PO/test/mock/http"
)

func TestDebugGivenAnAlreadyAuthenticatedContext_WhenRun_ThenCallsNextAndStops(t *testing.T) {
	ctxMiddle := new(http2.MockMiddleware)
	ctxMiddle.On("GetValue", "session_authenticated").Return(true).Once()
	ctxMiddle.On("NextHandler").Once()
	ctx := &http.Context{
		Middleware: ctxMiddle,
	}
	middle := session.CreateAuthenticateDebugMiddleware()

	middle(ctx)

	ctxMiddle.AssertExpectations(t)
}

func TestDebugGivenASameSecretAndHeader_WhenRun_ThenContextIsSetAndHandlerNexts(t *testing.T) {
	ctxReader := new(http2.MockReader)
	ctxReader.On("GetHeader", "Authorization").Return("DEBUG").Once()
	ctxMiddle := new(http2.MockMiddleware)
	ctxMiddle.On("GetValue", "session_authenticated").Return(false).Once()
	ctxMiddle.On("SetValue", "session_authenticated", true).Once()
	ctxMiddle.On("NextHandler").Once()
	ctx := &http.Context{
		Reader:     ctxReader,
		Middleware: ctxMiddle,
	}
	middle := session.CreateAuthenticateDebugMiddleware()

	middle(ctx)

	ctxMiddle.AssertExpectations(t)
	ctxReader.AssertExpectations(t)
}

func TestDebugGivenADifferentSecretAndHeader_WhenRun_ThenHaltsWithUnauthorized(t *testing.T) {
	ctxReader := new(http2.MockReader)
	ctxReader.On("GetHeader", "Authorization").Return("not test").Once()
	ctxMiddle := new(http2.MockMiddleware)
	ctxMiddle.On("GetValue", "session_authenticated").Return(false).Once()
	ctx := &http.Context{
		Reader:     ctxReader,
		Middleware: ctxMiddle,
	}
	middle := session.CreateAuthenticateDebugMiddleware()

	middle(ctx)

	ctxMiddle.AssertExpectations(t)
	ctxReader.AssertExpectations(t)
}
