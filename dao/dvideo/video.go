package dvideo

import (
	"BAT-douyin/commen"
	"BAT-douyin/dao/database"
	"BAT-douyin/model"
	"fmt"
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
