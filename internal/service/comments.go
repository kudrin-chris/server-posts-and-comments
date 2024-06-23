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
func CreateComment(comment model.Comment) (model.Comment, error) {
	return model.Comment{}, nil
}
func GetCommentsByPost(postId int, page *int, pageSize *int) ([]*model.Comment, error) {
	return nil, nil
}
func GetRepliesOfComment(commentId int) ([]*model.Comment, error) {
	return nil, nil
}
