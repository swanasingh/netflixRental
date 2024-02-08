package MovieService

import (
	"netflixRental/internal/models/movie"
	"netflixRental/internal/repository/movie_repo"
)

type MovieService interface {
	Get(criteria movie.Criteria) []movie.Movie
	GetMovieDetails(id int) (movie.Movie, error)
	AddToCart(cartItem movie.CartItem) error
	ViewCart(user_id int) []movie.Movie
}

func NewMovieService(movieRespository movie_repo.MovieRepository) MovieService {
	return &movieService{movieRespository}
}

type movieService struct {
	movieRepo movie_repo.MovieRepository
}

func (m movieService) ViewCart(user_id int) []movie.Movie {
	return m.movieRepo.ViewCart(user_id)
}

func (m movieService) AddToCart(cartItem movie.CartItem) error {
	return m.movieRepo.SaveCartData(cartItem)
}

func (m movieService) GetMovieDetails(id int) (movie.Movie, error) {
	return m.movieRepo.GetMovieDetails(id)
}

func (m movieService) Get(criteria movie.Criteria) []movie.Movie {
	var movies []movie.Movie
	movies = m.movieRepo.Get(criteria)
	return movies
}

/*func (m movieService) Get() []movie.Movie {
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
	return movies.Movies
}*/
