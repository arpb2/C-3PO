package jwt

import (
	"github.com/arpb2/C-3PO/pkg/domain/infrastructure/http"
	token2 "github.com/arpb2/C-3PO/pkg/domain/session/token"
	"github.com/dgrijalva/jwt-go"
)

func CreateTokenHandler(secret []byte) token2.Handler {
	return &TokenHandler{
		Secret: secret,
	}
}

type token struct {
	*token2.Token
	jwt.StandardClaims
}

type TokenHandler struct {
	Secret []byte
}

func (t TokenHandler) Create(authToken *token2.Token) (string, error) {
	claims := &token{
		Token:          authToken,
		StandardClaims: jwt.StandardClaims{},
	}

	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := tkn.SignedString(t.Secret)

	if err != nil {
		return "", http.CreateInternalError()
	}
	return tokenString, nil
}

func (t TokenHandler) Retrieve(authToken string) (*token2.Token, error) {
	claims := &token{}

	tkn, err := jwt.ParseWithClaims(authToken, claims, func(token *jwt.Token) (interface{}, error) {
		return t.Secret, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, http.CreateUnauthorizedError()
		}
		return nil, http.CreateBadRequestError("malformed token")
	}

	if !tkn.Valid {
		return nil, http.CreateUnauthorizedError()
	}

	return claims.Token, nil
}
