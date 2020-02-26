package validation_test

import (
	"testing"

	model2 "github.com/arpb2/C-3PO/pkg/domain/user/model"

	"github.com/arpb2/C-3PO/pkg/presentation/user/validation"

	"github.com/stretchr/testify/assert"
)

func TestEmptySurnameValidation_Error(t *testing.T) {
	err := validation.EmptySurname(&model2.AuthenticatedUser{
		User:     model2.User{},
		Password: "",
	})

	assert.NotNil(t, err)
	assert.Equal(t, "no 'surname' provided", err.Error())
}

func TestEmptySurnameValidation_Success(t *testing.T) {
	err := validation.EmptySurname(&model2.AuthenticatedUser{
		User: model2.User{
			Surname: "TestSurname",
		},
		Password: "",
	})

	assert.Nil(t, err)
}
