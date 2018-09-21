package controllers

import (
	"net/http"

	"github.com/cowboy-coding/go-gin-boilerplate/models"
	"github.com/gin-gonic/gin"
)

type UsersController struct{}

var userModel = new(models.User)

func (u UsersController) Show(c *gin.Context) {
	c.String(http.StatusOK, "User.show")
}
