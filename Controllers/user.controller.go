package controllers

import (
	models "gauth.com/Database/Models"
	"net/http"

	"gauth.com/services"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	AddUser(*gin.Context)
	GetUser(*gin.Context)
	GetAllUsers(*gin.Context)
	UpdateUser(*gin.Context)
	DeleteUser(ctx *gin.Context)
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

func (u userController) GetAllUsers(c *gin.Context) {
	users, err := u.userService.FetchAllUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (u userController) UpdateUser(c *gin.Context) {
	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u.userService.UpdateUser(c.Param("id"), updatedUser)
}

func (u userController) DeleteUser(c *gin.Context) {
	err := u.userService.DeleteUser(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.String(200, "User Deleted")
}
