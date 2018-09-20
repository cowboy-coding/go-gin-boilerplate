package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cowboy-coding/go-gin-boilerplate/server"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	router := server.NewRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Working!", w.Body.String())
}
