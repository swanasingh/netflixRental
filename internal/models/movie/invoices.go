package movie

type Invoices struct {
	User         User      `json:"user"`
	InvoicesList []Invoice `json:"products"`
}
