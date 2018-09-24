package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersController struct{}

func (u UsersController) Show(c *gin.Context) {
	c.String(http.StatusOK, "User.show")
}

func (u UsersController) SignUp(c *gin.Context) {
	fmt.Println("---UsersController")
	c.String(http.StatusOK, "User.show")
}
