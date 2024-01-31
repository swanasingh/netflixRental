package MovieService

import (
	"encoding/json"
	"log"
	"net/http"
	"netflixRental/internal/models/movie"
)

type MovieService interface {
	Get() []movie.Movie
}

type movieService struct {
	movie movie.Movie
}

func (m movieService) Get() []movie.Movie {
	var movies movie.MovieResponse
	resp, err := http.Get("http://www.omdbapi.com/?s=Barbie&type=movie&page=1&apikey=ead19c9f")
	if err != nil {
		log.Fatal("The service cannot fetch movies", err.Error())
	}

	if resp.StatusCode == http.StatusOK {
		err1 := json.NewDecoder(resp.Body).Decode(&movies)
		if err1 != nil {
			log.Fatal("cannot unmarshal the json response")
		}
	}
	return movies.Search
}

func NewMovieService() MovieService {
	return &movieService{}
}
