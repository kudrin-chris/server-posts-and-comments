package service

import "server/graph/model"

type CommentsService struct {
	PostGetter PostGetter
}

type PostGetter interface {
	GetPostById(id int) (model.Post, error)
}

func NewCommentsService(getter PostGetter) *CommentsService {
	return &CommentsService{PostGetter: getter}
}
func (c CommentsService) CreateComment(comment model.Comment) (model.Comment, error) {
	return model.Comment{}, nil
}
func (c CommentsService) GetCommentsByPost(postId int, limit int) ([]*model.Comment, error) {
	return nil, nil
}
func (c CommentsService) GetRepliesOfComment(commentId int) ([]*model.Comment, error) {
	return nil, nil
}
