package integration_test

import (
	"bytes"
	"fmt"
	"github.com/arpb2/C-3PO/src/api/engine/c3po"
	"github.com/arpb2/C-3PO/src/api/golden"
	"github.com/arpb2/C-3PO/src/api/server"
	"github.com/arpb2/C-3PO/src/api/service/code_service"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"
)

func dial(t *testing.T) {
	connStablished := false
	for i := 0 ; i < 5 && !connStablished ; i++ {
		fmt.Printf("Dialing localhost:8080. Retry: %d", i)
		_, err := net.DialTimeout("tcp","localhost:8080", 500 * time.Millisecond)
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
	engine := c3po.New()

	defer engine.Shutdown()
	go server.StartApplication(engine)

	// Add code to retrieve
	code := "expected code"
	codeId, err := code_service.GetService().CreateCode("1000", &code)

	assert.Nil(t, err)

	// Create request to perform
	req, err := http.NewRequest("GET", "http://localhost:8080/users/1000/codes/" + codeId, strings.NewReader(""))
	assert.Nil(t, err)

	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEwMDB9.GVS-KC5nOCHybzzFIIH864u4KcGu-ZSd-96krqTUGWo")

	// Perform request
	dial(t)
	resp, err := http.DefaultClient.Do(req)

	// High level Assertions
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Parse response body
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
	engine := c3po.New()

	defer engine.Shutdown()
	go server.StartApplication(engine)

	// Create request to perform
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

	// Parse response body
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
	engine := c3po.New()

	defer engine.Shutdown()
	go server.StartApplication(engine)

	// Add code to replace
	code := "test code"
	codeId, err := code_service.GetService().CreateCode("1000", &code)

	assert.Nil(t, err)

	// Create request to perform
	data := url.Values{}
	data["code"] = []string{"some code i'm replacing"}
	req, err := http.NewRequest("PUT", "http://localhost:8080/users/1000/codes/" + codeId, strings.NewReader(data.Encode()))
	assert.Nil(t, err)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEwMDB9.GVS-KC5nOCHybzzFIIH864u4KcGu-ZSd-96krqTUGWo")

	// Perform request
	dial(t)
	resp, err := http.DefaultClient.Do(req)

	// High level Assertions
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Parse response body
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	assert.Nil(t, err)

	// Low level assertions
	actual := bytes.TrimSpace(bodyBytes)
	expected := golden.Get(t, actual, "ok.put_code.golden.json")

	assert.Equal(t, expected, actual)
}
