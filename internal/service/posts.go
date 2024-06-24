package service

import "server/graph/model"

type PostsServices struct {
	Posts
	Comments
}

func NewPostsServices(posts Posts, comments Comments) *PostsServices {
	return &PostsServices{
		Posts:    posts,
		Comments: comments,
	}
}

func (p PostsServices) CreatePost(post model.Post) (model.Post, error) {
	return model.Post{}, nil
}
func (p PostsServices) GetPostById(id int) (model.Post, error) {
	return model.Post{}, nil
}
func (p PostsServices) GetAllPosts(limit int) ([]model.Post, error) {
	return []model.Post{}, nil
}
