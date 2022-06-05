package model

import "time"

//这里面是表结构

type User struct {
	ID              uint `gorm:"primaryKey"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Avatar          string `json:"avatar"`
	Signature       string `json:"signature"`
	BackgroundImage string `json:"background_image"`
	UserName        string `json:"name,omitempty"`
	Password        string
	FollowCount     uint `json:"follow_count,omitempty"`
	FollowerCount   uint `json:"follower_count,omitempty"`
	TotalFavorited  uint `json:"total_favorited"`
	FavoriteCount   uint `json:"favorite_count"`
}

type Video struct {
	ID            uint `gorm:"primaryKey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Author        User   `gorm:"foreignKey:AuthorID;references:ID;"`
	AuthorID      uint   `json:"author"`
	PlayUrl       string `json:"play_url ,omitempty"`
	Title         string `json:"title,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount uint   `json:"favorite_count,omitempty"`
	CommentCount  uint   `json:"comment_count,omitempty"`
}

type Comment struct {
	ID        uint      `gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_date"`
	UpdatedAt time.Time
	Author    User   `gorm:"foreignKey:AuthorID;references:ID;"`
	AuthorID  uint   `json:"user_id"`
	Video     Video  `gorm:"foreignKey:VideoID;references:ID;"`
	VideoID   uint   `json:"video_id"`
	Content   string `json:"content"`
}

type FavoriteVideo struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User  `gorm:"foreignKey:UserID;references:ID;"`
	UserID    uint  `json:"user_id"`
	Video     Video `gorm:"foreignKey:VideoID;references:ID;"`
	VideoID   uint  `json:"video_id"`
}

type FollowUser struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User `gorm:"foreignKey:UserID;references:ID;"`
	UserID    uint `json:"user_id"`
	ToUser    User `gorm:"foreignKey:ToUserID;references:ID;"`
	ToUserID  uint `json:"to_user_id"`
}
