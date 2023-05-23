package controllers

import "github.com/gin-gonic/gin"

type IUserController interface {
	GetUsers() gin.HandlerFunc
	GetUserByEmail() gin.HandlerFunc
	GetMe() gin.HandlerFunc
	SetUserRole() gin.HandlerFunc
	SetUserLock() gin.HandlerFunc
}
