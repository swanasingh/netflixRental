CREATE TABLE IF NOT EXISTS Cart
(
    id SERIAL PRIMARY KEY,
    user_id int NOT NULL,
    status bool default true,
    movie_id int NOT NULL,
    quantity INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (movie_id) references Movies(id),
    FOREIGN KEY (user_id) references Users(id)
)