package validation_test

import (
	"testing"

	"github.com/arpb2/C-3PO/pkg/data/usecase/user/validation"
	"github.com/arpb2/C-3PO/pkg/domain/model/user"

	"github.com/stretchr/testify/assert"
)

func TestEmptySurnameValidation_Error(t *testing.T) {
	err := validation.EmptySurname(&user.AuthenticatedUser{
		User:     user.User{},
		Password: "",
	})

	assert.NotNil(t, err)
	assert.Equal(t, "no 'surname' provided", err.Error())
}

func TestEmptySurnameValidation_Success(t *testing.T) {
	err := validation.EmptySurname(&user.AuthenticatedUser{
		User: user.User{
			Surname: "TestSurname",
		},
		Password: "",
	})

	assert.Nil(t, err)
}
