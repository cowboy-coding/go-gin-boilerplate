package services

import (
	jwt "github.com/appleboy/gin-jwt"
	"github.com/cowboy-coding/go-gin-boilerplate/forms"
	"github.com/cowboy-coding/go-gin-boilerplate/models"
	"github.com/gin-gonic/gin"
)

type AuthService struct{}

var identityKey = "id"

func (u AuthService) PayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(*models.User); ok {
		return jwt.MapClaims{
			identityKey: v.UserName,
		}
	}
	return jwt.MapClaims{}
}

func (u AuthService) IdentityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return &models.User{
		UserName: claims["id"].(string),
	}
}

func (u AuthService) Authenticator(c *gin.Context) (interface{}, error) {
	var loginVals forms.LoginForm
	if err := c.ShouldBind(&loginVals); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	userID := loginVals.Username
	password := loginVals.Password

	if (userID == "admin" && password == "admin") || (userID == "test" && password == "test") {
		return &models.User{
			UserName:  userID,
			LastName:  "Bo-Yi",
			FirstName: "Wu",
		}, nil
	}

	return nil, jwt.ErrFailedAuthentication
}

func (u AuthService) Authorizator(data interface{}, c *gin.Context) bool {
	if v, ok := data.(*models.User); ok && v.UserName == "admin" {
		return true
	}

	return false
}

func (u AuthService) Unauthorized(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}
