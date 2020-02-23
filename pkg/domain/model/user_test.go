package model_test

import (
	"encoding/json"
	"testing"

	"github.com/arpb2/C-3PO/pkg/domain/model"
	"github.com/stretchr/testify/assert"
)

func TestUser_ToJson(t *testing.T) {
	expectedJson := `{"id":0,"email":"test@email.com","name":"test name","surname":"test surname"}`

	user := &model.User{
		Id:      0,
		Name:    "test name",
		Surname: "test surname",
		Email:   "test@email.com",
	}

	data, err := json.Marshal(user)

	assert.Nil(t, err)
	assert.Equal(t, expectedJson, string(data))
}

func TestUser_FromJson(t *testing.T) {
	expectedUser := &model.User{
		Id:      0,
		Name:    "test name",
		Surname: "test surname",
		Email:   "test@email.com",
	}

	data := `{"id":0,"email":"test@email.com","name":"test name","surname":"test surname"}`
	var user model.User

	err := json.Unmarshal([]byte(data), &user)

	assert.Nil(t, err)

	assert.Equal(t, expectedUser.Id, user.Id)
	assert.Equal(t, expectedUser.Email, user.Email)
	assert.Equal(t, expectedUser.Name, user.Name)
	assert.Equal(t, expectedUser.Surname, user.Surname)
}
