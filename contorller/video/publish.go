package video

import (
	"BAT-douyin/commen"
	"BAT-douyin/dao/duser"
	"BAT-douyin/dao/dvideo"
	Res "BAT-douyin/entity/res"
	"BAT-douyin/pkg/utils"
	"BAT-douyin/redis"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
)

func Publish(c *gin.Context) {
	token := c.PostForm("token")

	title := c.PostForm("title")

	user := duser.GetByToken(token)

	if user == nil {
		Res.SendErrMessage(c, commen.UserNotExist, "user does not exists")
		return
	}

	data, err2 := c.FormFile("data")
	if err2 != nil {
		Res.SendErrMessage(c, commen.GetVideoError, "video error")
		return
	}

	baseFinalName := fmt.Sprintf("%d_%d", user.ID, time.Now().Unix())
	finalVideoName := fmt.Sprintf("%s.mp4", baseFinalName)
	videoPath := filepath.Join("./static/video/", finalVideoName)
	if err := c.SaveUploadedFile(data, videoPath); err != nil {
		Res.SendErrMessage(c, commen.SaveVideoError, "error occurred when saving video files")
		return
	}

	go utils.GetCover(videoPath, baseFinalName)

	v, ok := dvideo.Save(baseFinalName, user.ID, title)
	if !ok {
		Res.SendErrMessage(c, commen.SaveVideoError, "error occurred when saving video files")
		return
	}

	ok = redis.Redis.Set("v_"+strconv.Itoa(int(v.ID)), v, 1*time.Hour)
	if !ok {
		zap.L().Error("cache video error")
	}

	c.JSON(http.StatusOK, Res.MyResponse{
		StatusCode: commen.Success,
		StatusMsg:  finalVideoName + " uploaded successfully",
	})
}
