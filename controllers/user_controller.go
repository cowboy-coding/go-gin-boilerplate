package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersController struct{}

func (u UsersController) Show(c *gin.Context) {
	c.String(http.StatusOK, "User.show")
}
