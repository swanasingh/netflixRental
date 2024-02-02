package movies

import (
	"encoding/json"
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
	responseRecorder := getResponse(t, handler.ListMovies, "/netflix/api/movies", "/netflix/api/movies")
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
	responseRecorder := getResponse(t, handler.ListMovies, "/netflix/api/movies", "/netflix/api/movies?actor=Kelly Sheridan&genre=Animation&year=2023")
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
	responseRecorder := getResponse(t, handler.GetMovieDetails, "/netflix/api/movies/:id", "/netflix/api/movies/5")
	var response movie2.Movie
	err := json.NewDecoder(responseRecorder.Body).Decode(&response)
	require.NoError(t, err)
	assert.Equal(t, responseRecorder.Code, http.StatusOK)
	assert.Equal(t, response.Id, mockResponse.Id)
	assert.Equal(t, response, mockResponse)
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
