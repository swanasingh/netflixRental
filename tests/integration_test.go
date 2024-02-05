package tests

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"netflixRental/configs"
	"netflixRental/database/db"
	"netflixRental/internal/handler/movies"
	movie2 "netflixRental/internal/models/movie"
	"netflixRental/internal/repository/movie_repo"
	"netflixRental/internal/service/MovieService"
	"testing"
)

func TestAllEndpoints(t *testing.T) {
	//Arrange
	var config = configs.Config{}
	dbConnect := db.CreateConnection(config)
	movieRepository := movie_repo.NewMovieRepository(dbConnect)
	movieService := MovieService.NewMovieService(movieRepository)
	var movieHandler movies.MovieHandler
	movieHandler = movies.NewMovieHandler(movieService)

	t.Run("TestGetAllMoviesWhenMoviesPresentInDB", func(t *testing.T) {
		//Act
		responseRecorder := getResponse(t, movieHandler.ListMovies, "/netflix/api/movies", "/netflix/api/movies")
		var response []movie2.Movie
		err := json.NewDecoder(responseRecorder.Body).Decode(&response)
		require.NoError(t, err)

		// Assert
		assert.Equal(t, responseRecorder.Code, http.StatusOK)
		assert.Equal(t, 7, len(response))
	})

	t.Run("TestGetMovieByIdWhenIsIdIsCorrect", func(t *testing.T) {
		responseRecorder := getResponse(t, movieHandler.GetMovieDetails, "/netflix/api/movies/:id", "/netflix/api/movies/1")
		var response movie2.Movie
		err := json.NewDecoder(responseRecorder.Body).Decode(&response)
		require.NoError(t, err)

		// Assert
		fmt.Println(response)
		assert.Equal(t, http.StatusOK, responseRecorder.Code)
		assert.Equal(t, 1, response.Id)
	})

	t.Run("TestShouldNotReturnMovieByIdWhenIsIdIsIncorrect", func(t *testing.T) {

		responseRecorder := getResponse(t, movieHandler.GetMovieDetails, "/netflix/api/movies/:id", "/netflix/api/movies/11")
		// Assert
		assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	})

	t.Run("TestShouldAddToCartWhenCartItemIsValid", func(t *testing.T) {
		cartRequest := movie2.CartRequest{MovieId: 1, UserId: 2}

		jsonData, err := json.Marshal(cartRequest)
		if err != nil {
			fmt.Println("Error encoding JSON:", err)
			return
		}
		responseRecorder := postResponse(t, movieHandler.AddToCart, "/netflix/api/movies/add_to_cart", "/netflix/api/movies/add_to_cart", jsonData)
		// Assert
		assert.Equal(t, http.StatusCreated, responseRecorder.Code)
	})

	t.Run("TestShouldNotAddToCartWhenCartItemIsInvalid", func(t *testing.T) {
		cartRequest := movie2.CartRequest{MovieId: 11, UserId: 2}

		jsonData, err := json.Marshal(cartRequest)
		if err != nil {
			fmt.Println("Error encoding JSON:", err)
			return
		}
		responseRecorder := postResponse(t, movieHandler.AddToCart, "/netflix/api/movies/add_to_cart", "/netflix/api/movies/add_to_cart", jsonData)
		// Assert
		assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	})
}

func getResponse(t *testing.T, handlerFunc gin.HandlerFunc, handlerUrl, url string) *httptest.ResponseRecorder {

	engine := gin.Default()
	engine.GET(handlerUrl, handlerFunc)
	request, err := http.NewRequest(http.MethodGet, url, nil)
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
