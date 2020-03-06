package validation_test

import (
	"testing"

	"github.com/arpb2/C-3PO/pkg/data/usecase/user/validation"
	"github.com/arpb2/C-3PO/pkg/domain/model/user"

	"github.com/stretchr/testify/assert"
)

func TestUserValidation_Error(t *testing.T) {
	err := validation.EmptyUser(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "no user found", err.Error())
}

func TestUserValidation_Success(t *testing.T) {
	err := validation.EmptyUser(&user.AuthenticatedUser{})

	assert.Nil(t, err)
}
