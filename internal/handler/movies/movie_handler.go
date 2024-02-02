package movies

import (
	"github.com/gin-gonic/gin"
	"net/http"
	movie2 "netflixRental/internal/models/movie"
	"netflixRental/internal/service/MovieService"
	"strconv"
)

type MovieHandler interface {
	ListMovies(ctx *gin.Context)
	SearchMovies(ctx *gin.Context)
}

type movie struct {
	movieService MovieService.MovieService
}

func (m movie) ListMovies(ctx *gin.Context) {

	genre, ok1 := ctx.GetQuery("genre")
	actor, ok2 := ctx.GetQuery("actor")
	year, ok3 := ctx.GetQuery("year")
	var criteria movie2.Criteria
	if ok1 || ok2 || ok3 {
		criteria = movie2.Criteria{genre, actor, year}
	}
	response := m.movieService.Get(criteria)
	ctx.JSON(http.StatusOK, response)
}

func (m movie) SearchMovies(ctx *gin.Context) {
	genre, _ := ctx.GetQuery("genre")
	actor, _ := ctx.GetQuery("actor")
	year, _ := ctx.GetQuery("year")
	Year, _ := strconv.Atoi(year)

	response := m.movieService.FilterByCriteria(genre, actor, Year)
	ctx.JSON(http.StatusOK, response)
}

func NewMovieHandler(movieService MovieService.MovieService) MovieHandler {
	return &movie{movieService: movieService}
}
