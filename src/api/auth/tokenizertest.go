package auth

import "github.com/stretchr/testify/mock"

type MockTokenHandler struct{
	mock.Mock
}

func (t MockTokenHandler) Create(token *Token) (tokenStr string, err error) {
	args := t.Called(token)

	tokenStr = args.String(0)

	err = args.Error(1)

	return
}

func (t MockTokenHandler) Retrieve(token string) (tkn *Token, err error) {
	args := t.Called(token)

	tokenParam := args.Get(0)
	if tokenParam != nil {
		tkn = tokenParam.(*Token)
	}

	err = args.Error(1)

	return
}
