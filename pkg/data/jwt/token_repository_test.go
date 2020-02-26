package jwt

import (
	"github.com/arpb2/C-3PO/pkg/domain/session/repository"
	"net/http"
	"os"
	"testing"

	httpwrapper "github.com/arpb2/C-3PO/pkg/domain/architecture/http"
	"github.com/stretchr/testify/assert"
	"golang.org/x/xerrors"
)

var defaultSha = "52bfd2de0a2e69dff4517518590ac32a46bd76606ec22a258f99584a6e70aca2"

var DefaultTokenHandler = tokenRepository{
	Secret: []byte("52bfd2de0a2e69dff4517518590ac32a46bd76606ec22a258f99584a6e70aca2"),
}

func TestSecret_DefaultValue(t *testing.T) {
	err := os.Unsetenv("JWT_SECRET")

	assert.NoError(t, err)

	secret := []byte("52bfd2de0a2e69dff4517518590ac32a46bd76606ec22a258f99584a6e70aca2")

	assert.Equal(t, []byte(defaultSha), secret)
}

var expectedToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEwMDB9.GVS-KC5nOCHybzzFIIH864u4KcGu-ZSd-96krqTUGWo"

func TestCreate_CreatesExpectedToken(t *testing.T) {
	token, err := DefaultTokenHandler.Create(&repository.Token{
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

	var httpError httpwrapper.Error
	assert.True(t, xerrors.As(err, &httpError))
	assert.Equal(t, http.StatusBadRequest, httpError.Code)

	assert.Nil(t, token)
}
