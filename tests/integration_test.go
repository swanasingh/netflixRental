package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"log"
	"net/http"
	"net/http/httptest"
	"netflixRental/database/db"
	movie2 "netflixRental/internal/models/movie"
	"netflixRental/internal/router"
	"os"
	"testing"
)

var engine *gin.Engine

func TestMain(m *testing.M) {
	test_container, ctx, dbConnect, _ := CreatePostgresTestContainer()
	db.RunMigration(dbConnect)
	engine = gin.Default()
	router.RegisterRoutes(engine, dbConnect)
	code := m.Run()

	if err := test_container.Terminate(ctx); err != nil {
		log.Fatal(err)
	}
	os.Exit(code)
}

func TestListAllMovies(t *testing.T) {

	t.Run("TestGetAllMoviesWhenMoviesPresentInDB", func(t *testing.T) {

		request, err := http.NewRequest(http.MethodGet, "/netflix/api/movies", nil)
		require.NoError(t, err)
		responseRecorder := httptest.NewRecorder()
		engine.ServeHTTP(responseRecorder, request)

		var response []movie2.Movie
		err = json.NewDecoder(responseRecorder.Body).Decode(&response)
		require.NoError(t, err)

		// Assert
		assert.Equal(t, responseRecorder.Code, http.StatusOK)
		assert.Equal(t, 7, len(response))
	})

	t.Run("TestGetMovieByIdWhenIsIdIsCorrect", func(t *testing.T) {

		request, err := http.NewRequest(http.MethodGet, "/netflix/api/movies/1", nil)
		require.NoError(t, err)
		responseRecorder := httptest.NewRecorder()
		engine.ServeHTTP(responseRecorder, request)

		var response movie2.Movie
		err = json.NewDecoder(responseRecorder.Body).Decode(&response)
		require.NoError(t, err)

		// Assert

		assert.Equal(t, http.StatusOK, responseRecorder.Code)
		assert.Equal(t, 1, response.Id)
	})

	t.Run("TestShouldNotReturnMovieByIdWhenIsIdIsIncorrect", func(t *testing.T) {

		request, err := http.NewRequest(http.MethodGet, "/netflix/api/movies/11", nil)
		require.NoError(t, err)
		responseRecorder := httptest.NewRecorder()
		engine.ServeHTTP(responseRecorder, request)
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
		request, err := http.NewRequest(http.MethodPost, "/netflix/api/movies/add_to_cart", bytes.NewBuffer(jsonData))
		require.NoError(t, err)
		responseRecorder := httptest.NewRecorder()
		engine.ServeHTTP(responseRecorder, request)
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
		request, err := http.NewRequest(http.MethodPost, "/netflix/api/movies/add_to_cart", bytes.NewBuffer(jsonData))
		require.NoError(t, err)
		responseRecorder := httptest.NewRecorder()
		engine.ServeHTTP(responseRecorder, request)
		// Assert
		assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	})

}
