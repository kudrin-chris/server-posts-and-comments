package postgres

import (
	"server/graph/model"

	"github.com/jmoiron/sqlx"
)

type CommentsPostgres struct {
	db *sqlx.DB
}

func NewCommentsPostgres(db *sqlx.DB) *CommentsPostgres {
	return &CommentsPostgres{db: db}
}

func (c CommentsPostgres) CreateComment(comment model.Comment) (model.Comment, error) {

	tx, err := c.db.Begin()
	if err != nil {
		return model.Comment{}, err
	}

	sqlText := "INSERT INTO comments (content, author, post, reply_to) VALUES ($1, $2, $3, $4) RETURNING id"
	row := tx.QueryRow(sqlText, comment.Content, comment.Author, comment.Post, comment.ReplyTo)
	if err := row.Scan(&comment.ID); err != nil {
		tx.Rollback()
		return model.Comment{}, err
	}

	return comment, tx.Commit()

}

func (c CommentsPostgres) GetCommentsByPost(postId, limit int) ([]*model.Comment, error) {

	sqlText := "SELECT * FROM comments WHERE post = $1 AND reply_to IS NULL ORDER BY created_at LIMIT $2"

	var comments []*model.Comment

	if err := c.db.Select(&comments, sqlText, limit); err != nil {
		return nil, err
	}

	return comments, nil
}

func (c CommentsPostgres) GetRepliesOfComment(commentId int) ([]*model.Comment, error) {

	sqlText := "SELECT * FROM comments WHERE reply_to = $1"

	var comments []*model.Comment

	if err := c.db.Select(&comments, sqlText, commentId); err != nil {
		return nil, err
	}

	return comments, nil
}
