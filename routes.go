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
	userService := services.NewUserService(userRepository)

	userController := controllers.NewUserController(userService)

	users := router.Group("users")

	users.GET("/:id", userController.GetUser)
	users.GET("/", userController.GetAllUsers)
	users.POST("/", userController.AddUser)
	users.PATCH("/:id", userController.UpdateUser)
	users.DELETE("/:id", userController.DeleteUser)
	router.Run(":9022")

}
