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
		Posts:    NewPostsService(posts, comments),
		Comments: NewCommentsService(),
	}
}

type Posts interface {
	CreatePost(post model.Post) (model.Post, error)
	GetPostById(id int) (model.Post, error)
	GetAllPosts(page, pageSize *int) ([]model.Post, error)
}

type Comments interface {
	CreateComment(comment model.Comment) (model.Comment, error)
	GetCommentsByPost(postId int, page *int, pageSize *int) ([]*model.Comment, error)
	GetRepliesOfComment(commentId int) ([]*model.Comment, error)
}
