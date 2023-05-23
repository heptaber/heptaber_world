package helper

import (
	"fmt"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type UserRole string

const (
	ADMIN     UserRole = "ADMIN"
	MODERATOR UserRole = "MODERATOR"
	REGULAR   UserRole = "REGULAR"
)

type JwtClaims struct {
	UserId   string
	Email    string
	Username string
	Role     UserRole
	jwt.RegisteredClaims
}

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func ValidateToken(jwtToken string) (*JwtClaims, error) {
	claims, err := GetClaimsFromToken(jwtToken)
	if err != nil {
		return nil, err
	}

	if claims.ExpiresAt.Unix() < time.Now().UTC().Unix() {
		return nil, fmt.Errorf("the token is expired")
	}

	return claims, nil
}

func GetClaimsFromToken(jwtToken string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(
		jwtToken,
		&JwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JwtClaims)
	if !ok {
		return nil, fmt.Errorf("the token is invalid")
	}
	return claims, nil
}
