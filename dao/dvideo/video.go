package dvideo

import (
	"BAT-douyin/commen"
	"BAT-douyin/dao/database"
	"BAT-douyin/model"
	"fmt"
	"gorm.io/gorm"
	"sync"
)

func GetAllVideos() []model.Video {
	var Videos []model.Video
	//var video model.Video
	////如果没有创建表，就创建出这个video表 这个操作在dbInit的时候做了，注释掉。
	//database.DB.AutoMigrate(&video)
	database.DB.Limit(30).Find(&Videos)
	return Videos
}

func Save(path string, userid uint, title string) bool {
	m := sync.Mutex{}
	//下面的127.0.0.1应该换成自己的服务器的ip地址
	playurl := fmt.Sprintf("http://%s:8080/static/video/%s.mp4", commen.SelfIp, path)
	coverurl := fmt.Sprintf("http://%s:8080/static/cover/%s.jpg", commen.SelfIp, path)
	video := model.Video{PlayUrl: playurl, AuthorID: userid, CoverUrl: coverurl, CommentCount: 0, FavoriteCount: 0, Title: title}
	m.Lock()
	create := database.DB.Create(&video)
	m.Unlock()
	if create.RowsAffected == 0 {
		return false
	}
	return true
}

func LikeVideo(u *model.User, video *model.Video) bool {
	var m sync.Mutex
	uid := u.ID
	vid := video.ID
	tarVideo := model.FavoriteVideo{}
	find := database.DB.Where("user_id=? and video_id=?", uid, vid).Find(&tarVideo)
	if find.RowsAffected != 0 {
		return false
	} else {
		m.Lock()
		txx := database.DB.Begin()
		create := txx.Where("user_id=? and video_id=?", uid, vid).Create(&model.FavoriteVideo{UserID: uid, VideoID: vid})

		up1 := txx.Model(&model.Video{ID: vid}).Update("favorite_count", gorm.Expr("favorite_count+?", 1))
		if up1.RowsAffected == 0 {
			txx.Rollback()
			return false
		}
		up2 := txx.Model(&model.User{ID: video.AuthorID}).Update("total_favorited", gorm.Expr("total_favorited+?", 1))
		if up2.RowsAffected == 0 {
			txx.Rollback()
			return false
		}
		up3 := txx.Model(&model.User{ID: uid}).Update("favorite_count", gorm.Expr("favorite_count+?", 1))
		if up3.RowsAffected == 0 {
			txx.Rollback()
			return false
		}
		if create.RowsAffected == 0 {
			txx.Rollback()
			return false
		}
		m.Unlock()
		txx.Commit()
		return true
	}
}

func UnlikeVideo(u *model.User, video *model.Video) bool {
	var m sync.Mutex
	uid := u.ID
	vid := video.ID
	tarVideo := model.FavoriteVideo{}
	find := database.DB.Where("user_id=? and video_id=?", uid, vid).Find(&tarVideo)
	if find.RowsAffected == 0 {
		return false
	} else {
		m.Lock()
		txx := database.DB.Begin()
		tx := txx.Unscoped().Where("user_id=? and video_id=?", uid, vid).Delete(&model.FavoriteVideo{})
		if tx.RowsAffected == 0 {
			txx.Commit()
			return false
		}
		up1 := txx.Model(&model.Video{ID: vid}).Update("favorite_count", gorm.Expr("favorite_count-?", 1))
		if up1.RowsAffected == 0 {
			txx.Rollback()
			return false
		}
		up2 := txx.Model(&model.User{ID: video.AuthorID}).Update("total_favorited", gorm.Expr("total_favorited-?", 1))
		if up2.RowsAffected == 0 {
			txx.Rollback()
			return false
		}
		up3 := txx.Model(&model.User{ID: uid}).Update("favorite_count", gorm.Expr("favorite_count-?", 1))
		if up3.RowsAffected == 0 {
			txx.Rollback()
			return false
		}
		m.Unlock()
		txx.Commit()
		return true
	}
}
