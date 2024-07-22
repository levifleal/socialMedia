package router

import (
	"github.com/gin-gonic/gin"
	"github.com/levifleal/socialMedia/backEnd/handlers"
	"github.com/levifleal/socialMedia/backEnd/handlers/auth"
)

func initRoutes(r *gin.Engine) {

	//initializing package handlers
	handlers.Init()

	basePathV1 := "/api/v1"

	//default v1 router --Protected
	v1 := r.Group(basePathV1).Use(authMiddleware())
	{
		v1.POST("/RedefinePassword", auth.RedefinePasswordHandler)
	}

	//auth router --Unprotected
	authRoute := r.Group("/Auth")
	{
		authRoute.POST("/SignIn", auth.SignInUserHandler)
		authRoute.POST("/SignUp", auth.SignUpUserHandler)
	}

}
