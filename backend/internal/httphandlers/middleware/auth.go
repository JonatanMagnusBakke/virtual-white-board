package middleware

import (
	"fmt"
	"net/http"
	"virtual-white-board-service/internal/services"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//AuthMiddleware middleware for auth
func AuthMiddleware(jwtService services.JWTService) func(c *gin.Context) {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		token, err := jwtService.ValidateToken(tokenString)
		if err == nil && token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println(claims)
		} else {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}
