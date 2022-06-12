package video

import (
	"BAT-douyin/commen"
	"BAT-douyin/dao/duser"
	"BAT-douyin/dao/dvideo"
	Res "BAT-douyin/entity/res"
	"BAT-douyin/model"
	"BAT-douyin/pkg/utils"
	"BAT-douyin/pkg/utils/convert"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func Feed(c *gin.Context) {
	claim := new(utils.UserClaim)
	var ok bool
	token := c.DefaultQuery("token", "")
	if token == "" {
	} else {
		claim, ok = utils.ValidateJwt(token)
		if !ok {
			Res.SendErrMessage(c, commen.TokenError, "token is wrong")
			return
		}
	}
	u, ok := duser.GetById(claim.UserId)
	if !ok {
		u = new(model.User)
	}

	t, err := strconv.ParseInt(c.Query("next_time"), 10, 64)
	if err != nil {
		Res.SendErrMessage(c, commen.ParseError, "parse error")
		return
	}
	objtime := time.Unix(t, 0)
	Videos := dvideo.GetAllVideos(nil, objtime)

	videoList, err := convert.ConvertVideoList(Videos, u)
	if err != nil {
		Res.SendErrMessage(c, commen.ParseError, "error occurred when parsing videoList")
		return
	}
	nextTime := time.Now().Unix()
	if len(videoList) != 0 {
		nextTime = Videos[0].CreatedAt.Unix()
	}

	c.JSON(http.StatusOK, Res.VideoListRes{
		MyResponse: Res.MyResponse{StatusCode: 0},
		NextTime:   nextTime,
		VideoList:  videoList,
	})
}
