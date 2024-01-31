package router

import (
	"github.com/gin-gonic/gin"
	"netflixRental/internal/handler/helloworld"
	"netflixRental/internal/handler/movies"
	"netflixRental/internal/service/MovieService"
)

func RegisterRoutes(engine *gin.Engine) {

	movieService := MovieService.NewMovieService()
	var movieHandler movies.MovieHandler
	movieHandler = movies.NewMovieHandler(movieService)

	group := engine.Group("/netflix/api")
	{
		group.GET("/hello", helloworld.HelloWorld)
		group.GET("/movies", movieHandler.ListMovies)
	}
}
