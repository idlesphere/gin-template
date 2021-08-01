package middleware

import (
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		// allow Cross-domain
		c.Header("Access-Control-Allow-Origin", "*")

		// allow headers
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")

		// allow methods
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")

		// c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		// c.Header("Access-Control-Allow-Credentials", "true")

	}
}
