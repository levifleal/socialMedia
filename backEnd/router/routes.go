package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initRoutes(r *gin.Engine) {

	basePathV1 := "/api/v1"

	v1 := r.Group(basePathV1)
	{
		v1.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "Ok",
			})
		})
	}

	// auth := r.Group("/Auth")

}
