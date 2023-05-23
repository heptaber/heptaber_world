package helper

import (
	"fmt"
	"os"
	"time"

	"heptaber/auth/domain/model"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JwtClaims struct {
	UserId   string
	Email    string
	Username string
	Role     model.UserRole
	jwt.RegisteredClaims
}

var SECRET_KEY string = os.Getenv("SECRET_KEY")

// TODO: impl get tokens TTLs from env
// var ACCESS_TOKEN_TTL = os.Getenv("ACCESS_TOKEN_TTL")
// var REFRESH_TOKEN_TTL = os.Getenv("REFRESH_TOKEN_TTL")

func GenerateAccessToken(user model.User, currentTime time.Time) (string, error) {
	claims := &JwtClaims{
		UserId:   user.ID.String(),
		Email:    user.Email,
		Username: user.Username,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    user.ID.String(),
			IssuedAt:  jwt.NewNumericDate(currentTime),
			ExpiresAt: jwt.NewNumericDate(currentTime.Add(time.Hour * time.Duration(72))), // 3 days
		},
	}
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	return accessToken, err
}

func GenerateRefreshToken(userId uuid.UUID, currentTime time.Time) (string, error) {
	refreshClaims := &JwtClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    userId.String(),
			IssuedAt:  jwt.NewNumericDate(currentTime),
			ExpiresAt: jwt.NewNumericDate(currentTime.Add(time.Hour * time.Duration(4320))), // 180 days
		},
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	return refreshToken, err
}

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
