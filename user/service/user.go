package service

import (
	"myapp/config"
	"myapp/graph/model"
)

func UsersGet() ([]*model.User, error) {

	db,sql := config.ConnectGorm()
	defer sql.Close()

	var users []*model.User
	err := db.Table("users").Find(&users).Error

	return users,err
}

func UserCreate(input model.NewUser) (*model.User, error) {
	
	db,sql := config.ConnectGorm()
	defer sql.Close()

	newUser := model.User{
		Name: input.Name,
	}

	err := db.Table("users").Create(&newUser).Error

	return &newUser,err
}

func UserFindByID(userID int) (*model.User, error){
	
	db,sql := config.ConnectGorm()
	defer sql.Close()

	var user model.User
	err := db.Table("users").Where("id = ?", userID).First(&user).Error

	return &user,err
}