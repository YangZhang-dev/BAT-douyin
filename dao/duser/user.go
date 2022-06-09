package duser

import (
	"BAT-douyin/commen"
	"BAT-douyin/dao/database"
	"BAT-douyin/model"
	"BAT-douyin/pkg/utils"
)

func GetById(uid uint) (*model.User, bool) {
	if uid == 0 {
		return nil, false
	}
	user := &model.User{}
	affected := database.DB.Where("id=?", uid).Find(user).RowsAffected
	if affected == 0 {
		return nil, false
	}
	return user, true
}

func GetByToken(token string) *model.User {
	claims, ok := utils.ValidateJwt(token)
	if !ok {
		return nil
	}
	u, ok := GetById(claims.UserId)
	if !ok {
		return nil
	}
	return u
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

func GetByName(username string) (*model.User, bool) {
	user := &model.User{}
	result := database.DB.Where("UserName=?", username).Find(user)
	if result.RowsAffected == 0 {
		return nil, false
	}
	return user, true
}
