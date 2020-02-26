package token

type Token struct {
	UserId uint
}

type Handler interface {
	Create(token *Token) (string, error)
	Retrieve(token string) (*Token, error)
}
