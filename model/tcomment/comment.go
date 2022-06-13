package tcomment

import (
	"BAT-douyin/model/tuser"
	"BAT-douyin/model/tvideo"
	"github.com/vmihailenco/msgpack"
	"time"
)

type Comment struct {
	ID        uint      `gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_date"`
	UpdatedAt time.Time
	Author    tuser.User   `gorm:"foreignKey:AuthorID;references:ID;"`
	AuthorID  uint         `json:"user_id"`
	Video     tvideo.Video `gorm:"foreignKey:VideoID;references:ID;"`
	VideoID   uint         `json:"video_id"`
	Content   string       `json:"content"`
}

func (s *Comment) MarshalBinary() ([]byte, error) {
	return msgpack.Marshal(s)
}

func (s *Comment) UnmarshalBinary(data []byte) error {
	return msgpack.Unmarshal(data, s)
}
