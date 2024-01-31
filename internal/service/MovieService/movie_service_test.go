package MovieService

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMovieService(t *testing.T) {
	//t.Run("Return List Of Movies", Get())

	//want := []movie.Movie{{
	//	//	Title: "Barbie",
	//	//	Year: "2023",
	//	//	ImdbId: "tt1517268",
	//	//	Type: "movie",
	//	//	Poster: "https://m.media-amazon.com/images/M/MV5BNjU3N2QxNzYtMjk1NC00MTc4LTk1NTQtMmUxNTljM2I0NDA5XkEyXkFqcGdeQXVyODE5NzE3OTE@._V1_SX300.jpg"
	//	//},{
	//	//	Title: "Barbie as The Princess and the Pauper",
	//	//	Year: "2004",
	//	//	ImdbId: "tt0426955",
	//	//	Type: "movie",
	//	//	Poster: "https://m.media-amazon.com/images/M/MV5BMGY5MzU3MzItNDBjMC00YjQzLWEzMTUtMGMxMTEzYjhkMGNkXkEyXkFqcGdeQXVyNDE5MTU2MDE@._V1_SX300.jpg"
	//	//},
	//	//}
	movieService := NewMovieService()
	got := len(movieService.Get())
	assert.Equal(t, got, 10)
}
