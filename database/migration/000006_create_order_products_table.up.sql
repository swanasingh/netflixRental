CREATE TABLE IF NOT EXISTS order_products (
    order_id INT NOT NULL,
    movie_id INT NOT NULL,
    quantity INT NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (movie_id) references Movies(id),
    FOREIGN KEY (order_id) references Orders(id),
    PRIMARY KEY (order_id, movie_id)
);