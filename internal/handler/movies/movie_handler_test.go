package movies

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	movie2 "netflixRental/internal/models/movie"
	"netflixRental/internal/service/MovieService/mocks"
	"testing"
)

func TestShouldReturnListOfMovieWithMock(t *testing.T) {

	movieService := mocks.MovieService{}
	mockResponse := []movie2.Movie{{
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
	},
	}
	movieService.On(
		"Get",
		movie2.Criteria{},
	).Return(mockResponse)

	handler := NewMovieHandler(&movieService)
	responseRecorder := getResponse(t, handler.ListMovies, "/netflix/api/movies", "/netflix/api/movies", nil)
	var response []movie2.Movie
	err := json.NewDecoder(responseRecorder.Body).Decode(&response)
	require.NoError(t, err)

	assert.Equal(t, responseRecorder.Code, http.StatusOK)
	assert.Equal(t, response, mockResponse)

}

func TestShouldReturnListOfMoviesFilteredByCriteria(t *testing.T) {

	movieService := mocks.MovieService{}
	mockResponse := []movie2.Movie{{
		Title:  "Barbie",
		Year:   "2023",
		ImdbId: "tt1517268",
		Type:   "movie",
		Poster: "https://m.media-amazon.com/images/M/MV5BNjU3N2QxNzYtMjk1NC00MTc4LTk1NTQtMmUxNTljM2I0NDA5XkEyXkFqcGdeQXVyODE5NzE3OTE@._V1_SX300.jpg",
	}, {
		Title:  "Barbie as The Princess and the Pauper",
		Year:   "2023",
		ImdbId: "tt0426955",
		Type:   "movie",
		Poster: "https://m.media-amazon.com/images/M/MV5BMGY5MzU3MzItNDBjMC00YjQzLWEzMTUtMGMxMTEzYjhkMGNkXkEyXkFqcGdeQXVyNDE5MTU2MDE@._V1_SX300.jpg",
	},
	}

	criteria := movie2.Criteria{Actors: "Kelly Sheridan", Genre: "Animation", Year: 2023}
	movieService.On("Get", criteria).Return(mockResponse, nil)
	handler := NewMovieHandler(&movieService)
	responseRecorder := getResponse(t, handler.ListMovies, "/netflix/api/movies", "/netflix/api/movies?actor=Kelly Sheridan&genre=Animation&year=2023", nil)
	var response []movie2.Movie

	err := json.NewDecoder(responseRecorder.Body).Decode(&response)
	require.NoError(t, err)

	assert.Equal(t, responseRecorder.Code, http.StatusOK)
	assert.Equal(t, response, mockResponse)
	assert.Equal(t, mockResponse[0].Year, response[0].Year)

}
func TestShouldReturnMovieDetailsWhenCorrectIdGiven(t *testing.T) {

	movieService := mocks.MovieService{}
	mockResponse := movie2.Movie{
		Id:     5,
		Title:  "Barbie",
		Year:   "2023",
		ImdbId: "tt1517268",
		Type:   "movie",
		Poster: "https://m.media-amazon.com/images/M/MV5BNjU3N2QxNzYtMjk1NC00MTc4LTk1NTQtMmUxNTljM2I0NDA5XkEyXkFqcGdeQXVyODE5NzE3OTE@._V1_SX300.jpg",
	}
	movieService.On("GetMovieDetails", 5).Return(mockResponse, nil)
	handler := NewMovieHandler(&movieService)
	responseRecorder := getResponse(t, handler.GetMovieDetails, "/netflix/api/movies/:id", "/netflix/api/movies/5", nil)
	var response movie2.Movie
	err := json.NewDecoder(responseRecorder.Body).Decode(&response)
	require.NoError(t, err)
	assert.Equal(t, responseRecorder.Code, http.StatusOK)
	assert.Equal(t, response.Id, mockResponse.Id)
	assert.Equal(t, response, mockResponse)
}

