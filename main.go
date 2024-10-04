package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sanda-bunescu/ExploreRO/controllers"
	"github.com/sanda-bunescu/ExploreRO/initializers"
)

func init() {
	initializers.LoadEnvFiles()
	initializers.FirebaseInitialization()
	initializers.ConnectToDB()
	initializers.MigrateDB()
}

func main() {
	r := gin.Default()

	userController, err := controllers.NewUserController()
	if err != nil {
		log.Fatalf("Failed to initialize UserController: %v", err)
	}

	r.POST("/CreateUser", userController.CreateUser)
	r.POST("/DeleteUser", userController.DeleteUser)
	
	r.Run()
}
