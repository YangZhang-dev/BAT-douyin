package dvideo

import (
	"BAT-douyin/commen"
	"BAT-douyin/dao/database"
	"BAT-douyin/model"
	"fmt"
	"gorm.io/gorm"
	"sync"
	"time"
)

func GetAllVideos(u *model.User, time time.Time) []*model.Video {
	var videos []*model.Video
	if u != nil {
		database.DB.Order("created_at").Limit(30).Where("author_id=?", u.ID).Find(&videos)
	} else {
		database.DB.Order("created_at").Limit(30).Where("created_at>?", time).Find(&videos)
	}
	return videos
}

func Save(path string, userid uint, title string) (*model.Video, bool) {
	m := sync.Mutex{}
	playurl := fmt.Sprintf("%s/static/video/%s.mp4", commen.SelfIp, path)
	coverurl := fmt.Sprintf("%s/static/cover/%s.jpg", commen.SelfIp, path)
	video := model.Video{PlayUrl: playurl, AuthorID: userid, CoverUrl: coverurl, CommentCount: 0, FavoriteCount: 0, Title: title}
	m.Lock()
	create := database.DB.Create(&video)
	m.Unlock()
	if create.RowsAffected == 0 {
		return nil, false
	}
	return &video, true
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
		tx := database.DB.Begin()
		create := tx.Where("user_id=? and video_id=?", uid, vid).Create(&model.FavoriteVideo{UserID: uid, VideoID: vid})
		up1 := tx.Model(&model.Video{ID: vid}).Update("favorite_count", gorm.Expr("favorite_count+?", 1))
		up2 := tx.Model(&model.User{ID: video.AuthorID}).Update("total_favorited", gorm.Expr("total_favorited+?", 1))
		up3 := tx.Model(&model.User{ID: uid}).Update("favorite_count", gorm.Expr("favorite_count+?", 1))
		m.Unlock()
		if up1.RowsAffected != 0 && up2.RowsAffected != 0 && up3.RowsAffected != 0 && create.RowsAffected != 0 {
			tx.Commit()
			return true
		}
		tx.Rollback()
		return false
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
		tx := database.DB.Begin()
		up := tx.Unscoped().Where("user_id=? and video_id=?", uid, vid).Delete(&model.FavoriteVideo{})
		up1 := tx.Model(&model.Video{ID: vid}).Update("favorite_count", gorm.Expr("favorite_count-?", 1))
		up2 := tx.Model(&model.User{ID: video.AuthorID}).Update("total_favorited", gorm.Expr("total_favorited-?", 1))
		up3 := tx.Model(&model.User{ID: uid}).Update("favorite_count", gorm.Expr("favorite_count-?", 1))
		m.Unlock()
		if up1.RowsAffected != 0 && up2.RowsAffected != 0 && up3.RowsAffected != 0 && up.RowsAffected != 0 {
			tx.Commit()
			return true
		}
		tx.Rollback()
		return false
	}
}

func GetById(id uint) (*model.Video, bool) {
	video := &model.Video{}
	affected := database.DB.Where("id=?", id).Find(video).RowsAffected
	if affected == 0 {
		return nil, false
	}
	return video, true
}

func IsFavoriteVideo(u *model.User, v *model.Video) bool {
	affected := database.DB.Where("user_id=? and video_id=?", u.ID, v.ID).Find(&model.FavoriteVideo{}).RowsAffected
	if affected == 0 {
		return false
	}
	return true
}

func UserFavoriteVideos(u *model.User) []*model.Video {
	var tmp []model.FavoriteVideo
	var videolist []*model.Video
	find := database.DB.Where("user_id=?", u.ID).Find(&tmp)
	if find.RowsAffected == 0 {
		return nil
	}
	for i := 0; i < len(tmp); i++ {
		video, ok := GetById(tmp[i].VideoID)
		if !ok {
			return nil
		}
		videolist = append(videolist, video)
	}
	return videolist
}
