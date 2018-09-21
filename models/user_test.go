package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserStruct(t *testing.T) {
	var user User
	user.ID = "1"

	assert.Equal(t, user.ID, "1")
}
