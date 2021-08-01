package router

import (
"gin-template/app/controller"

"github.com/gin-gonic/gin"
"gin-template/app/middleware"
)

func SetRouter(app *gin.Engine) {

	app.NoRoute(func(c *gin.Context) {
		c.String(404, "404 not found")
	})

	app.GET("/ping", new(controller.HealthController).Run)
	app.POST("/login", new(controller.UserController).Login)

	// Use custom middleware in the "v1" group.
	v1 := app.Group("/v1")
	v1.Use(middleware.Auth())
	{
		v1.GET("/hello", new(controller.HealthController).Run)
	}

}

