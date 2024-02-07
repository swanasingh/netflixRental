package router

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"netflixRental/internal/handler/helloworld"
	"netflixRental/internal/handler/movies"
	"netflixRental/internal/repository/movie_repo"
	"netflixRental/internal/service/MovieService"
)

func RegisterRoutes(engine *gin.Engine, db *sql.DB) {

	movieRepository := movie_repo.NewMovieRepository(db)
	movieService := MovieService.NewMovieService(movieRepository)
	var movieHandler movies.MovieHandler
	movieHandler = movies.NewMovieHandler(movieService)

	group := engine.Group("/netflix/api")
	{
		group.GET("/hello", helloworld.HelloWorld)
		group.GET("/movies", movieHandler.ListMovies)
		group.GET("/movies/:id", movieHandler.GetMovieDetails)
		group.POST("/movies/add_to_cart", movieHandler.AddToCart)
	}

}
