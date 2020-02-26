package token

import (
	token2 "github.com/arpb2/C-3PO/pkg/domain/session/token"
	"github.com/stretchr/testify/mock"
)

type MockTokenHandler struct {
	mock.Mock
}

func (m *MockTokenHandler) Create(token *token2.Token) (string, error) {
	args := m.Called(token)
	return args.String(0), args.Error(1)
}

func (m *MockTokenHandler) Retrieve(token string) (*token2.Token, error) {
	args := m.Called(token)

	var tkn *token2.Token
	tknParam := args.Get(0)
	if tknParam != nil {
		tkn = tknParam.(*token2.Token)
	}

	return tkn, args.Error(1)
}
