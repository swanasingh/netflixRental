package tests

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
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

func TestGetMovieByIdWhenIsIdIsCorrect(t *testing.T) {
	//Arrange
	engine := gin.Default()
	var config = configs.Config{}
	dbConnect := db.CreateConnection(config)
	movieRepository := movie_repo.NewMovieRepository(dbConnect)
	movieService := MovieService.NewMovieService(movieRepository)
	var movieHandler movies.MovieHandler
	movieHandler = movies.NewMovieHandler(movieService)

	//Act
	engine.GET("/netflix/api/movies/:id", movieHandler.GetMovieDetails)
	request, err := http.NewRequest(http.MethodGet, "/netflix/api/movies/1", nil)
	require.NoError(t, err)
	responseRecorder := httptest.NewRecorder()
	engine.ServeHTTP(responseRecorder, request)
	var response movie2.Movie

	err = json.NewDecoder(responseRecorder.Body).Decode(&response)
	require.NoError(t, err)

	// Assert
	fmt.Println(response)
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, 1, response.Id)
}

func TestShouldNotReturnMovieByIdWhenIsIdIsIncorrect(t *testing.T) {
	//Arrange
	engine := gin.Default()
	var config = configs.Config{}
	dbConnect := db.CreateConnection(config)
	movieRepository := movie_repo.NewMovieRepository(dbConnect)
	movieService := MovieService.NewMovieService(movieRepository)
	var movieHandler movies.MovieHandler
	movieHandler = movies.NewMovieHandler(movieService)

	//Act
	engine.GET("/netflix/api/movies/:id", movieHandler.GetMovieDetails)
	request, _ := http.NewRequest(http.MethodGet, "/netflix/api/movies/11", nil)

	responseRecorder := httptest.NewRecorder()
	engine.ServeHTTP(responseRecorder, request)

	// Assert
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestPostgreSQLIntegration(t *testing.T) {
	ctx := context.Background()

	// Define a PostgreSQL container
	req := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5434/tcp"},
		WaitingFor:   wait.ForLog("database system is ready to accept connections"),
	}

	// Create the container
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		fmt.Println("container not created")
		t.Fatal(err)
	}
	defer container.Terminate(ctx)

	// Get the PostgreSQL port
	host, _ := container.Host(ctx)
	fmt.Println(host)
	/*dsn := "user=postgres password=postgres dbname=postgres sslmode=disable host=" + host + " port=" + host

	// Connect to the database
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// Your database test code here

	// Clean up
	if err := container.Terminate(ctx); err != nil {
		t.Fatal(err)
	}*/
}
