package auth

import (
	"github.com/arpb2/C-3PO/pkg/domain/session/repository"
	"github.com/stretchr/testify/mock"
)

type MockTokenHandler struct {
	mock.Mock
}

func (t *MockTokenHandler) Create(token *repository.Token) (tokenStr string, err error) {
	args := t.Called(token)

	tokenStr = args.String(0)

	err = args.Error(1)

	return
}

func (t *MockTokenHandler) Retrieve(tokn string) (tkn *repository.Token, err error) {
	args := t.Called(tokn)

	tokenParam := args.Get(0)
	if tokenParam != nil {
		tkn = tokenParam.(*repository.Token)
	}

	err = args.Error(1)

	return
}
