package movie_repo

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"netflixRental/internal/models/movie"
	"strings"
)

type MovieRepository interface {
	Get(criteria movie.Criteria) []movie.Movie
	GetMovieDetails(id int) (movie.Movie, error)
	SaveCartData(cartItem movie.CartItem) error
	ViewCart(user_id int) []movie.Movie
	CreateOrder(order movie.OrderPayload) error
}

type movieRepo struct {
	*sql.DB
}

func (m movieRepo) CreateOrder(order movie.OrderPayload) error {
	/*Result, err := m.DB.Exec("insert into orders (user_id,total_amount,status) values($1, $2, $3)", order.UserId, order.TotalAmount, order.Status)
	if err != nil {
		return errors.New("error in creating order")
	}*/

	stmt, err := m.DB.Prepare("INSERT INTO orders (user_id,total_amount,status) values($1, $2, $3) RETURNING id")
	if err != nil {
		panic(err)
	}

	var orderId int
	err = stmt.QueryRow(order.UserId, order.TotalAmount, order.Status).Scan(&orderId)
	if err != nil {
		return errors.New("error in creating order")
	}

	for _, mv := range order.Items {

		fmt.Println(mv)
		_, err = m.DB.Exec("insert into order_products (order_id,movie_id,quantity,price) values($1, $2, $3,$4)", orderId, mv.MovieId, mv.Quantity, mv.Price)
		if err != nil {
			return errors.New("error in saving order_products")
		}
	}

	return nil
}

func (m movieRepo) ViewCart(userId int) []movie.Movie {
	var cartItems []movie.Movie
	rows, err := m.DB.Query("select m.id,title,year,price from movies m "+
		"inner join cart c on m.id = c.movie_id where c.user_id=$1", userId)

	if err != nil {
		log.Fatal("Could Not Fetch data From DB")
	}
	fmt.Println(rows)
	for rows.Next() {
		var movie movie.Movie
		if err = rows.Scan(&movie.Id, &movie.Title,
			&movie.Year, &movie.Price); err != nil {
			log.Fatal("Could Not Fetch data From DB")
		}
		cartItems = append(cartItems, movie)
	}
	return cartItems
}

func (m movieRepo) SaveCartData(cartItem movie.CartItem) error {

	_, err := m.DB.Exec("insert into cart (user_id,quantity,movie_id) values($1, $2, $3)", cartItem.UserId, cartItem.Quantity, cartItem.MovieId)
	if err != nil {
		return errors.New("movie is not present in database")
	}

	return nil
}

func (m movieRepo) GetMovieDetails(id int) (movie.Movie, error) {
	var mv movie.Movie
	query := fmt.Sprintf("select * from movies where id =%d", id)
	row := m.DB.QueryRow(query)

	err := row.Scan(&mv.Id, &mv.Title,
		&mv.Year, &mv.Rated, &mv.Released,
		&mv.Runtime, &mv.Genre, &mv.Director,
		&mv.Writer, &mv.Actors, &mv.Plot, &mv.Language,
		&mv.Country, &mv.Awards, &mv.Poster, &mv.Metascore,
		&mv.ImdbRating, &mv.ImdbVotes, &mv.ImdbId, &mv.Type, &mv.Dvd, &mv.BoxOffice,
		&mv.Production, &mv.Website, &mv.Response, &mv.Price)

	if err != nil {
		if err == sql.ErrNoRows {
			return mv, errors.New("Invalid Id")
		}
	}
	return mv, nil
}

func NewMovieRepository(db *sql.DB) MovieRepository {
	return &movieRepo{db}
}

func (m movieRepo) Get(criteria movie.Criteria) []movie.Movie {

	var rows *sql.Rows
	var err error
	var movieList movie.MovieResponse

	fmt.Println(criteria)
	if criteria != (movie.Criteria{}) {
		criteria.Genre = strings.ReplaceAll(criteria.Genre, "\"", "")
		criteria.Actors = strings.ReplaceAll(criteria.Actors, "\"", "")
		query := fmt.Sprintf("select * from movies where actors like '%s' or year = %d or genre like '%s'", criteria.Actors, criteria.Year, criteria.Genre)
		rows, err = m.DB.Query(query)
		//rows, err = m.DB.Query("select * from movies where actors like '%$1' or year = $2 or genre like '%$3'", criteria.Actors, criteria.Year, criteria.Genre)
	} else {
		rows, err = m.DB.Query("select * from movies")
	}

	if err != nil {
		log.Fatal("Could Not Fetch data From DB")
	}
	for rows.Next() {
		var movie movie.Movie
		if err := rows.Scan(&movie.Id, &movie.Title,
			&movie.Year, &movie.Rated, &movie.Released,
			&movie.Runtime, &movie.Genre, &movie.Director,
			&movie.Writer, &movie.Actors, &movie.Plot, &movie.Language,
			&movie.Country, &movie.Awards, &movie.Poster, &movie.Metascore,
			&movie.ImdbRating, &movie.ImdbVotes, &movie.ImdbId, &movie.Type, &movie.Dvd, &movie.BoxOffice,
			&movie.Production, &movie.Website, &movie.Response, &movie.Price); err != nil {
			movieList.Movies = append(movieList.Movies, movie)
		}

	}
	return movieList.Movies
}
