package dcomment

import (
	"BAT-douyin/dao/database"

	"BAT-douyin/model/tcomment"
	"BAT-douyin/model/tuser"
	"BAT-douyin/model/tvideo"
	"gorm.io/gorm"
	"sync"
)

func GetById(cid uint) (*tcomment.Comment, bool) {
	if cid == 0 {
		return nil, false
	}
	comment := &tcomment.Comment{}
	affected := database.DB.Where("id=?", cid).Find(comment).RowsAffected
	if affected == 0 {
		return nil, false
	}
	return comment, true
}

func Create(u *tuser.User, v *tvideo.Video, content string) *tcomment.Comment {
	var m sync.Mutex
	comment := tcomment.Comment{AuthorID: u.ID, VideoID: v.ID, Content: content}
	m.Lock()
	tx := database.DB.Begin()
	find := tx.Create(&comment)
	up := tx.Model(&tvideo.Video{ID: v.ID}).Update("comment_count", gorm.Expr("comment_count+?", 1))
	m.Unlock()
	if find.RowsAffected != 0 && up.RowsAffected != 0 {
		tx.Commit()
		return &comment
	}
	tx.Rollback()
	return nil
}
func Delete(c *tcomment.Comment, v *tvideo.Video) bool {
	var m sync.Mutex
	tarcomment := tcomment.Comment{}
	find := database.DB.Where("id=?", c.ID).Find(&tarcomment)
	if find.RowsAffected != 0 {
		m.Lock()
		tx := database.DB.Begin()
		up1 := tx.Unscoped().Where("id=?", c.ID).Delete(&tcomment.Comment{})
		up2 := tx.Model(&tvideo.Video{ID: v.ID}).Update("comment_count", gorm.Expr("comment_count-?", 1))
		m.Unlock()
		if up1.RowsAffected != 0 && up2.RowsAffected != 0 {
			tx.Commit()
			return true
		}
		tx.Rollback()
		return false

	}
	return false
}

func GetList(v *tvideo.Video) []tcomment.Comment {
	var comments []tcomment.Comment
	database.DB.Where("video_id=?", v.ID).Find(&comments)
	return comments
}
