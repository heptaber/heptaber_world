package controllers

import "github.com/gin-gonic/gin"

type IAuthController interface {
	Signup() gin.HandlerFunc
	Login() gin.HandlerFunc
	Logout() gin.HandlerFunc
	RefreshToken() gin.HandlerFunc
	VerifyUser() gin.HandlerFunc
}
