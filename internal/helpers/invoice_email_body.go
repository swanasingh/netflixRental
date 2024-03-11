package helpers

import (
	"fmt"
	movie2 "netflixRental/internal/models/movie"
	"strconv"
)

func GenerateInvoiceEmailBody(invoices movie2.Invoices) string {
	table := "<table border=\"1\"><tr><th>InvoiceID</th><th>InvoiceDate</th><th>Movie</th><th>Price</th><th>Qty</th></tr>"
	orderId := invoices.InvoicesList[0].OrderId
	for _, item := range invoices.InvoicesList {
		row := fmt.Sprintf("<tr><td>%d</td><td>%s</td><td>%s</td><td>%f</td><td>%d</td></tr>",
			item.InvoiceId, item.InvoiceDate,
			item.MovieName, item.Price, item.Quantity)
		table += row
	}
	table += "</table>"

	// Email body with HTML formatting
	body := fmt.Sprintf("Subject: Order Invoice "+strconv.Itoa(orderId)+" \r\n"+
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"+
		"<h2>Information Table</h2>"+
		"%s", table)

	return body
}
