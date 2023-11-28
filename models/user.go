package models

import "gin-ranking/dao"

type User struct {
	Id   int
	Name string
}

func (User) TableName() string {
	return "user"
}

func GetUser(id int) (User, error) {
	var user User
	err := dao.Db.Where("id = ?", id).First(&user).Error
	return user, err
}

func AddUser(username string) (User, error) {
	user := User{Name: username}
	err := dao.Db.Create(&user).Error
	return user, err
}
