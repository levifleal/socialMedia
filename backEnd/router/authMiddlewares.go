package router

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func authMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//verifing if has Autorization key
		bearerToken := ctx.Request.Header.Get("Authorization")
		token, _ := strings.CutPrefix(bearerToken, "Bearer ")
		if token == "" {
			ctx.Header("content-type", "application/json")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "token is missing",
			})
			return
		}
		//verify if the token is valid
		validator := jwt.NewParser(jwt.WithExpirationRequired())
		_, err := validator.ParseWithClaims(token, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			key := []byte(os.Getenv("JWT_SECRET"))
			return key, nil
		})
		if err != nil {
			logger.Errorf("error validating token: %v", err)
			ctx.Header("content-type", "application/json")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": fmt.Sprint(err.Error()),
			})
			return
		}

		ctx.Next()
	}
}
