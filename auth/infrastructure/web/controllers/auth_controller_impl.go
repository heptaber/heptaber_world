package controllers

import (
	"net/http"
	"time"

	"heptaber/auth/app/helper"
	"heptaber/auth/domain/model"
	"heptaber/auth/domain/service"

	"github.com/gin-gonic/gin"

	// validator "github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// var validate = validator.New()

type authController struct {
	authService service.IAuthService
}

func NewAuthController(authService service.IAuthService) *authController {
	return &authController{authService: authService}
}

func (ac *authController) Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var signUpRequestDTO model.SignUpRequestDTO
		if err := c.ShouldBindJSON(&signUpRequestDTO); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		createdUser, err := ac.authService.Signup(signUpRequestDTO)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		c.JSON(http.StatusCreated, createdUser.GetUserDTO())
	}
}

func (ac *authController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginRequestDTO model.LoginRequestDTO
		if err := c.ShouldBindJSON(&loginRequestDTO); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		accessToken, refreshToken, err := ac.authService.Login(loginRequestDTO)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.Header("Authorization", "Bearer "+accessToken)
		c.SetCookie(
			"refresh_token",
			refreshToken,
			int((time.Hour * time.Duration(4320)).Seconds()), // maxAge=180 days more
			"/",
			"",
			false,
			true,
		)

		c.JSON(http.StatusOK, gin.H{"message": "login successfuly"})
	}
}

func (ac *authController) Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.SetCookie("refresh_token", "", -1, "/", "", false, true)
		userId, ok := c.Get("userId")
		if ok {
			if err := ac.authService.Logout(userId.(string)); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": "already logged out"})
			}
		}
		c.JSON(http.StatusOK, gin.H{"message": "logout successful"})
	}
}

func (ac *authController) RefreshToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		refreshToken, err := c.Cookie("refresh_token")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		rClaims, err := helper.ValidateToken(refreshToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid refresh token provided"})
			return
		}

		accessToken, err := ac.authService.GetNewAccessToken(rClaims.Issuer)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.Header("Authorization", "Bearer "+accessToken)
		c.JSON(http.StatusOK, gin.H{"message": "access token refreshed successfuly"})
	}
}

func (ac *authController) VerifyUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		vCode := c.Param("code")
		code, err := uuid.Parse(vCode)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid verification code format"})
			return
		}

		err = ac.authService.VerifyUser(code)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "email verified"})
	}
}
