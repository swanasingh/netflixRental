package movie

type CartRequest struct {
	MovieId int `json:"movie_id"`
	UserId  int `json:"user_id"`
}
