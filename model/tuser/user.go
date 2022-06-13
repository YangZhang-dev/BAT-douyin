package tuser

import (
	"github.com/vmihailenco/msgpack"
	"time"
)

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

func (s *User) MarshalBinary() ([]byte, error) {
	return msgpack.Marshal(s)
}

func (s *User) UnmarshalBinary(data []byte) error {
	return msgpack.Unmarshal(data, s)
}
