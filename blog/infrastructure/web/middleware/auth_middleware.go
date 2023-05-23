package middleware

import (
	"heptaber/blog/app/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := helper.GetBearerTokenValue(c)
		if clientToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No Authorization header provided"})
			c.Abort()
			return
		}

		claims, err := helper.ValidateToken(clientToken)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)
		c.Set("userId", claims.UserId)
		c.Next()
	}
}

func HasAdminPermission() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helper.CheckUserRole(c, "ADMIN"); err != nil {
			c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Insufficient privileges"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func HasModeratorPermission() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := helper.CheckUserRole(c, "MODERATOR"); err != nil {
			if err2 := helper.CheckUserRole(c, "ADMIN"); err2 != nil {
				c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Insufficient privileges"})
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
