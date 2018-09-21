package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cowboy-coding/go-gin-boilerplate/server"
	"github.com/stretchr/testify/assert"
)

func TestAuthControllerCreate(t *testing.T) {
	router := server.NewRouter()

	w := httptest.NewRecorder()

	values := map[string]string{"username": "foo", "password": "bar"}
	s, _ := json.Marshal(values)
	req, _ := http.NewRequest("POST", "/v1/auth", bytes.NewBuffer(s))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	// make suere err is nil
	assert.Nil(t, err)

	// make suere JWT was created
	value, exists := response["token"]
	assert.True(t, exists)
}
