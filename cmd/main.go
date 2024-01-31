package main

import (
	"github.com/gin-gonic/gin"
	"netflixRental/internal/router"
)

func main() {
	engine := gin.Default()
	router.RegisterRoutes(engine)

	engine.Run("localhost:8080")

}
