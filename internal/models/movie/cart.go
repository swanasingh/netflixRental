package movie

type CartItem struct {
	MovieId int
	UserId  int
	Status  bool
	movie   Movie
}

type Cart struct {
	CartItems []CartItem
}
