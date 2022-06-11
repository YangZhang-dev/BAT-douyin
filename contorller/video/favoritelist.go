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
)

func FavoriteList(c *gin.Context) {
	claim := new(utils.UserClaim)
	var ok bool
	token := c.Query("token")
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

	//查询目标用户
	strid := c.Query("user_id")
	uid, ok := utils.ParseStr2Uint(strid)
	if !ok {
		Res.SendErrMessage(c, commen.ParseError, "pares error")
		return
	}
	taru, ok := duser.GetById(uid)
	if !ok {
		u = new(model.User)
	}
	videos := dvideo.UserFavoriteVideos(taru)
	videoList, err := convert.ConvertVideoList(videos, u)
	if err != nil {
		Res.SendErrMessage(c, commen.ParseError, "error occurred when parsing videoList")
		return
	}
	c.JSON(http.StatusOK, Res.VideoListRes{
		MyResponse: Res.MyResponse{
			StatusCode: commen.Success,
			StatusMsg:  "Get user favorite list success",
		},
		VideoList: videoList,
	})
}
