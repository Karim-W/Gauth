package routes

import (
	"log"

	repos "gauth.com/Repos"
	"gauth.com/dbmanager"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(db dbmanager.DBContext) {
	router := gin.Default()
	//pass DB to repos
	userRepository := repos.NewUserRepository(db.Ctx)
	if err := userRepository.Migrate(); err != nil {
		log.Fatal("User migrate err", err)
	}
	router.Run("9022")

}
