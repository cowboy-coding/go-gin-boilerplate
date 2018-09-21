package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cowboy-coding/go-gin-boilerplate/server"
	"github.com/stretchr/testify/assert"
)

func TestUsersControllerShow(t *testing.T) {
	router := server.NewRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/users/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "User.show", w.Body.String())
}
