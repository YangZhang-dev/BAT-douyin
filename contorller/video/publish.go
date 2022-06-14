package video

import (
	"BAT-douyin/commen"
	"BAT-douyin/cos"
	"BAT-douyin/dao/duser"
	"BAT-douyin/dao/dvideo"
	"BAT-douyin/dao/redis"
	Res "BAT-douyin/entity/res"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
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
	//videoPath := filepath.Join("./static/video/", finalVideoName)
	file, err := (*data).Open()
	if err != nil {
		Res.SendErrMessage(c, commen.SaveVideoError, "save video files error")
		return
	}

	err = cos.PutVideo2Cos(cos.CosClient, baseFinalName, file)
	if err != nil {
		Res.SendErrMessage(c, commen.SaveVideoError, "save video to cos wrong")
		return
	}
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
