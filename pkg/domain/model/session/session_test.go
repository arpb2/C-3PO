package session_test

import (
	"encoding/json"
	"testing"

	session2 "github.com/arpb2/C-3PO/pkg/domain/model/session"

	"github.com/stretchr/testify/assert"
)

func TestSession_ToJson(t *testing.T) {
	expectedJson := `{"user_id":1000,"token":"test_token"}`

	session := &session2.Session{
		Token:  "test_token",
		UserId: uint(1000),
	}

	data, err := json.Marshal(session)

	assert.Nil(t, err)
	assert.Equal(t, expectedJson, string(data))
}

func TestSession_FromJson(t *testing.T) {
	expectedSession := &session2.Session{
		Token:  "test_token",
		UserId: uint(1000),
	}

	data := `{"user_id":1000,"token":"test_token"}`
	var session session2.Session

	err := json.Unmarshal([]byte(data), &session)

	assert.Nil(t, err)

	assert.Equal(t, expectedSession.UserId, session.UserId)
	assert.Equal(t, expectedSession.Token, session.Token)
}
