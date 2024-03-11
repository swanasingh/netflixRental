package MovieService

import (
	"errors"
	"netflixRental/internal/models/movie"
	"netflixRental/internal/repository/movie_repo"
)

type MovieService interface {
	Get(criteria movie.Criteria) []movie.Movie
	GetMovieDetails(id int) (movie.Movie, error)
	AddToCart(cartItem movie.CartItem) error
	ViewCart(user_id int) []movie.Movie
	CreateOrder(order movie.OrderPayload) error
	GetInvoice(orderId int) ([]movie.Invoice, movie.User, error)
}

func NewMovieService(movieRepository movie_repo.MovieRepository) MovieService {
	return &movieService{movieRepository}
}

type movieService struct {
	movieRepo movie_repo.MovieRepository
}

func (m movieService) GetInvoice(orderId int) ([]movie.Invoice, movie.User, error) {
	invoices, user, err := m.movieRepo.GetInvoice(orderId)
	if (err != nil) || user == (movie.User{}) {
		return nil, user, errors.New("provide valid input")
	} else {
		return invoices, user, nil
	}

	return m.movieRepo.GetInvoice(orderId)
}

func (m movieService) CreateOrder(order movie.OrderPayload) error {
	return m.movieRepo.CreateOrder(order)
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
