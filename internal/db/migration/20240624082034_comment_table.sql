-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Comment
(
    id serial primary key,
    date varchar(20),
    content varchar(2000),
    author varchar(100) not null,
    post int not null,
    FOREIGN KEY (post) REFERENCES Post(id) ON DELETE CASCADE ON UPDATE CASCADE ,
    reply_to int,
    FOREIGN KEY (reply_to) REFERENCES Comment(id) ON DELETE SET NULL ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Comment;
-- +goose StatementEnd
