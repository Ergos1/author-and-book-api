-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE books(
    id BIGSERIAL PRIMARY KEY NOT NULL, 
    name TEXT NOT NULL DEFAULT '',
    rating INTEGER NOT NULL DEFAULT 0,
    author_id INTEGER REFERENCES authors(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE books;
-- +goose StatementEnd
