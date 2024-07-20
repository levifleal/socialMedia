package auth

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/levifleal/socialMedia/backEnd/schemas"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"golang.org/x/crypto/bcrypt"
)

type MyUserClaim struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func SignUpUserHandler(ctx *gin.Context) {
	request := CreateUserRequest{}

	ctx.BindJSON(&request)

	id, err := gonanoid.New()
	if err != nil {
		logger.Errorf("error creating id: %s", err)
		sendError(ctx, http.StatusInternalServerError, "error creating id")
	}

	err = request.Validate()
	if err != nil {
		logger.Errorf("validation error: %v", err)
		sendError(ctx, http.StatusBadRequest, fmt.Sprintf("validation error: %v", err))
		return
	}

	// encriping password
	hashPass, err := hashPassword(request.Password)
	if err != nil {
		logger.Error("error crypting password")
		sendError(ctx, http.StatusInternalServerError, "error crypting password")
		return
	}

	user := schemas.User{
		Id:           id,
		Name:         request.Name,
		Email:        request.Email,
		PasswordHash: hashPass,
	}

	// creating user in database
	if err := db.Create(&user).Error; err != nil {
		logger.Errorf("error creating content: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating content on database")
		return
	}

	//generating token
	token, err := newJwt(&user)
	if err != nil {
		logger.Errorf("error generating token: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error generating token")
		return
	}

	sendSuccess(ctx, token, schemas.UserRespose{Name: user.Name, Email: user.Email, Id: user.Id, CreatedAt: user.CreatedAt, UpdatedAt: user.UpdatedAt})
}

func hashPassword(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(bytes), err
}

func newJwt(data *schemas.User) (string, error) {
	key := []byte(os.Getenv("JWT_SECRET"))

	claims := MyUserClaim{
		Id:    data.Id,
		Name:  data.Name,
		Email: data.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
		},
	}

	//generating token
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//signing token
	token, err := t.SignedString(key)
	if err != nil {
		return "", err
	}

	return token, nil

}
