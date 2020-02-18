package integration_test

// TODO: We need the level endpoint before we can do this. We can't mock because it's integration tests.
/*
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
	// CreateService request to perform
	req, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:8080/users/%d/levels/%d", 1, 1), strings.NewReader(""))
	assert.Nil(t, err)

	req.Header.Set("Authorization", generateToken(1))

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

func generateToken(userId uint) string {
	str, err := jwt.TokenHandler{Secret: jwt.FetchJwtSecret()}.Create(&auth.Token{UserId: userId})
	if err != nil {
		panic(err)
	}
	return str
}

func Test_Put(t *testing.T) {
	// CreateService request to perform
	data := url.Values{}
	data["code"] = []string{"some code i'm replacing"}
	data["workspace"] = []string{"some workspace i'm replacing"}
	req, err := http.NewRequest("PUT", fmt.Sprintf("http://localhost:8080/users/%d/levels/%d", 1, 1), strings.NewReader(data.Encode()))
	assert.Nil(t, err)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", generateToken(1))

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
*/
