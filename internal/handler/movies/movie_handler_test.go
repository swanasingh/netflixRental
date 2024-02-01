package movies

import (
	"encoding/json"
	//"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	movie2 "netflixRental/internal/models/movie"
	"netflixRental/internal/service/MovieService/mocks"
	//mockRepo "netflixRental/internal/repository/movie_repo/mocks"
	"testing"
)

func TestShouldReturnListOfMovieWithMock(t *testing.T) {

	movieService := mocks.MovieService{}
	mockResponse := []movie2.Movie{{
		Title:  "Barbie",
		Year:   "2023",
		ImdbId: "tt1517268",
		Type:   "movie",
		Poster: "https://m.media-amazon.com/images/M/MV5BNjU3N2QxNzYtMjk1NC00MTc4LTk1NTQtMmUxNTljM2I0NDA5XkEyXkFqcGdeQXVyODE5NzE3OTE@._V1_SX300.jpg",
	}, {
		Title:  "Barbie as The Princess and the Pauper",
		Year:   "2004",
		ImdbId: "tt0426955",
		Type:   "movie",
		Poster: "https://m.media-amazon.com/images/M/MV5BMGY5MzU3MzItNDBjMC00YjQzLWEzMTUtMGMxMTEzYjhkMGNkXkEyXkFqcGdeQXVyNDE5MTU2MDE@._V1_SX300.jpg",
	},
	}
	movieService.On("Get").Return(mockResponse, nil)
	handler := NewMovieHandler(&movieService)
	responseRecorder := getResponse(t, handler.ListMovies, "/netflix/api/movies/")
	var response []movie2.Movie
	err := json.NewDecoder(responseRecorder.Body).Decode(&response)
	require.NoError(t, err)

	assert.Equal(t, responseRecorder.Code, http.StatusOK)
	assert.Equal(t, response, mockResponse)

}

//
//func TestShouldReturnListOfMoviesFilteredByCriteria(t *testing.T) {
//
//	movieService := mocks.MovieService{}
//	mockResponse := []movie2.Movie{{
//		Title:  "Barbie",
//		Year:   "2023",
//		ImdbId: "tt1517268",
//		Type:   "movie",
//		Poster: "https://m.media-amazon.com/images/M/MV5BNjU3N2QxNzYtMjk1NC00MTc4LTk1NTQtMmUxNTljM2I0NDA5XkEyXkFqcGdeQXVyODE5NzE3OTE@._V1_SX300.jpg",
//	}, {
//		Title:  "Barbie as The Princess and the Pauper",
//		Year:   "2004",
//		ImdbId: "tt0426955",
//		Type:   "movie",
//		Poster: "https://m.media-amazon.com/images/M/MV5BMGY5MzU3MzItNDBjMC00YjQzLWEzMTUtMGMxMTEzYjhkMGNkXkEyXkFqcGdeQXVyNDE5MTU2MDE@._V1_SX300.jpg",
//	},
//	}
//	movieService.On("FilterByCriteria").Return(mockResponse, nil)
//	handler := NewMovieHandler(&movieService)
//	responseRecorder := getResponse(t, handler.SearchMovies, "/netflix/api/movies/search")
//	var response []movie2.Movie
//
//	fmt.Println("hereeee")
//	fmt.Println(responseRecorder.Body, responseRecorder.Code)
//
//	err := json.NewDecoder(responseRecorder.Body).Decode(&response)
//	require.NoError(t, err)
//
//	assert.Equal(t, responseRecorder.Code, http.StatusOK)
//	assert.Equal(t, response, mockResponse)
//
//}

func getResponse(t *testing.T, handlerFunc gin.HandlerFunc, url string) *httptest.ResponseRecorder {
	engine := gin.Default()
	engine.GET(url, handlerFunc)
	request, err := http.NewRequest(http.MethodGet, "/netflix/api/movies", nil)
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
