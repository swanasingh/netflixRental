package movies

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"netflixRental/internal/service/MovieService"
)

type MovieHandler interface {
	ListMovies(ctx *gin.Context)
}

type movie struct {
	movieService MovieService.MovieService
}

func (m movie) ListMovies(ctx *gin.Context) {

	response := m.movieService.Get()
	ctx.AbortWithStatusJSON(http.StatusOK, response)
}

func NewMovieHandler(movieService MovieService.MovieService) MovieHandler {
	return &movie{movieService: movieService}
}
