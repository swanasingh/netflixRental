package MovieService

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"netflixRental/internal/models/movie"
	"netflixRental/internal/repository/movie_repo/mocks"
	"testing"
)

func TestMovieService(t *testing.T) {

	m1 := movie.Movie{
		1,
		"TestTitle",
		"2003",
		"TV-Y",
		"2003-09-30T00:00:00Z",
		"81",
		"Animation, Family",
		"Owen Hurley",
		"Elana Lesser, Cliff Ruby",
		"Kelly Sheridan, Mark Hildreth, Kelsey Grammer, Maggie Wheeler",
		"Barbie comes to life in her third animated movie, based on the beloved fairy tale and set to the brilliant music of Tchaikovsky.",
		"English",
		"USA",
		"2 nominations.",
		"https://m.media-amazon.com/images/M/MV5BNDAzZDBhODAtNmU4My00NWY5LTlmYTYtZDVkOTk3MDcyMDEyXkEyXkFqcGdeQXVyNDE5MTU2MDE@._V1_SX300.jpg",
		0,
		0,
		0,
		"",
		"",
		"",
		"",
		"",
		"",
		false,
	}

	t.Run("validate count of Movies", func(t *testing.T) {

		var mockResponse []movie.Movie
		mockResponse = append(mockResponse, m1)

		mockRepository := mocks.MovieRepository{}
		mockRepository.On("Get", movie.Criteria{}).Return(mockResponse)
		movieService := NewMovieService(&mockRepository)

		got := movieService.Get(movie.Criteria{})
		assert.Equal(t, len(got), 1)
		assert.Equal(t, mockResponse, got)
	})

	t.Run("validate filter by criteria", func(t *testing.T) {
		var mockResponse []movie.Movie
		mockResponse = append(mockResponse, m1)

		mockRepository := mocks.MovieRepository{}
		criteria := movie.Criteria{Actors: "Animation", Genre: "Kelly Sheridan", Year: 2003}
		mockRepository.On("Get", criteria).Return(mockResponse)
		movieService := NewMovieService(&mockRepository)

		got := movieService.Get(criteria)

		assert.Equal(t, mockResponse[0].Year, got[0].Year)
		assert.Equal(t, mockResponse[0].Genre, got[0].Genre)
	})
	t.Run("get movie details if correct id is given", func(t *testing.T) {

		mockRepository := mocks.MovieRepository{}
		mockRepository.On("GetMovieDetails", 1).Return(m1, nil)
		movieService := NewMovieService(&mockRepository)

		got, _ := movieService.GetMovieDetails(1)

		assert.Equal(t, m1.Id, got.Id)
		assert.Equal(t, m1.Title, got.Title)
	})

	t.Run("should return empty movie details if incorrect id is given", func(t *testing.T) {

		mockRepository := mocks.MovieRepository{}
		mockRepository.On("GetMovieDetails", 9).Return(movie.Movie{}, errors.New("Invalid Id"))
		movieService := NewMovieService(&mockRepository)

		_, err := movieService.GetMovieDetails(9)

		assert.Equal(t, "Invalid Id", err.Error())

	})

	t.Run("should add cartItems to cart", func(t *testing.T) {
		mockRepository := mocks.MovieRepository{}
		cartItem := movie.CartItem{MovieId: 1, UserId: 10}
		mockRepository.On("SaveCartData", cartItem).Return(nil)
		movieService := NewMovieService(&mockRepository)
		err := movieService.AddToCart(cartItem)

		assert.NoError(t, err)

	})

	t.Run("should not add invalid cartItems to cart", func(t *testing.T) {
		expectedError := errors.New("Movie is not present in database")
		mockRepository := mocks.MovieRepository{}
		cartItem := movie.CartItem{MovieId: 10, UserId: 10}
		mockRepository.On("SaveCartData", cartItem).Return(expectedError)
		movieService := NewMovieService(&mockRepository)
		err := movieService.AddToCart(cartItem)

		assert.Equal(t, expectedError, err)

	})

	t.Run("should return cartItems for valid user", func(t *testing.T) {

		var mockResponse []movie.Movie
		mockResponse = append(mockResponse, m1)
		mockRepository := mocks.MovieRepository{}
		mockRepository.On("ViewCart", 1).Return(mockResponse)
		movieService := NewMovieService(&mockRepository)
		response := movieService.ViewCart(1)
		assert.Equal(t, mockResponse, response)

	})

	t.Run("should return empty cart for invalid user", func(t *testing.T) {

		var mockResponse []movie.Movie
		mockRepository := mocks.MovieRepository{}
		mockRepository.On("ViewCart", 5).Return(mockResponse)
		movieService := NewMovieService(&mockRepository)
		response := movieService.ViewCart(5)
		assert.Equal(t, mockResponse, response)

	})

}
