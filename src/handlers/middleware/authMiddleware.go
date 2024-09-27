package middleware

import (
	"DidlyDoodash-api/src/data"
	"DidlyDoodash-api/src/utils"
	"DidlyDoodash-api/src/utils/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract access token
		tokenString := jwt.ExtractToken(c)
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.NotAuthenticated)
			return
		}

		token, err := jwt.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.InvalidToken)
			return
		}

		sub, err := token.Claims.GetSubject()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.NotAuthenticated)
			return
		}
		data.CurrentUser = &sub

		c.Next()
	}
}
