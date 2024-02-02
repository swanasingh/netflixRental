package movie_repo

import (
	"database/sql"
	"fmt"
	"log"
	"netflixRental/internal/models/movie"
	"strings"
)

type MovieRepository interface {
	Get(criteria movie.Criteria) []movie.Movie
	GetMovieDetails(id int) movie.Movie
}

type movieRepo struct {
	*sql.DB
}

func (m movieRepo) GetMovieDetails(id int) movie.Movie {
	var mv movie.Movie
	query := fmt.Sprintf("select * from movies where id =%d", id)
	row := m.DB.QueryRow(query)
	row.Scan(&mv.Id, &mv.Title,
		&mv.Year, &mv.Rated, &mv.Released,
		&mv.Runtime, &mv.Genre, &mv.Director,
		&mv.Writer, &mv.Actors, &mv.Plot, &mv.Language,
		&mv.Country, &mv.Awards, &mv.Poster, &mv.Metascore,
		&mv.ImdbRating, &mv.ImdbVotes, &mv.ImdbId, &mv.Type, &mv.Dvd, &mv.BoxOffice,
		&mv.Production, &mv.Website, &mv.Response)
	return mv
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
			&movie.Production, &movie.Website, &movie.Response); err != nil {
			movieList.Movies = append(movieList.Movies, movie)
		}

	}
	return movieList.Movies
}
