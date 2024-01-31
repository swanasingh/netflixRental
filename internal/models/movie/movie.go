package movie

type Movie struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbId string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

type MovieResponse struct {
	Search []Movie
}
