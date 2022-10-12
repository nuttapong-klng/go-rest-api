package authentication

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer"
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenString := authHeader[len(BEARER_SCHEMA)+1:]
		token, err := JWTAuthService().ValidateToken(tokenString)
		if err != nil || !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
