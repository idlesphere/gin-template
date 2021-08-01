package app

import (
	"gin-template/app/middleware"
	"gin-template/app/router"

	"github.com/gin-gonic/gin"
)

func Start(port string) {
	app := gin.New()

	// Global middleware
	// middleware -> Logger will write the logs to gin.DefaultWriter(os.Stdout) by default gin.DefaultWriter
	app.Use(gin.Logger())
	// middleware -> Recovery recovers from any panics and writes a 500 if there was one.
	app.Use(gin.Recovery())
	// middleware -> cors
	app.Use(middleware.Cors())

	// Routers
	router.SetRouter(app)

	app.Run(":" + port)
}
