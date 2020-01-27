package jwt

import (
	"fmt"
	"os"

	"github.com/arpb2/C-3PO/api/auth"
	"github.com/arpb2/C-3PO/api/http"
	"github.com/dgrijalva/jwt-go"
)

func CreateTokenHandler() auth.TokenHandler {
	return &TokenHandler{
		Secret: FetchJwtSecret(),
	}
}

func FetchJwtSecret() []byte {
	osValue := os.Getenv("JWT_SECRET")
	if osValue == "" {
		// TODO Remove this. As it's a serious security problem if we ever go to production.
		osValue = "52bfd2de0a2e69dff4517518590ac32a46bd76606ec22a258f99584a6e70aca2" // "test_secret" SHA256
		fmt.Printf("[WARN] JWT Tokenizer - Setting test secret '%s' for JWT as secret environment variable wasn't found\n", osValue)
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
		return "", http.CreateInternalError()
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
			return nil, http.CreateUnauthorizedError()
		}
		return nil, http.CreateBadRequestError("malformed token")
	}

	if !tkn.Valid {
		return nil, http.CreateUnauthorizedError()
	}

	return claims.Token, nil
}
