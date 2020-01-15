package jwt

import (
	"fmt"
	"github.com/arpb2/C-3PO/src/api/auth"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
)

var DefaultTokenHandler = TokenHandler{
	secret: FetchJwtSecret(),
}

func FetchJwtSecret() []byte {
	osValue := os.Getenv("JWT_SECRET")
	if osValue == "" {
		osValue = "52bfd2de0a2e69dff4517518590ac32a46bd76606ec22a258f99584a6e70aca2" // "test_secret" SHA256
		fmt.Printf("[WARN] Setting test secret '%s' for JWT as secret environment variable wasn't found", osValue)
	}
	return []byte(osValue)
}

type token struct{
	auth.Token
	jwt.StandardClaims
}

type TokenHandler struct{
	secret []byte
}

func (t TokenHandler) Create(authToken auth.Token) (*string, *auth.TokenError) {
	claims := &token{
		Token: authToken,
		StandardClaims: jwt.StandardClaims{},
	}

	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := tkn.SignedString(t.secret)

	if err != nil {
		return nil, &auth.TokenError{
			Error:  err,
			Status: http.StatusInternalServerError,
		}
	}
	return &tokenString, nil
}

func (t TokenHandler) Retrieve(authToken string) (*auth.Token, *auth.TokenError) {
	claims := &token{}

	tkn, err := jwt.ParseWithClaims(authToken, claims, func(token *jwt.Token) (interface{}, error) {
		return t.secret, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, &auth.TokenError{
				Error:  err,
				Status: http.StatusUnauthorized,
			}
		}
		return nil, &auth.TokenError{
			Error: err,
			Status: http.StatusBadRequest,
		}
	}

	if !tkn.Valid {
		return nil, &auth.TokenError{
			Error:  err,
			Status: http.StatusUnauthorized,
		}
	}

	return &claims.Token, nil
}