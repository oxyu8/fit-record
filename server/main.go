package main

import (
	"fit_record_server/src/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	middleware.SetupRouter(app)
	app.Run(":3001")
}
