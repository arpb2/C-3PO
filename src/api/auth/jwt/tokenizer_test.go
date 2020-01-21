package jwt

import (
	"github.com/arpb2/C-3PO/src/api/auth"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

var defaultSha = "52bfd2de0a2e69dff4517518590ac32a46bd76606ec22a258f99584a6e70aca2"

var DefaultTokenHandler = TokenHandler{
	Secret: FetchJwtSecret(),
}

func TestSecret_DefaultValue(t *testing.T) {
	err := os.Unsetenv("JWT_SECRET")

	assert.NoError(t, err)

	secret := FetchJwtSecret()

	assert.Equal(t, []byte(defaultSha), secret)
}

func TestSecret_UsesOsEnv(t *testing.T) {
	value := "some secret value"
	err := os.Setenv("JWT_SECRET", value)

	assert.NoError(t, err)

	defer os.Unsetenv("JWT_SECRET")

	secret := FetchJwtSecret()

	assert.Equal(t, []byte(value), secret)
}

var expectedToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEwMDB9.GVS-KC5nOCHybzzFIIH864u4KcGu-ZSd-96krqTUGWo"

func TestCreate_CreatesExpectedToken(t *testing.T) {
	token, err := DefaultTokenHandler.Create(&auth.Token{
		UserId: 1000,
	})

	assert.Nil(t, err)
	assert.Equal(t, expectedToken, token)
}

func TestRetrieve_GetsExpectedUserId(t *testing.T) {
	token, err := DefaultTokenHandler.Retrieve(expectedToken)

	assert.Nil(t, err)
	assert.Equal(t, uint(1000), token.UserId)
}

func TestRetrieve_OnBadToken(t *testing.T) {
	token, err := DefaultTokenHandler.Retrieve("bad token")

	assert.NotNil(t, err)
	assert.Equal(t, http.StatusBadRequest, err.Status)
	assert.Nil(t, token)
}