-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE books
DROP CONSTRAINT books_author_id_fkey,
ADD CONSTRAINT books_author_id_fkey
    FOREIGN KEY (author_id)
    REFERENCES authors(id)
    ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE books
DROP CONSTRAINT books_author_id_fkey,
ADD CONSTRAINT books_author_id_fkey
    FOREIGN KEY (author_id)
    REFERENCES authors(id);
-- +goose StatementEnd
