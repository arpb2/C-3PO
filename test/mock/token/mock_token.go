package token

import (
	"github.com/arpb2/C-3PO/pkg/data/repository/session"
	"github.com/stretchr/testify/mock"
)

type MockTokenHandler struct {
	mock.Mock
}

func (m *MockTokenHandler) Create(token *session.Token) (string, error) {
	args := m.Called(token)
	return args.String(0), args.Error(1)
}

func (m *MockTokenHandler) Retrieve(token string) (*session.Token, error) {
	args := m.Called(token)

	var tkn *session.Token
	tknParam := args.Get(0)
	if tknParam != nil {
		tkn = tknParam.(*session.Token)
	}

	return tkn, args.Error(1)
}
