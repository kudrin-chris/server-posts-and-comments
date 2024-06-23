// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Comment struct {
	ID      string     `json:"id"`
	Date    string     `json:"date"`
	Author  string     `json:"author"`
	Content string     `json:"content"`
	Post    string     `json:"post"`
	Replies []*Comment `json:"replies,omitempty"`
	ReplyTo *string    `json:"replyTo,omitempty"`
}

type InputComment struct {
	Author  string  `json:"author"`
	Content string  `json:"content"`
	Post    string  `json:"post"`
	ReplyTo *string `json:"replyTo,omitempty"`
}

type InputPost struct {
	Name            string `json:"name"`
	Content         string `json:"content"`
	Author          string `json:"author"`
	CommentsAllowed bool   `json:"commentsAllowed"`
}

type Mutation struct {
}

type Post struct {
	ID              string     `json:"id"`
	Date            string     `json:"date"`
	Name            string     `json:"name"`
	Author          string     `json:"author"`
	Content         string     `json:"content"`
	CommentsAllowed bool       `json:"commentsAllowed"`
	Comments        []*Comment `json:"comments,omitempty"`
}

type PostGraph struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Name    string `json:"name"`
	Author  string `json:"author"`
	Content string `json:"content"`
}

type Query struct {
}

type Subscription struct {
}
