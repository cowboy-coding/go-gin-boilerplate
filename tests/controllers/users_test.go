package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cowboy-coding/go-gin-boilerplate/server"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUsersControllerWithoutAuth(t *testing.T) {
	// {"code":401,"message":"auth header is empty"}
	// log.Print(response)
	body := gin.H{
		"code":    401,
		"message": "auth header is empty",
	}

	router := server.NewRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/users/1", nil)
	router.ServeHTTP(w, req)

	// Convert the JSON response to a map
	var response map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	value, _ := response["message"]

	assert.Equal(t, 401, w.Code)
	assert.Equal(t, body["code"], 401)
	assert.Nil(t, err)
	assert.Equal(t, body["message"], value)
}
