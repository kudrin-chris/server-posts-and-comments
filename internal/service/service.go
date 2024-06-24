package service

import (
	"server/graph/model"
)

type Services struct {
	Posts
	Comments
}

func NewServices(posts Posts, comments Comments) *Services {
	return &Services{
		Posts: NewPostsServices(posts, comments),
		Comments: NewCommentsService(PostsServices{
			Posts:    posts,
			Comments: comments,
		}),
	}
}

type Posts interface {
	CreatePost(post model.Post) (model.Post, error)
	GetPostById(id int) (model.Post, error)
	GetAllPosts(limit int) ([]model.Post, error)
}

type Comments interface {
	CreateComment(comment model.Comment) (model.Comment, error)
	GetCommentsByPost(postId int, limit int) ([]*model.Comment, error)
	GetRepliesOfComment(commentId int) ([]*model.Comment, error)
}
