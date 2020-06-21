package user_test

import (
	"encoding/json"
	"testing"

	user2 "github.com/arpb2/C-3PO/pkg/domain/model/user"

	"github.com/stretchr/testify/assert"
)

func TestAuthenticatedUser_ToJson(t *testing.T) {
	expectedJson := `{"id":0,"type":"student","classroom_id":0,"email":"test@email.com","name":"test name","surname":"test surname","password":"test password"}`

	user := &user2.AuthenticatedUser{
		User: user2.User{
			Id:      0,
			ClassroomID: 0,
			Type:    user2.TypeStudent,
			Name:    "test name",
			Surname: "test surname",
			Email:   "test@email.com",
		},
		Password: "test password",
	}

	data, err := json.Marshal(user)

	assert.Nil(t, err)
	assert.Equal(t, expectedJson, string(data))
}

func TestAuthenticatedUser_FromJson(t *testing.T) {
	expectedUser := &user2.AuthenticatedUser{
		User: user2.User{
			Id:      0,
			Type:    user2.TypeStudent,
			Name:    "test name",
			Surname: "test surname",
			Email:   "test@email.com",
		},
		Password: "test password",
	}

	data := `{"id":0,"type":"student","classroom_id":0,"email":"test@email.com","name":"test name","surname":"test surname","password":"test password"}`
	var user user2.AuthenticatedUser

	err := json.Unmarshal([]byte(data), &user)

	assert.Nil(t, err)

	assert.Equal(t, expectedUser.Id, user.Id)
	assert.Equal(t, expectedUser.Email, user.Email)
	assert.Equal(t, expectedUser.Name, user.Name)
	assert.Equal(t, expectedUser.Surname, user.Surname)
	assert.Equal(t, expectedUser.Password, user.Password)
}
