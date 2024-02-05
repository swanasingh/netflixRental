CREATE TABLE IF NOT EXISTS Cart
(
    id SERIAL PRIMARY KEY,
    user_id int,
    status bool,
    movie_id int,
    FOREIGN KEY (movie_id) references Movies(id)
)