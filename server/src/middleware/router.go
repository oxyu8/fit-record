package middleware

import "github.com/gin-gonic/gin"

func SetupRouter(app *gin.Engine) {
	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "It works! ",
		})
	})
	app.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
}
