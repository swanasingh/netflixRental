package movie

import (
	"time"
)

type Invoice struct {
	InvoiceId   int       `json:"invoice_id"`
	OrderId     int       `json:"order_id"`
	InvoiceDate time.Time `json:"invoice_date"`
	MovieName   string    `json:"movie_name"`
	Price       float32   `json:"price"`
	Quantity    int       `json:"quantity"`
}
