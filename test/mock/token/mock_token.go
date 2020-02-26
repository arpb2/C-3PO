package token

import (
	"github.com/arpb2/C-3PO/pkg/domain/session/repository"
	"github.com/stretchr/testify/mock"
)

type MockTokenHandler struct {
	mock.Mock
}

func (m *MockTokenHandler) Create(token *repository.Token) (string, error) {
	args := m.Called(token)
	return args.String(0), args.Error(1)
}

func (m *MockTokenHandler) Retrieve(token string) (*repository.Token, error) {
	args := m.Called(token)

	var tkn *repository.Token
	tknParam := args.Get(0)
	if tknParam != nil {
		tkn = tknParam.(*repository.Token)
	}

	return tkn, args.Error(1)
}
