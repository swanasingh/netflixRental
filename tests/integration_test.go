package tests

import (
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

func TestGetAllMoviesWhenMoviesPresentInDB(t *testing.T) {
	//Arrange
	engine := gin.Default()
	var config = configs.Config{}
	dbConnect := db.CreateConnection(config)
	movieRepository := movie_repo.NewMovieRepository(dbConnect)
	movieService := MovieService.NewMovieService(movieRepository)
	var movieHandler movies.MovieHandler
	movieHandler = movies.NewMovieHandler(movieService)

	//Act
	engine.GET("/netflix/api/movies", movieHandler.ListMovies)
	request, err := http.NewRequest(http.MethodGet, "/netflix/api/movies", nil)
	require.NoError(t, err)
	responseRecorder := httptest.NewRecorder()
	engine.ServeHTTP(responseRecorder, request)
	var response []movie2.Movie

	err = json.NewDecoder(responseRecorder.Body).Decode(&response)
	require.NoError(t, err)

	// Assert
	fmt.Println(response)
	assert.Equal(t, responseRecorder.Code, http.StatusOK)
	assert.Equal(t, 7, len(response))
}
