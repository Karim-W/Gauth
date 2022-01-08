package controllers

import (
	"net/http"

	"gauth.com/Dbmanager/models"
	"gauth.com/services"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	AddUser(*gin.Context)
	GetUser(*gin.Context)
}

func NewUserController(s services.UserService) UserController {
	return userController{
		userService: s,
	}
}

type userController struct {
	userService services.UserService
}

func (u userController) AddUser(c *gin.Context) {
	var newlyCreatedUser models.User
	if err := c.ShouldBindJSON(&newlyCreatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u.userService.CreateNewUser(newlyCreatedUser)
}

func (u userController) GetUser(c *gin.Context) {
	user, err := u.userService.GetUser(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
