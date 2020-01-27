package user_test

import (
	"testing"

	"github.com/arpb2/C-3PO/api/model"
	uservalidation "github.com/arpb2/C-3PO/pkg/validation/user"
	"github.com/stretchr/testify/assert"
)

func TestUserValidation_Error(t *testing.T) {
	err := uservalidation.EmptyUser(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "no user found", err.Error())
}

func TestUserValidation_Success(t *testing.T) {
	err := uservalidation.EmptyUser(&model.AuthenticatedUser{})

	assert.Nil(t, err)
}
