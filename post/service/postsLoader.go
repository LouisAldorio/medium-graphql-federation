package service

import (
	"myapp/graph/model"
	"myapp/config"
)

func PostsLoader(userIDs []int) ([][]*model.Post, []error) {
	db,sql := config.ConnectGorm()
	defer sql.Close()

	var posts []*model.Post
	err := db.Table("posts").Where("user_id IN (?)",userIDs).Find(&posts).Error
	if err != nil {
		return nil, []error{err}
	}

	postsMappedByUserID := map[int][]*model.Post{}
	for _,v := range posts {
		postsMappedByUserID[v.UserID] = append(postsMappedByUserID[v.UserID], v)
	}

	response := make([][]*model.Post, len(userIDs))
	for i,userID := range userIDs {
		response[i] = postsMappedByUserID[userID]
	}

	return response, nil
}