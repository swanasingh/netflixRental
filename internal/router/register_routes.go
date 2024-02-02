package router

import (
	"github.com/gin-gonic/gin"
	"netflixRental/configs"
	"netflixRental/database/db"
	"netflixRental/internal/handler/helloworld"
	"netflixRental/internal/handler/movies"
	"netflixRental/internal/repository/movie_repo"
	"netflixRental/internal/service/MovieService"
)

func RegisterRoutes(engine *gin.Engine) {
	var config = configs.Config{}
	dbConnect := db.CreateConnection(config)
	movieRepository := movie_repo.NewMovieRepository(dbConnect)
	movieService := MovieService.NewMovieService(movieRepository)
	var movieHandler movies.MovieHandler
	movieHandler = movies.NewMovieHandler(movieService)

	group := engine.Group("/netflix/api")
	{
		group.GET("/hello", helloworld.HelloWorld)
		group.GET("/movies", movieHandler.ListMovies)
		//group.GET("/movies/search", movieHandler.SearchMovies)
	}
}
