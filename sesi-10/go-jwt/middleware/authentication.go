package middleware

import (
	"fgd-golang/sesi-10/go-jwt/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		validateToken, err := helpers.ValidateToken(c)
		_ = validateToken

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "unauthorized",
				"message": err.Error(),
			})
			return
		}
		c.Set("userData", validateToken)
		c.Next()
	}
}
