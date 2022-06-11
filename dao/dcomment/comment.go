package dcomment

import (
	"BAT-douyin/dao/database"
	"BAT-douyin/model"
	"gorm.io/gorm"
	"sync"
)

func GetById(cid uint) (*model.Comment, bool) {
	if cid == 0 {
		return nil, false
	}
	comment := &model.Comment{}
	affected := database.DB.Where("id=?", cid).Find(comment).RowsAffected
	if affected == 0 {
		return nil, false
	}
	return comment, true
}

func Create(u *model.User, v *model.Video, content string) *model.Comment {
	var m sync.Mutex
	comment := model.Comment{AuthorID: u.ID, VideoID: v.ID, Content: content}
	m.Lock()
	tx := database.DB.Begin()
	find := tx.Create(&comment)
	up := tx.Model(&model.Video{ID: v.ID}).Update("comment_count", gorm.Expr("comment_count+?", 1))
	m.Unlock()
	if find.RowsAffected != 0 && up.RowsAffected != 0 {
		tx.Commit()
		return &comment
	}
	tx.Rollback()
	return nil
}
func Delete(c *model.Comment, v *model.Video) bool {
	var m sync.Mutex
	tarcomment := model.Comment{}
	find := database.DB.Where("id=?", c.ID).Find(&tarcomment)
	if find.RowsAffected != 0 {
		m.Lock()
		tx := database.DB.Begin()
		up1 := tx.Unscoped().Where("id=?", c.ID).Delete(&model.Comment{})
		up2 := tx.Model(&model.Video{ID: v.ID}).Update("comment_count", gorm.Expr("comment_count-?", 1))
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

func GetList(v *model.Video) []model.Comment {
	var comments []model.Comment
	database.DB.Where("video_id=?", v.ID).Find(&comments)
	return comments
}
