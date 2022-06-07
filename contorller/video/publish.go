package video

import (
	"BAT-douyin/dao/database"
	"BAT-douyin/dao/dvideo"
	Res "BAT-douyin/entity/res"
	"BAT-douyin/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"time"
)

func Publish(c *gin.Context) {
	token := c.PostForm("token")

	title := c.PostForm("title")

	user := model.User{}

	if err := database.DB.First(&user, "token=?", token).Error; err != nil {
		c.JSON(http.StatusOK, Res.MyResponse{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	data, err2 := c.FormFile("data")
	if err2 != nil {
		c.JSON(http.StatusOK, Res.MyResponse{
			StatusCode: 1,
			StatusMsg:  err2.Error(),
		})
		return
	}

	//filename := filepath.Base(data.Filename)
	//user := usersLoginInfo[token]
	baseFinalName := fmt.Sprintf("%d_%d", user.ID, time.Now().Unix())
	finalVideoName := fmt.Sprintf("%s.mp4", baseFinalName)
	videoPath := filepath.Join("./static/video/", finalVideoName)
	if err := c.SaveUploadedFile(data, videoPath); err != nil {
		c.JSON(http.StatusOK, Res.MyResponse{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	dvideo.Save(baseFinalName, user.ID, title)

	c.JSON(http.StatusOK, Res.MyResponse{
		StatusCode: 0,
		StatusMsg:  finalVideoName + " uploaded successfully",
	})
}
