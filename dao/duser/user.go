package duser

import (
	"BAT-douyin/commen"
	"BAT-douyin/dao/database"
	"BAT-douyin/model"
	"BAT-douyin/pkg/utils"
	"fmt"
	"gorm.io/gorm"

	"sync"
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
	var m sync.Mutex
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
	m.Lock()
	result := database.DB.Create(user)
	m.Unlock()
	if result.RowsAffected == 0 {
		return nil, false
	}
	return user, true
}

func GetByName(username string) (*model.User, bool) {
	user := &model.User{}
	result := database.DB.Where("user_name=?", username).Find(user)
	if result.RowsAffected == 0 {
		return nil, false
	}
	return user, true
}

func IsFollowUser(u, taru *model.User) bool {
	//自己默认关注自己，但在表里不记录
	if u.ID == taru.ID {
		return true
	}
	find := database.DB.Where("user_id=? and to_user_id=?", u.ID, taru.ID).Find(&model.FollowUser{})
	if find.RowsAffected == 0 {
		return false
	}
	return true
}

func FollowUser(u, taru *model.User) bool {
	var m sync.Mutex
	m.Lock()
	tx := database.DB.Begin()
	up1 := tx.Model(&model.User{}).Where("id=?", u.ID).Update("follow_count", gorm.Expr("follow_count+?", 1))
	up2 := tx.Model(&model.User{}).Where("id=?", taru.ID).Update("follower_count", gorm.Expr("follower_count+?", 1))
	up3 := tx.Create(&model.FollowUser{UserID: u.ID, ToUserID: taru.ID})

	m.Unlock()
	if up1.RowsAffected != 0 && up2.RowsAffected != 0 && up3.RowsAffected != 0 {
		tx.Commit()

		return true
	}
	fmt.Println("success")
	tx.Rollback()
	return false

}
func UnFollowUser(u, taru *model.User) bool {
	var m sync.Mutex
	m.Lock()
	tx := database.DB.Begin()
	up1 := tx.Model(&model.User{}).Where("id=?", u.ID).Update("follow_count", gorm.Expr("follow_count-?", 1))
	up2 := tx.Model(&model.User{}).Where("id=?", taru.ID).Update("follower_count", gorm.Expr("follower_count-?", 1))
	up3 := tx.Unscoped().Where("user_id=? and to_user_id=?", u.ID, taru.ID).Delete(&model.FollowUser{})
	m.Unlock()
	if up1.RowsAffected != 0 && up2.RowsAffected != 0 && up3.RowsAffected != 0 {
		tx.Commit()
		return true
	}
	tx.Rollback()
	return false
}
func FollowerUserList(u *model.User) []model.FollowUser {
	res := &[]model.FollowUser{}
	database.DB.Where("to_user_id", u.ID).Find(res)
	return *res

}
func FollowUserList(u *model.User) []model.FollowUser {
	res := &[]model.FollowUser{}
	database.DB.Where("user_id", u.ID).Find(res)
	return *res

}
