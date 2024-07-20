package router

import (
	"github.com/gin-gonic/gin"
	"github.com/levifleal/socialMedia/backEnd/handlers/auth"
)

func initRoutes(r *gin.Engine) {

	auth.InitAuthHandler()

	basePathV1 := "/api/v1"

	v1 := r.Group(basePathV1)

	authRoute := v1.Group("/Auth")
	{
		authRoute.POST("/SignIn", auth.SignInUserHandler)
		authRoute.POST("/SignUp", auth.SignUpUserHandler)
	}

}
