package movie

type Movie struct {
	Id         int     `json:"id,string"`
	Title      string  `json:"title"`
	Year       string  `json:"year"`
	Rated      string  `json:"rated"`
	Released   string  `json:"released"`
	Runtime    string  `json:"runtime"`
	Genre      string  `json:"genre"`
	Director   string  `json:"director"`
	Writer     string  `json:"writer"`
	Actors     string  `json:"actors"`
	Plot       string  `json:"plot"`
	Language   string  `json:"language"`
	Country    string  `json:"country"`
	Awards     string  `json:"awards"`
	Poster     string  `json:"poster"`
	Metascore  int     `json:"metascore,string"`
	ImdbRating float32 `json:"imdbRating, string"`
	ImdbVotes  int     `json:"imdbVotes,string"`
	ImdbId     string  `json:"imdbId"`
	Type       string  `json:"type"`
	Dvd        string  `json:"dvd"`
	BoxOffice  string  `json:"boxOffice"`
	Production string  `json:"production"`
	Website    string  `json:"website"`
	Response   bool    `json:"response"`
}

type MovieResponse struct {
	Movies []Movie
}
