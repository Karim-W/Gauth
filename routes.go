package main

import (
	"log"

	controllers "gauth.com/Controllers"
	"gauth.com/Database"
	repos "gauth.com/Repos"
	"gauth.com/services"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(db Database.DBContext) {
	router := gin.Default()
	//pass DB to repos
	userRepository := repos.NewUserRepository(db.Ctx)
	if err := userRepository.Migrate(); err != nil {
		log.Fatal("User migrate err", err)
	}
	u, err := userRepository.GetUser("2fce1c6a-444d-42f6-9ae2-d614d71c5e07")
	if err != nil {
		log.Fatal("Get user err", err)
	}
	log.Print("User", u)
	userService := services.NewUserService(userRepository)

	userController := controllers.NewUserController(userService)

	users := router.Group("users")

	users.GET("/:id", userController.GetUser)
	users.GET("/", userController.GetAllUsers)
	users.POST("/", userController.AddUser)
	users.PATCH("/:id", userController.UpdateUser)
	users.DELETE("/:id", userController.DeleteUser)
	router.Run(":8022")

}
