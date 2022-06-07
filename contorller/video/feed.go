package video

import (
	"BAT-douyin/dao/dvideo"
	Res "BAT-douyin/entity/res"
	"BAT-douyin/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Feed(c *gin.Context) {
	Videos := dvideo.GetAllVideos()
	videoList := utils.ConvertVideoList(Videos)
	for i := range Videos {
		//Videos[i].PlayURL = "http://" + c.Request.Host + Videos[i].PlayURL
		fmt.Print(Videos[i], "\n")
	}
	//fmt.Print(Videos)
	c.JSON(http.StatusOK, Res.VideoListRes{
		MyResponse: Res.MyResponse{StatusCode: 0},
		NextTime:   time.Now().Unix(),
		VideoList:  videoList,
	})
}
