package integration_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/arpb2/C-3PO/api/model"

	ginengine "github.com/arpb2/C-3PO/pkg/engine/gin"
	"github.com/arpb2/C-3PO/pkg/server"
	userlevelservice "github.com/arpb2/C-3PO/pkg/service/user_level"
	"github.com/arpb2/C-3PO/test/mock/golden"
	"github.com/stretchr/testify/assert"
)

func dial(t *testing.T) {
	connStablished := false
	for i := 0; i < 5 && !connStablished; i++ {
		fmt.Printf("Dialing localhost:8080. Retry: %d", i)
		_, err := net.DialTimeout("tcp", "localhost:8080", 500*time.Millisecond)
		if err == nil {
			connStablished = true
		} else {
			fmt.Printf("Error: %s", err.Error())
		}
	}

	assert.True(t, connStablished)
}

func Test_Get(t *testing.T) {
	// Ignite server
	engine := ginengine.New()

	defer engine.Shutdown()
	go server.StartApplication(engine)

	// Add code to retrieve
	code := "expected code"
	createdCode, err := userlevelservice.CreateService().StoreUserLevel(model.UserLevel{
		UserId:  1000,
		LevelId: 1,
		UserLevelData: model.UserLevelData{
			Code: code,
		},
	})

	assert.Nil(t, err)

	// CreateService request to perform
	req, err := http.NewRequest("GET", "http://localhost:8080/users/1000/levels/"+strconv.FormatUint(uint64(createdCode.LevelId), 10), strings.NewReader(""))
	assert.Nil(t, err)

	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEwMDB9.GVS-KC5nOCHybzzFIIH864u4KcGu-ZSd-96krqTUGWo")

	// Perform request
	dial(t)
	resp, err := http.DefaultClient.Do(req)

	// High level Assertions
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// ReadBody response body
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	assert.Nil(t, err)

	// Low level assertions
	actual := bytes.TrimSpace(bodyBytes)
	expected := golden.Get(t, actual, "ok.get_user_level.golden.json")

	assert.Equal(t, expected, actual)
}

func Test_Put(t *testing.T) {
	// Ignite server
	engine := ginengine.New()

	defer engine.Shutdown()
	go server.StartApplication(engine)

	// Add code to replace
	code := "test code"
	userLevel, err := userlevelservice.CreateService().StoreUserLevel(model.UserLevel{
		UserId:  1000,
		LevelId: 1,
		UserLevelData: model.UserLevelData{
			Code: code,
		},
	})

	assert.Nil(t, err)

	// CreateService request to perform
	data := url.Values{}
	data["code"] = []string{"some code i'm replacing"}
	data["workspace"] = []string{"some workspace i'm replacing"}
	req, err := http.NewRequest("PUT", "http://localhost:8080/users/1000/levels/"+strconv.FormatUint(uint64(userLevel.LevelId), 10), strings.NewReader(data.Encode()))
	assert.Nil(t, err)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEwMDB9.GVS-KC5nOCHybzzFIIH864u4KcGu-ZSd-96krqTUGWo")

	// Perform request
	dial(t)
	resp, err := http.DefaultClient.Do(req)

	// High level Assertions
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// ReadBody response body
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	assert.Nil(t, err)

	// Low level assertions
	actual := bytes.TrimSpace(bodyBytes)
	expected := golden.Get(t, actual, "ok.put_user_level.golden.json")

	assert.Equal(t, expected, actual)
}
