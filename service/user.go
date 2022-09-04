package service

import (
	"ginDemo/global"
	"ginDemo/model/database"
)

func QueryUserByUsername(username string) (user database.User, notFound bool) {
	notFound = global.DB.Where("username = ?", username).First(&user).RecordNotFound()
	return user, notFound
}

func QueryUserByUserID(userID uint64) (user database.User, notFound bool) {
	notFound = global.DB.Where("user_id = ?", userID).First(&user).RecordNotFound()
	return user, notFound
}

func QueryUserByEmail(email string) (user database.User, notFound bool) {
	notFound = global.DB.Where("email = ?", email).First(&user).RecordNotFound()
	return user, notFound
}

func CreateUser(user *database.User) (err error) {
	err = global.DB.Create(&user).Error
	return err
}

func UpdateUser(user *database.User) (err error) {
	err = global.DB.Save(user).Error
	return err
}
