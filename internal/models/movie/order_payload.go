package movie

type OrderPayload struct {
	UserId      int     `json:"user_id"`
	Status      string  `json:"status"`
	TotalAmount float64 `json:"total_amount"`
	Items       []Item  `json:"movies"`
}

type Item struct {
	MovieId  int     `json:"movie_id"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}
