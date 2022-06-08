package video

import (
	"BAT-douyin/commen"
	"BAT-douyin/dao/dvideo"
	Res "BAT-douyin/entity/res"
	"BAT-douyin/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PublishList(c *gin.Context) {
	videos := dvideo.GetAllVideos()
	videoList, _ := utils.ConvertVideoList(videos)
	c.JSON(http.StatusOK, Res.VideoListRes{
		MyResponse: Res.MyResponse{
			StatusCode: commen.Success,
		},
		VideoList: videoList,
	})
}
