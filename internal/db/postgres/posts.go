package postgres

import (
	"github.com/jmoiron/sqlx"
)

type PostsPostgres struct {
	db *sqlx.DB
}

func NewPostsPostgres(db *sqlx.DB) *PostsPostgres {
	return &PostsPostgres{db: db}
}

func CreatePost() {

}
func GetAllPosts() {

}
func GetPostByID() {

}
