package duser

import (
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
