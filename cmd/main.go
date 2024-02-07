package main

import (
	"github.com/gin-gonic/gin"
	"netflixRental/configs"
	"netflixRental/database/db"
	"netflixRental/internal/router"
)

func main() {
	engine := gin.Default()
	var config = configs.Config{}
	dbConnect := db.CreateConnection(config)
	db.RunMigration(dbConnect)
	router.RegisterRoutes(engine, dbConnect)

	engine.Run("localhost:8080")
}
