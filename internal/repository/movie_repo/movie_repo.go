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
	GetInvoice(orderId int) ([]movie.Invoice, movie.User, error)
	GetUserDetails(userId int) (movie.User, error)
}

type movieRepo struct {
	*sql.DB
}

func (m movieRepo) GetUserDetails(userId int) (movie.User, error) {
	var user movie.User
	row := m.DB.QueryRow("SELECT id,name,email,address from users where id=$1", userId)

	err := row.Scan(&user.UserId, &user.Name, &user.Email, &user.Address)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("Invalid Id")
		}
	}
	return user, nil
}

func (m movieRepo) GetInvoice(orderId int) ([]movie.Invoice, movie.User, error) {

	var invoices []movie.Invoice
	var user movie.User
	rows, err := m.DB.Query("SELECT i.id,op.order_id,m.title,op.quantity,op.price,o.created_at,u.id "+
		"FROM invoices i inner join order_products op on i.order_id=op.order_id "+
		"inner join orders o on o.id =i.order_id "+
		"inner join movies m on m.id=op.movie_id "+
		"inner join users u on u.id =o.user_id "+
		"where i.order_id =$1", orderId)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal("Could Not Fetch data From DB")
	}

	var userId int
	for rows.Next() {
		var inv movie.Invoice
		if err = rows.Scan(&inv.InvoiceId, &inv.OrderId, &inv.MovieName, &inv.Quantity, &inv.Price, &inv.InvoiceDate, &userId); err != nil {
			return nil, user, err
		}
		invoices = append(invoices, inv)
	}
	if userId != 0 {
		user, _ = m.GetUserDetails(userId)
	}

	return invoices, user, nil
}

func (m movieRepo) CreateOrder(order movie.OrderPayload) error {
	/*Result, err := m.DB.Exec("insert into orders (user_id,total_amount,status) values($1, $2, $3)", order.UserId, order.TotalAmount, order.Status)
	if err != nil {
		return errors.New("error in creating order")
	}*/

	stmt, err := m.DB.Prepare("INSERT INTO orders (user_id,total_amount,status) values($1, $2, $3) RETURNING id")
	if err != nil {
		return err
	}

	var orderId int
	err = stmt.QueryRow(order.UserId, order.TotalAmount, order.Status).Scan(&orderId)
	if err != nil {
		return errors.New("error in creating order")
	}

	for _, mv := range order.Items {
		_, err = m.DB.Exec("insert into order_products (order_id,movie_id,quantity,price) values($1, $2, $3,$4)", orderId, mv.MovieId, mv.Quantity, mv.Price)
		if err != nil {
			return errors.New("error in saving order_products")
		}
	}
	_, err = m.DB.Exec("INSERT INTO invoices (order_id) values($1)", orderId)
	if err != nil {
		return err
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
