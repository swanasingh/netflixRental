package movie_repo

import (
	"database/sql"
	"fmt"
	"log"
	"netflixRental/internal/models/movie"
	"strings"
)

type MovieRepository interface {
	Get() []movie.Movie
	FetchByCriteria(Genre, Actor string, Year int) []movie.Movie
}

type movieRepo struct {
	*sql.DB
}

func (m movieRepo) Get() []movie.Movie {
	query := "select * from movies"
	return executeQuery(m, query)
}

func (m movieRepo) FetchByCriteria(Genre, Actor string, Year int) []movie.Movie {
	fmt.Println("inside repo filter by criteria")
	fmt.Println(Genre, Actor, Year)
	Genre = strings.ReplaceAll(Genre, "\"", "")
	Actor = strings.ReplaceAll(Actor, "\"", "")
	//if Genre == "" {
	//	Genre = strings.ReplaceAll(Genre, "\"", "'")
	//}
	//if Actor == "" {
	//	Actor = "" + "''"
	//}

	query := fmt.Sprintf("select * from movies where actors like '%s' or year = %d or genre like '%s'", Actor, Year, Genre)
	return executeQuery(m, query)
}

func executeQuery(m movieRepo, query string) []movie.Movie {
	var movieList movie.MovieResponse
	fmt.Println("inside repo execute query")
	fmt.Println(query)
	rows, err := m.DB.Query(query)
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
			&movie.Production, &movie.Website, &movie.Response); err != nil {
			movieList.Movies = append(movieList.Movies, movie)
		}

	}
	return movieList.Movies
}

func NewMovieRepository(db *sql.DB) MovieRepository {
	return &movieRepo{db}
}
