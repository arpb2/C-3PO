package jwt

import (
	"github.com/arpb2/C-3PO/pkg/data/repository/session"
	"github.com/arpb2/C-3PO/pkg/domain/http"
	"github.com/dgrijalva/jwt-go"
)

func CreateTokenRepository(secret []byte) session.TokenRepository {
	return &tokenRepository{
		Secret: secret,
	}
}

type token struct {
	*session.Token
	jwt.StandardClaims
}

type tokenRepository struct {
	Secret []byte
}

func (t tokenRepository) Create(authToken *session.Token) (string, error) {
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

func (t tokenRepository) Retrieve(authToken string) (*session.Token, error) {
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
