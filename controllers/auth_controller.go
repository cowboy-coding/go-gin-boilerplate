package controllers

import (
	"github.com/gin-gonic/gin"
)

type AuthController struct{}

func (u AuthController) Create(c *gin.Context) {
}

func (u AuthController) Update(c *gin.Context) {
}

func (u AuthController) Destroy(c *gin.Context) {
}

// var user models.User

// // bind Post params to user
// if err := c.ShouldBindJSON(&user); err != nil {
// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 	return
// }
// // validate username and passwort against DB

// // just create a token
// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 	"username": user.Username,
// 	"password": user.Password,
// })

// tokenString, error := token.SignedString([]byte("secret"))
// if error != nil {
// 	fmt.Println(error)
// }

// c.JSON(http.StatusOK, JwtToken{Token: tokenString})
