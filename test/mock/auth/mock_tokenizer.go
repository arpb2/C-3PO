package auth

import (
	"github.com/arpb2/C-3PO/pkg/data/repository/session"
	"github.com/stretchr/testify/mock"
)

type MockTokenHandler struct {
	mock.Mock
}

func (t *MockTokenHandler) Create(token *session.Token) (tokenStr string, err error) {
	args := t.Called(token)

	tokenStr = args.String(0)

	err = args.Error(1)

	return
}

func (t *MockTokenHandler) Retrieve(tokn string) (tkn *session.Token, err error) {
	args := t.Called(tokn)

	tokenParam := args.Get(0)
	if tokenParam != nil {
		tkn = tokenParam.(*session.Token)
	}

	err = args.Error(1)

	return
}
