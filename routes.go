package main

import (
	"log"

	controllers "gauth.com/Controllers"
	repos "gauth.com/Repos"
	"gauth.com/dbmanager"
	"gauth.com/services"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(db dbmanager.DBContext) {
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
	users.POST("/", userController.AddUser)

	router.Run(":9022")

}