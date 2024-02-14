package movie

type CartItem struct {
	MovieId  int
	UserId   int
	Status   bool
	Quantity int
	movie    Movie
}

type Cart struct {
	CartItems []CartItem
}
