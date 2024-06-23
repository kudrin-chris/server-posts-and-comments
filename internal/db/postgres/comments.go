package postgres

import (
	"github.com/jmoiron/sqlx"
)

type CommentsPostgres struct {
	db *sqlx.DB
}

func NewCommentsPostgres(db *sqlx.DB) *CommentsPostgres {
	return &CommentsPostgres{db: db}
}
func CreateComment() {

}
func CommentsSubscription() {

}
func SubscriptionResolver() {

}
