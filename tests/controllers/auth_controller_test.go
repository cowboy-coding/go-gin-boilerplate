package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cowboy-coding/go-gin-boilerplate/server"
	"github.com/stretchr/testify/assert"
)

func TestAuthControllerCreate(t *testing.T) {
	router := server.NewRouter()

	w := httptest.NewRecorder()

	// build request
	var jsonStr = []byte(`{"username":"admin", "password": "admin"}`)
	req, _ := http.NewRequest("POST", "/v1/auth", bytes.NewBuffer(jsonStr))
	router.ServeHTTP(w, req)

	// validate success
	assert.Equal(t, 200, w.Code)

	// validate response body
	// fmt.Println(w.Body.String())
	var response map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	// make sure err is nil
	assert.Nil(t, err)

	// make suere JWT was created
	_, exists := response["token"]
	assert.True(t, exists)
}

func TestAuthControllerCreateFail(t *testing.T) {
	router := server.NewRouter()

	w := httptest.NewRecorder()

	// build request
	var jsonStr = []byte(`{"username":"admin1", "password": "admin1"}`)
	req, _ := http.NewRequest("POST", "/v1/auth", bytes.NewBuffer(jsonStr))
	router.ServeHTTP(w, req)

	// validate failure
	assert.Equal(t, 401, w.Code)

	// validate response body
	fmt.Println(w.Body.String())
	var response map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	// make sure err is nil
	assert.Nil(t, err)

	// make sure messe is present
	_, exists := response["message"]
	assert.True(t, exists)
}
