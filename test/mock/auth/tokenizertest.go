package auth

import (
	"github.com/arpb2/C-3PO/pkg/domain/auth"
	"github.com/stretchr/testify/mock"
)

type MockTokenHandler struct {
	mock.Mock
}

func (t *MockTokenHandler) Create(token *auth.Token) (tokenStr string, err error) {
	args := t.Called(token)

	tokenStr = args.String(0)

	err = args.Error(1)

	return
}

func (t *MockTokenHandler) Retrieve(token string) (tkn *auth.Token, err error) {
	args := t.Called(token)

	tokenParam := args.Get(0)
	if tokenParam != nil {
		tkn = tokenParam.(*auth.Token)
	}

	err = args.Error(1)

	return
}
