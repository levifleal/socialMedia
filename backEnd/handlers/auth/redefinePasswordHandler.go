package auth

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/levifleal/socialMedia/backEnd/schemas"
)

func RedefinePasswordHandler(ctx *gin.Context) {

	r := RedefinePasswordRequest{}

	ctx.BindJSON(&r)

	//validating user input
	if err := r.Validate(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// separating "Bearer " from token
	bearer := ctx.Request.Header.Get("Authorization")
	token := strings.Split(bearer, " ")

	//decoding token
	data := MyUserClaim{}
	if err := decodeToken(token[1], &data); err != nil {
		logger.Errorf("error decoding token: %v", err)
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error decoding token: %v", err))
		return
	}

	//finding user in database
	user := schemas.User{}
	if err := db.First(&user, schemas.User{Id: data.Id}).Error; err != nil {
		sendError(ctx, http.StatusUnauthorized, "invalid credentials")
		return
	}

	//encrypting password
	hashPass, err := hashPassword(r.Password)
	if err != nil {
		logger.Error("error crypting password")
		sendError(ctx, http.StatusInternalServerError, "error crypting password")
		return
	}

	user.PasswordHash = hashPass

	//saving new password in database
	if err := db.Save(&user).Error; err != nil {
		logger.Errorf("error updating content: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating content")
		return
	}

	sendOk(ctx)

}

func decodeToken(token string, claim *MyUserClaim) error {
	_, err := jwt.ParseWithClaims(token, claim, func(token *jwt.Token) (interface{}, error) {
		if os.Getenv("JWT_SECRET") == "" {
			return "", fmt.Errorf("jwt secret is empty")
		}
		key := []byte(os.Getenv("JWT_SECRET"))
		return key, nil
	})
	if err != nil {
		return nil
	}

	return nil
}
