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

func (p InputPost) FromInput() Post {
	return Post{
		Name:            p.Name,
		Author:          p.Author,
		Content:         p.Content,
		CommentsAllowed: p.CommentsAllowed,
	}
}
func (c InputComment) FromInput() Comment {
	return Comment{
		Author:  c.Author,
		Content: c.Content,
		Post:    c.Post,
		ReplyTo: c.ReplyTo,
	}
}
func (p Post) ToGraph() *PostGraph {
	return &PostGraph{
		ID:        p.ID,
		Date: 	   p.Date,
		Name:      p.Name,
		Author:    p.Author,
		Content:   p.Content,
	}
}

func ToPostGraph(posts []Post) []*PostGraph {
	res := make([]*PostGraph, 0, len(posts))

	for _, post := range posts {
		res = append(res, post.ToGraph())
	}

	return res
}