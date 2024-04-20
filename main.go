package main

import (
	"os"

	"github.com/Nikittansk/go-auth/controllers"
	"github.com/Nikittansk/go-auth/database"
	"github.com/Nikittansk/go-auth/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Database
	database.Connect(os.Getenv("DATABASE_URL"))
	// Initialize Router
	router := initRouter()
	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}

	return router
}