func TestShouldReturnEmptyMovieDetailsWhenIncorrectIdGiven(t *testing.T) {

	movieService := mocks.MovieService{}

	movieService.On("GetMovieDetails", 9).Return(movie2.Movie{}, errors.New("Invalid Id"))
	handler := NewMovieHandler(&movieService)
	responseRecorder := getResponse(t, handler.GetMovieDetails, "/netflix/api/movies/:id", "/netflix/api/movies/9", nil)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestShouldAddToCartWhenCartItemIsValid(t *testing.T) {
	movieService := mocks.MovieService{}
	cartRequest := movie2.CartRequest{MovieId: 1, UserId: 2}
	cartItem := movie2.CartItem{MovieId: 1, UserId: 2, Status: true}
	jsonData, err := json.Marshal(cartRequest)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	movieService.On("AddToCart", cartItem).Return(nil)
	handler := NewMovieHandler(&movieService)
	responseRecorder := postResponse(t, handler.AddToCart, "/netflix/api/movies/add_to_cart", "/netflix/api/movies/add_to_cart", jsonData)

	assert.Equal(t, http.StatusCreated, responseRecorder.Code)
}

func TestShouldNotAddToCartWhenCartItemIsInValid(t *testing.T) {
	movieService := mocks.MovieService{}
	cartRequest := movie2.CartRequest{MovieId: 10, UserId: 2}
	cartItem := movie2.CartItem{MovieId: 10, UserId: 2, Status: true}
	expectedError := errors.New("Movie is not present in database")
	jsonData, err := json.Marshal(cartRequest)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	movieService.On("AddToCart", cartItem).Return(expectedError)
	handler := NewMovieHandler(&movieService)
	responseRecorder := postResponse(t, handler.AddToCart, "/netflix/api/movies/add_to_cart", "/netflix/api/movies/add_to_cart", jsonData)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestShouldReturnCartItemsWhenValidUserIdIsGiven(t *testing.T) {
	mockResponse := []movie2.Movie{{
		Title:  "Barbie",
		Year:   "2023",
		ImdbId: "tt1517268",
	}, {
		Title:  "Barbie as The Princess and the Pauper",
		Year:   "2023",
		ImdbId: "tt0426955",
	},
	}

	movieService := mocks.MovieService{}
	cartRequest := movie2.CartRequest{UserId: 2}

	jsonData, err := json.Marshal(cartRequest)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
	movieService.On("ViewCart", cartRequest.UserId).Return(mockResponse)
	handler := NewMovieHandler(&movieService)
	responseRecorder := getResponse(t, handler.ViewCart, "/netflix/api/movies/cart", "/netflix/api/movies/cart", jsonData)

	var response []movie2.Movie

	err = json.NewDecoder(responseRecorder.Body).Decode(&response)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, mockResponse, response)
}

func TestShouldReturnEmptyCartMessageWhenInvalidUserIdIsGiven(t *testing.T) {
	var mockResponse []movie2.Movie

	movieService := mocks.MovieService{}
	cartRequest := movie2.CartRequest{UserId: 12}

	jsonData, err := json.Marshal(cartRequest)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
	movieService.On("ViewCart", cartRequest.UserId).Return(mockResponse)
	handler := NewMovieHandler(&movieService)
	responseRecorder := getResponse(t, handler.ViewCart, "/netflix/api/movies/cart", "/netflix/api/movies/cart", jsonData)

	var response string

	err = json.NewDecoder(responseRecorder.Body).Decode(&response)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, response, "Cart is empty")
}

func getResponse(t *testing.T, handlerFunc gin.HandlerFunc, handlerUrl, url string, body []byte) *httptest.ResponseRecorder {
	engine := gin.Default()
	engine.GET(handlerUrl, handlerFunc)
	request, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(body))
	require.NoError(t, err)
	responseRecorder := httptest.NewRecorder()
	engine.ServeHTTP(responseRecorder, request)
	return responseRecorder
}

func postResponse(t *testing.T, handlerFunc gin.HandlerFunc, handlerUrl, url string, body []byte) *httptest.ResponseRecorder {
	engine := gin.Default()
	engine.POST(handlerUrl, handlerFunc)
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	require.NoError(t, err)
	responseRecorder := httptest.NewRecorder()
	engine.ServeHTTP(responseRecorder, request)
	return responseRecorder
}

//func TestShouldReturnListOfMovie(t *testing.T) {
//
//	movieService := MovieService.NewMovieService()
//	handler := NewMovieHandler(movieService)
//	responseRecorder := getResponse(t, handler)
//	var response []movie2.Movie
//	err := json.NewDecoder(responseRecorder.Body).Decode(&response)
//	require.NoError(t, err)
//
//	assert.Equal(t, responseRecorder.Code, http.StatusOK)
//}
