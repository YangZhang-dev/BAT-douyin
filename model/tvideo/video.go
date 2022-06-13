package tvideo

import (
	"BAT-douyin/model/tuser"
	"github.com/vmihailenco/msgpack"
	"time"
)

type Video struct {
	ID            uint `gorm:"primaryKey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Author        tuser.User `gorm:"foreignKey:AuthorID;references:ID;"`
	AuthorID      uint       `json:"author"`
	PlayUrl       string     `json:"play_url ,omitempty"`
	Title         string     `json:"title,omitempty"`
	CoverUrl      string     `json:"cover_url,omitempty"`
	FavoriteCount uint       `json:"favorite_count,omitempty"`
	CommentCount  uint       `json:"comment_count,omitempty"`
}

func (s *Video) MarshalBinary() ([]byte, error) {
	return msgpack.Marshal(s)
}

func (s *Video) UnmarshalBinary(data []byte) error {
	return msgpack.Unmarshal(data, s)
}
