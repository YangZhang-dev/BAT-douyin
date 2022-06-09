package duser

import (
	"BAT-douyin/commen"
	"BAT-douyin/dao/database"
	"BAT-douyin/model"
	"golang.org/x/crypto/bcrypt"
)

func GetByName(username string) (*model.User, bool) {
	user := &model.User{}
	result := database.DB.Where("UserName=?", username).Find(user)
	if result.RowsAffected == 0 {
		return nil, false
	}
	return user, true
}
func Create(username, password string) (*model.User, bool) {
	user := &model.User{
		Avatar:          commen.Avatar,
		Signature:       commen.Signature,
		BackgroundImage: commen.BackgroundImage,
		UserName:        username,
		Password:        password,
		FollowCount:     0,
		FollowerCount:   0,
		TotalFavorited:  0,
		FavoriteCount:   0,
	}
	result := database.DB.Create(user)
	if result.RowsAffected == 0 {
		return nil, false
	}
	return user, true
}
func CheckPassword(user model.User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false
	}
	return true
}
