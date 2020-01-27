package model_test

import (
	"encoding/json"
	"testing"

	"github.com/arpb2/C-3PO/api/model"
	"github.com/stretchr/testify/assert"
)

func TestSession_ToJson(t *testing.T) {
	expectedJson := `{"user_id":1000,"token":"test_token"}`

	session := &model.Session{
		Token:  "test_token",
		UserId: uint(1000),
	}

	data, err := json.Marshal(session)

	assert.Nil(t, err)
	assert.Equal(t, expectedJson, string(data))
}

func TestSession_FromJson(t *testing.T) {
	expectedSession := &model.Session{
		Token:  "test_token",
		UserId: uint(1000),
	}

	data := `{"user_id":1000,"token":"test_token"}`
	var session model.Session

	err := json.Unmarshal([]byte(data), &session)

	assert.Nil(t, err)

	assert.Equal(t, expectedSession.UserId, session.UserId)
	assert.Equal(t, expectedSession.Token, session.Token)
}
