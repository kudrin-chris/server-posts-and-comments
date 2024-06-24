-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Post (
    id serial primary key,
    date varchar(20) not null,
    name varchar(100) not null,
    content varchar(2000) not null,
    author varchar(100) not null,
    comments_allowed boolean
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Post;
-- +goose StatementEnd
