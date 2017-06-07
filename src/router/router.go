package router

import (
	sample "api/v1/sample"
	"github.com/gin-gonic/gin"
)

func V1(api *gin.Engine) *gin.Engine {
	v1 := api.Group("/api/v1")
	{
		test := v1.Group("/test")
		test.GET("/sample", sample.Hello)
		test.GET("/world",  sample.World)
	}

	return api
}

func HealthCheck(app *gin.Engine) *gin.Engine {
	app.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "ok"})
	})

	return app
}
