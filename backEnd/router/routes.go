package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/levifleal/socialMedia/backEnd/handlers/auth"
)

func initRoutes(r *gin.Engine) {

	auth.InitAuthHandler()

	basePathV1 := "/api/v1"

	v1 := r.Group(basePathV1).Use(authMiddleware())
	{
		v1.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"msg": "ok"})
		})
	}

	authRoute := r.Group("/Auth")
	{
		authRoute.POST("/SignIn", auth.SignInUserHandler)
		authRoute.POST("/SignUp", auth.SignUpUserHandler)
	}

}
