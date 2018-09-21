package middlewares

import (
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/cowboy-coding/go-gin-boilerplate/lib/services"
)

func JwtAuthMiddleware() *jwt.GinJWTMiddleware {
	// later: config := config.GetConfig()
	authService := new(services.AuthService)
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "jwt.realm",
		Key:             []byte("jwt.secret_key"),
		Timeout:         time.Hour,
		MaxRefresh:      time.Hour,
		IdentityKey:     "id",
		PayloadFunc:     authService.PayloadFunc,
		IdentityHandler: authService.IdentityHandler,
		Authenticator:   authService.Authenticator,
		Authorizator:    authService.Authorizator,
		Unauthorized:    authService.Unauthorized,
		TimeFunc:        time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	return authMiddleware
}
