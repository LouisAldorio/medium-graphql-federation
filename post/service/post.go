package service

import (
	"myapp/config"
	"myapp/graph/model"
)

func PostsGet() ([]*model.Post, error) {

	db,sql := config.ConnectGorm()
	defer sql.Close()

	var posts []*model.Post
	err := db.Table("posts").Find(&posts).Error

	return posts, err
}

func PostCreate(input model.NewPost) (*model.Post, error) {
	
	db,sql := config.ConnectGorm()
	defer sql.Close()

	newPost := model.Post{
		Title: input.Title,
		Content: input.Content,
		UserID: input.UserID,
	}

	err := db.Table("posts").Create(&newPost).Error

	return &newPost, err
}