CREATE TABLE IF NOT EXISTS Invoices(
 id SERIAL PRIMARY KEY,
 order_id INT NOT NULL,
 FOREIGN KEY (order_id) references Orders(id)
)
