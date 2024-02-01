package movies

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

	response := m.movieService.Get()
	ctx.AbortWithStatusJSON(http.StatusOK, response)
}

func (m movie) SearchMovies(ctx *gin.Context) {
	genre, _ := ctx.GetQuery("genre")
	actor, _ := ctx.GetQuery("actor")
	year, _ := ctx.GetQuery("year")
	Year, _ := strconv.Atoi(year)

	response := m.movieService.FilterByCriteria(genre, actor, Year)
	ctx.AbortWithStatusJSON(http.StatusOK, response)
}

func NewMovieHandler(movieService MovieService.MovieService) MovieHandler {
	return &movie{movieService: movieService}
}
