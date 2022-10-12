package helpers

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ValidateToken(c *gin.Context) (interface{}, error) {
	errResponse := errors.New("Sign In to Processed")
	headerToken := c.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer ")

	if !bearer {
		return nil, errResponse
	}

	stringToken := strings.Split(headerToken, "")[1]

	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}
		return []byte(secretKey), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errResponse
	}

	return token.Claims.(jwt.MapClaims), nil
}
