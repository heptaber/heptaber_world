package controllers

import (
	"net/http"
	"strconv"

	"heptaber/auth/app/helper"
	"heptaber/auth/domain/service"

	"heptaber/auth/domain/model"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *userController {
	return &userController{userService: userService}
}

func (uc *userController) GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		recordPerPage, err := strconv.Atoi(c.DefaultQuery("recordPerPage", "10"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid page size param"})
			return
		}
		page, err := strconv.Atoi(c.DefaultQuery("pageNumber", "1"))
		if err != nil || page < 1 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid page number param"})
			return
		}

		users, totalCount, err := uc.userService.GetUsers(recordPerPage, page)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while retrieving users"})
			return
		}

		shortUsersDTO := func(usrs []model.User) []model.ShortUserDTO {
			var res []model.ShortUserDTO = make([]model.ShortUserDTO, len(usrs))
			for i := 0; i < len(usrs); i++ {
				res[i] = *usrs[i].GetShortUserDTO()
			}
			return res
		}(*users)

		c.JSON(http.StatusOK, gin.H{
			"users":      shortUsersDTO,
			"totalCount": totalCount,
		})
	}
}

func (uc *userController) GetUserByEmail() gin.HandlerFunc {
	return func(c *gin.Context) {
		userEmail := c.Param("email")
		foundUser, err := uc.userService.GetUserByEmail(userEmail)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, foundUser.GetUserDTO())
	}
}

func (uc *userController) GetMe() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := helper.ValidateToken(helper.GetBearerTokenValue(c))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user, err := uc.userService.GetUserById(claims.UserId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, user.GetUserDTO())
	}
}

func (uc *userController) SetUserRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		var roleQuery struct {
			UserId string `json:"user_id" validate:"required"`
			Role   string `json:"role" validate:"required"`
		}
		if err := c.ShouldBindJSON(&roleQuery); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request data"})
			return
		}
		updatedUser, err := uc.userService.SetUserRoleByUserId(roleQuery.UserId, roleQuery.Role)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
			return
		}

		c.JSON(http.StatusOK, updatedUser.GetUserDTO())
	}
}

func (uc *userController) SetUserLock() gin.HandlerFunc {
	return func(c *gin.Context) {
		var lockQuery struct {
			UserId   string `json:"user_id" validate:"required"`
			IsLocked bool   `json:"is_locked" validate:"required"`
		}
		if err := c.ShouldBindJSON(&lockQuery); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request data"})
			return
		}
		updatedUser, err := uc.userService.SetUserLockByUserId(lockQuery.UserId, lockQuery.IsLocked)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
			return
		}

		c.JSON(http.StatusOK, updatedUser.GetUserDTO())
	}
}
