package jwt

import (
	"github.com/arpb2/C-3PO/pkg/domain/architecture/http"
	"github.com/arpb2/C-3PO/pkg/domain/session/repository"
	"github.com/dgrijalva/jwt-go"
)

func CreateTokenRepository(secret []byte) repository.TokenRepository {
	return &tokenRepository{
		Secret: secret,
	}
}

type token struct {
	*repository.Token
	jwt.StandardClaims
}

type tokenRepository struct {
	Secret []byte
}

func (t tokenRepository) Create(authToken *repository.Token) (string, error) {
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

func (t tokenRepository) Retrieve(authToken string) (*repository.Token, error) {
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
