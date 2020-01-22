package auth

type Token struct {
	UserId uint
}

type TokenHandler interface {
	Create(token *Token) (string, error)
	Retrieve(token string) (*Token, error)
}
