package postgres

import (
	"server/graph/model"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type PostsPostgres struct {
	db *sqlx.DB
}

func NewPostsPostgres(db *sqlx.DB) *PostsPostgres {
	return &PostsPostgres{db: db}
}

func (p PostsPostgres) CreatePost(post model.Post) (model.Post, error) {
	sqlText := "INSERT INTO Posts (name, content, author, comments_allowed) VALUES ($1, $2, $3, $4) RETURNING id, created_at"
	tx, err := p.db.Begin()
	if err != nil {
		return model.Post{}, err
	}
	row := tx.QueryRow(sqlText, post.Name, post.Content, post.Author, post.CommentsAllowed)
	if err := row.Scan(&post.ID); err != nil {
		tx.Rollback()
		return model.Post{}, err
	}

	return post, tx.Commit()
}
func (p PostsPostgres) GetAllPosts(limit int) ([]model.Post, error) {
	sqlText := "SELECT * FROM post ORDER BY id LIMIT " + strconv.Itoa(limit)

	var posts []model.Post
	if err := p.db.Get(&posts, sqlText); err != nil {
		return posts, err
	}
	return posts, nil
}
func (p PostsPostgres) GetPostById(id int) (model.Post, error) {

	sqlText := "SELECT * FROM post WHERE id = " + strconv.Itoa(id)

	var post model.Post
	if err := p.db.Get(&post, sqlText); err != nil {
		return post, err
	}

	return model.Post{}, nil
}
