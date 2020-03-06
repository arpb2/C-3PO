package session

type Token struct {
	UserId uint
}

type TokenRepository interface {
	Create(token *Token) (string, error)
	Retrieve(token string) (*Token, error)
}
