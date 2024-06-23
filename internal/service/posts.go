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

func CreatePost(post model.Post) (model.Post, error) {
	return model.Post{}, nil
}
func GetPostById(id int) (model.Post, error) {
	return model.Post{}, nil
}
func GetAllPosts(page, pageSize *int) ([]model.Post, error) {
	return []model.Post{}, nil
}
