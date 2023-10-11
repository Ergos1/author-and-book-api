-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE authors(
    id BIGSERIAL PRIMARY KEY NOT NULL, 
    name TEXT NOT NULL DEFAULT ''
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE authors;
-- +goose StatementEnd
