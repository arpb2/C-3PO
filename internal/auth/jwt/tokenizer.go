package jwt

import (
	"fmt"
	"github.com/arpb2/C-3PO/api/auth"
	"github.com/arpb2/C-3PO/api/http_wrapper"
	"github.com/dgrijalva/jwt-go"
	"os"
)

func CreateTokenHandler() auth.TokenHandler {
	return &TokenHandler{
		Secret: FetchJwtSecret(),
	}
}

func FetchJwtSecret() []byte {
	osValue := os.Getenv("JWT_SECRET")
	if osValue == "" {
		osValue = "52bfd2de0a2e69dff4517518590ac32a46bd76606ec22a258f99584a6e70aca2" // "test_secret" SHA256
		fmt.Printf("[WARN] Setting test secret '%s' for JWT as secret environment variable wasn't found", osValue)
	}
	return []byte(osValue)
}

type token struct {
	*auth.Token
	jwt.StandardClaims
}

type TokenHandler struct {
	Secret []byte
}

func (t TokenHandler) Create(authToken *auth.Token) (string, error) {
	claims := &token{
		Token:          authToken,
		StandardClaims: jwt.StandardClaims{},
	}

	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := tkn.SignedString(t.Secret)

	if err != nil {
		return "", http_wrapper.CreateInternalError()
	}
	return tokenString, nil
}

func (t TokenHandler) Retrieve(authToken string) (*auth.Token, error) {
	claims := &token{}

	tkn, err := jwt.ParseWithClaims(authToken, claims, func(token *jwt.Token) (interface{}, error) {
		return t.Secret, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, http_wrapper.CreateUnauthorizedError()
		}
		return nil, http_wrapper.CreateBadRequestError("malformed token")
	}

	if !tkn.Valid {
		return nil, http_wrapper.CreateUnauthorizedError()
	}

	return claims.Token, nil
}
