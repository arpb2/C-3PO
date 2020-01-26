package integration_test

import (
	"bytes"
	"fmt"
	"github.com/arpb2/C-3PO/internal/engine/gin"
	"github.com/arpb2/C-3PO/internal/server"
	code_service "github.com/arpb2/C-3PO/internal/service/code"
	"github.com/arpb2/C-3PO/test/golden"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"
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
	engine := gin_engine.New()

	defer engine.Shutdown()
	go server.StartApplication(engine)

	// Add code to retrieve
	code := "expected code"
	createdCode, err := code_service.CreateService().CreateCode(uint(1000), code)

	assert.Nil(t, err)

	// CreateService request to perform
	req, err := http.NewRequest("GET", "http://localhost:8080/users/1000/codes/"+strconv.FormatUint(uint64(createdCode.Id), 10), strings.NewReader(""))
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
	expected := golden.Get(t, actual, "ok.get_code.golden.json")

	assert.Equal(t, expected, actual)
}

func Test_Post(t *testing.T) {
	// Ignite server
	engine := gin_engine.New()

	defer engine.Shutdown()
	go server.StartApplication(engine)

	// CreateService request to perform
	data := url.Values{}
	data["code"] = []string{"some code i'm uploading"}
	req, err := http.NewRequest("POST", "http://localhost:8080/users/1000/codes", strings.NewReader(data.Encode()))
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
	expected := golden.Get(t, actual, "ok.post_code.golden.json")

	assert.Equal(t, expected, actual)
}

func Test_Put(t *testing.T) {
	// Ignite server
	engine := gin_engine.New()

	defer engine.Shutdown()
	go server.StartApplication(engine)

	// Add code to replace
	code := "test code"
	codeModel, err := code_service.CreateService().CreateCode(uint(1000), code)

	assert.Nil(t, err)

	// CreateService request to perform
	data := url.Values{}
	data["code"] = []string{"some code i'm replacing"}
	req, err := http.NewRequest("PUT", "http://localhost:8080/users/1000/codes/"+strconv.FormatUint(uint64(codeModel.Id), 10), strings.NewReader(data.Encode()))
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
	expected := golden.Get(t, actual, "ok.put_code.golden.json")

	assert.Equal(t, expected, actual)
}
