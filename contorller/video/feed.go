package video

import (
	"BAT-douyin/commen"
	"BAT-douyin/dao/dvideo"
	Res "BAT-douyin/entity/res"
	"BAT-douyin/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Feed(c *gin.Context) {
	Videos := dvideo.GetAllVideos()
	videoList, err := utils.ConvertVideoList(Videos)
	if err != nil {
		Res.SendErrMessage(c, commen.ParseError, "error occurred when parsing videoList")
		return
	}
	//下面注释掉的部分可以用来打印videos的信息
	//for i := range Videos {
	//	//Videos[i].PlayURL = "http://" + c.Request.Host + Videos[i].PlayURL
	//	fmt.Print(Videos[i], "\n")
	//}
	//fmt.Print(Videos)
	c.JSON(http.StatusOK, Res.VideoListRes{
		MyResponse: Res.MyResponse{StatusCode: 0},
		NextTime:   time.Now().Unix(),
		VideoList:  videoList,
	})
}
