package MovieService

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMovieService(t *testing.T) {

	t.Run("validate count of Movies", func(t *testing.T) {
		movieService := NewMovieService()
		got := len(movieService.Get())
		assert.Equal(t, got, 10)
	})
}
