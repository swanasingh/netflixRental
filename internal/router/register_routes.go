package router

import (
	"github.com/gin-gonic/gin"
	"netflixRental/internal/handler/helloworld"
)

func RegisterRoutes(engine *gin.Engine) {

	group := engine.Group("/netflix/api")
	{
		group.GET("/hello", helloworld.HelloWorld)
	}
}
