package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/levifleal/socialMedia/backEnd/schemas"
	"golang.org/x/crypto/bcrypt"
)

func SignInUserHandler(ctx *gin.Context) {
	r := LoginUserRequest{}

	ctx.BindJSON(&r)

	//validating user input
	if err := r.Validate(); err != nil {
		logger.Errorf("Validation Error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	//finding user in database
	user := schemas.User{}
	if err := db.First(&user, schemas.User{Email: r.Email}).Error; err != nil {
		sendError(ctx, http.StatusUnauthorized, "invalid credentials")
		return
	}

	//checking if passwords match
	passMatch := CheckPasswordHash(r.Password, user.PasswordHash)
	if !passMatch {
		sendError(ctx, http.StatusUnauthorized, "invalid credentials")
		return
	}

	//generating token
	token, err := newJwt(&user)
	if err != nil {
		logger.Errorf("error generating token: %v", err.Error())
		return
	}

	sendSuccess(ctx, token, schemas.UserRespose{ Email: user.Email, Id: user.Id, BaseSchema: schemas.BaseSchema{
		CreatedAt: user.CreatedAt, UpdatedAt: user.UpdatedAt,
	}})
}

// checking if passwords match thru bcrypt
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
