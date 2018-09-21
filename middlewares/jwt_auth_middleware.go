package middlewares

import (
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/cowboy-coding/go-gin-boilerplate/config"
	"github.com/cowboy-coding/go-gin-boilerplate/forms"
	"github.com/cowboy-coding/go-gin-boilerplate/models"
	"github.com/gin-gonic/gin"
)

var identityKey = "id"

func payloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(*models.User); ok {
		return jwt.MapClaims{
			identityKey: v.UserName,
		}
	}
	return jwt.MapClaims{}
}

func identityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return &models.User{
		UserName: claims["id"].(string),
	}
}

func authenticator(c *gin.Context) (interface{}, error) {
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

func authorizator(data interface{}, c *gin.Context) bool {
	if v, ok := data.(*models.User); ok && v.UserName == "admin" {
		return true
	}

	return false
}

func unauthorized(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

func JwtAuthMiddleware() *jwt.GinJWTMiddleware {
	c := config.GetConfig()
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           c.GetString("jwt.realm"),
		Key:             []byte(c.GetString("jwt.secret_key")),
		Timeout:         time.Hour,
		MaxRefresh:      time.Hour,
		IdentityKey:     identityKey,
		PayloadFunc:     payloadFunc,
		IdentityHandler: identityHandler,
		Authenticator:   authenticator,
		Authorizator:    authorizator,
		Unauthorized:    unauthorized,
		TimeFunc:        time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	return authMiddleware
}
