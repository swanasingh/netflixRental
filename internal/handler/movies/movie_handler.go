package movies

import (
	"fmt"
	"net/http"
	"netflixRental/internal/helpers"
	movie2 "netflixRental/internal/models/movie"
	"netflixRental/internal/service/EmailService"
	"netflixRental/internal/service/MovieService"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MovieHandler interface {
	ListMovies(ctx *gin.Context)
	GetMovieDetails(ctx *gin.Context)
	AddToCart(ctx *gin.Context)
	ViewCart(ctx *gin.Context)
	CreateOrder(ctx *gin.Context)
	SendInvoice(ctx *gin.Context)
}

type movie struct {
	movieService MovieService.MovieService
	emailService EmailService.EmailService
}

func (m movie) SendInvoice(ctx *gin.Context) {

	var order struct {
		OrderId int `json:"order_id"`
	}
	if err := ctx.BindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if invoices, user, err := m.movieService.GetInvoice(order.OrderId); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	} else {
		resp := movie2.Invoices{user, invoices}
		ctx.JSON(http.StatusOK, resp)
		emailBody := helpers.GenerateInvoiceEmailBody(resp)
		m.emailService.SendInvoice(user.Email, emailBody)

	}
}

func (m movie) CreateOrder(ctx *gin.Context) {

	var order movie2.OrderPayload
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := m.movieService.CreateOrder(order); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	} else {
		ctx.JSON(http.StatusCreated, order)
	}
}

func (m movie) ViewCart(ctx *gin.Context) {
	var userInfo movie2.CartRequest
	if err := ctx.ShouldBindJSON(&userInfo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if userInfo.UserId == 0 {
		ctx.JSON(http.StatusBadRequest, "Provide valid user_id")
		return
	}
	cartItems := m.movieService.ViewCart(userInfo.UserId)
	if cartItems != nil {
		fmt.Println(cartItems)
		ctx.JSON(http.StatusOK, cartItems)
	} else {
		ctx.JSON(http.StatusOK, "Cart is empty")
	}

}

func (m movie) AddToCart(ctx *gin.Context) {
	var cartData movie2.CartRequest
	if err := ctx.ShouldBindJSON(&cartData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if cartData.MovieId == 0 {
		ctx.JSON(http.StatusBadRequest, "Provide valid movie_id")
		return
	}
	cartItem := movie2.CartItem{MovieId: cartData.MovieId, UserId: cartData.UserId, Status: true, Quantity: cartData.Quantity}
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

func NewMovieHandler(movieService MovieService.MovieService, emailService EmailService.EmailService) MovieHandler {
	return &movie{movieService: movieService, emailService: emailService}
}
