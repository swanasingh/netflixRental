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
	GetMovieDetails(ctx *gin.Context)
	AddToCart(ctx *gin.Context)
}

type movie struct {
	movieService MovieService.MovieService
}

func (m movie) AddToCart(ctx *gin.Context) {
	var cartData movie2.CartRequest
	fmt.Println("I FOUND THE API")
	if err := ctx.ShouldBindJSON(&cartData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if cartData.MovieId == 0 {
		ctx.JSON(http.StatusBadRequest, "Provide valid movie_id")
		return
	}
	cartItem := movie2.CartItem{MovieId: cartData.MovieId, UserId: cartData.UserId, Status: true}
	if err := m.movieService.AddToCart(cartItem); err != nil {
		ctx.JSON(http.StatusBadRequest, "This Movie Is Not Available")
		return
	} else {
		ctx.JSON(http.StatusCreated, cartItem)
	}

}

func (m movie) GetMovieDetails(ctx *gin.Context) {
	id := ctx.Param("id")
	Id, _ := strconv.Atoi(id)
	response, err := m.movieService.GetMovieDetails(Id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, response)
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
