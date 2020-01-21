package auth

type Token struct {
	UserId uint
}

type TokenError struct {
	Error error
	Status int
}

type TokenHandler interface {
	Create(token *Token) (string, *TokenError)
	Retrieve(token string) (*Token, *TokenError)
}
