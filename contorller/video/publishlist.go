package video

import (
	"BAT-douyin/dao/dvideo"
	Res "BAT-douyin/entity/res"
	"BAT-douyin/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PublishList(c *gin.Context) {
	videos := dvideo.GetAllVideos()
	videoList := utils.ConvertVideoList(videos)
	c.JSON(http.StatusOK, Res.VideoListRes{
		MyResponse: Res.MyResponse{
			StatusCode: 0,
		},
		VideoList: videoList,
	})
}
