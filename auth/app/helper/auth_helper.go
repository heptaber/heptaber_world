package helper

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	unauthorizedMessage = "unauthorized to access this resource"
)

func CheckUserRole(c *gin.Context, role string) (err error) {
	claims, err := ValidateToken(GetBearerTokenValue(c))
	userRole := claims.Role
	if err != nil {
		return err
	}
	if string(userRole) != role {
		return fmt.Errorf(unauthorizedMessage)
	}
	return nil
}

func MatchUserRoleToUid(c *gin.Context, userId string) (err error) {
	userRole := c.GetString("role")
	user_id := c.GetString("user_id")

	if userRole == "REGULAR" && userId != user_id {
		return fmt.Errorf(unauthorizedMessage)
	}

	err = CheckUserRole(c, userRole)
	return err
}

func GetBearerTokenValue(c *gin.Context) string {
	return strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
}
