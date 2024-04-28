CREATE TABLE IF NOT EXISTS books (
    id SERIAL PRIMARY KEY,
    isbn TEXT UNIQUE NOT NULL,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    author TEXT NOT NULL,
    image_url TEXT NOT NULL
);