package movies

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	movie2 "netflixRental/internal/models/movie"
	"netflixRental/internal/service/MovieService"
	"strconv"
)

type MovieHandler interface {
	ListMovies(ctx *gin.Context)
}

type movie struct {
	movieService MovieService.MovieService
}

func (m movie) ListMovies(ctx *gin.Context) {

	genre, ok1 := ctx.GetQuery("genre")
	actor, ok2 := ctx.GetQuery("actor")
	year, ok3 := ctx.GetQuery("year")
	Year, _ := strconv.Atoi(year)
	var criteria movie2.Criteria
	if ok1 || ok2 || ok3 {
		criteria = movie2.Criteria{Genre: genre, Actors: actor, Year: Year}
	}
	fmt.Println("criteria")
	fmt.Println(criteria)
	response := m.movieService.Get(criteria)
	ctx.JSON(http.StatusOK, response)
}

func NewMovieHandler(movieService MovieService.MovieService) MovieHandler {
	return &movie{movieService: movieService}
}
