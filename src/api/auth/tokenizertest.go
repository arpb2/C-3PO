package auth

import "github.com/stretchr/testify/mock"

type MockTokenHandler struct{
	mock.Mock
}

func (t MockTokenHandler) Create(token *Token) (tokenStr string, err *TokenError) {
	args := t.Called(token)

	tokenStr = args.String(0)

	errParam := args.Get(1)
	if errParam != nil {
		err = errParam.(*TokenError)
	}

	return
}

func (t MockTokenHandler) Retrieve(token string) (tkn *Token, err *TokenError) {
	args := t.Called(token)

	tokenParam := args.Get(0)
	if tokenParam != nil {
		tkn = tokenParam.(*Token)
	}

	errParam := args.Get(1)
	if errParam != nil {
		err = errParam.(*TokenError)
	}

	return
}
