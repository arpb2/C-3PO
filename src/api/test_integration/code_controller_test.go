package integration_test

import (
	"bytes"
	"github.com/arpb2/C-3PO/src/api/controller/code"
	"github.com/arpb2/C-3PO/src/api/golden"
	"github.com/arpb2/C-3PO/src/api/server"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"
)

type SharedInMemoryCodeService struct{}
func (s *SharedInMemoryCodeService) Read(userId string, codeId string) (code *string, err error) {
	c := "expected code"
	return &c, nil
}

func (s *SharedInMemoryCodeService) Write(userId string, code *string) (codeId string, err error) {
	return "456", nil
}

func (s *SharedInMemoryCodeService) Replace(userId string, codeId string, code *string) error {
	return nil
}

func init() {
	code.Service = &SharedInMemoryCodeService{}
}

func Test_Get(t *testing.T) {
	// Ignite server
	defer server.Engine.Shutdown()
	go server.StartApplication()

	// Create request to perform
	req, err := http.NewRequest("GET", "http://localhost:8080/users/1000/codes/1000", strings.NewReader(""))
	assert.Nil(t, err)

	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEwMDB9.GVS-KC5nOCHybzzFIIH864u4KcGu-ZSd-96krqTUGWo")

	// Perform request
	time.Sleep(1 * time.Second) // http doesn't have an ignite hook, so we halt for a second
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
	defer server.Engine.Shutdown()
	go server.StartApplication()

	// Create request to perform
	data := url.Values{}
	data["code"] = []string{"some code i'm uploading"}
	req, err := http.NewRequest("POST", "http://localhost:8080/users/1000/codes", strings.NewReader(data.Encode()))
	assert.Nil(t, err)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEwMDB9.GVS-KC5nOCHybzzFIIH864u4KcGu-ZSd-96krqTUGWo")

	// Perform request
	time.Sleep(1 * time.Second) // http doesn't have an ignite hook, so we halt for a second
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
	defer server.Engine.Shutdown()
	go server.StartApplication()

	// Create request to perform
	data := url.Values{}
	data["code"] = []string{"some code i'm replacing"}
	req, err := http.NewRequest("PUT", "http://localhost:8080/users/1000/codes/1000", strings.NewReader(data.Encode()))
	assert.Nil(t, err)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEwMDB9.GVS-KC5nOCHybzzFIIH864u4KcGu-ZSd-96krqTUGWo")

	// Perform request
	time.Sleep(1 * time.Second) // http doesn't have an ignite hook, so we halt for a second
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
