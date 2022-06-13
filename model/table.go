package model

import (
	"BAT-douyin/model/tuser"
	"BAT-douyin/model/tvideo"
	"time"
)

//这里面是表结构

type FavoriteVideo struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	User      tuser.User   `gorm:"foreignKey:UserID;references:ID;"`
	UserID    uint         `json:"user_id"`
	Video     tvideo.Video `gorm:"foreignKey:VideoID;references:ID;"`
	VideoID   uint         `json:"video_id"`
}

type FollowUser struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	User      tuser.User `gorm:"foreignKey:UserID;references:ID;"`
	UserID    uint       `json:"user_id"`
	ToUser    tuser.User `gorm:"foreignKey:ToUserID;references:ID;"`
	ToUserID  uint       `json:"to_user_id"`
}
