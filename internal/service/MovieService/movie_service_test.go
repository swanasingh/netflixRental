package MovieService

import (
	"github.com/stretchr/testify/assert"
	"netflixRental/internal/models/movie"
	"netflixRental/internal/repository/movie_repo/mocks"
	"testing"
)

func TestMovieService(t *testing.T) {

	t.Run("validate count of Movies", func(t *testing.T) {
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
		var mockResponse []movie.Movie
		movies := append(mockResponse, m1)

		mockRepository := mocks.MovieRepository{}
		mockRepository.On("Get").Return(mockResponse)
		movieService := NewMovieService(&mockRepository)

		got := movieService.Get()
		assert.Equal(t, len(got), 1)
		assert.Equal(t, movies, got)

	})
}
