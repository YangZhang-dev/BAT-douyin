package user

import (
	"BAT-douyin/commen"
	"BAT-douyin/dao/duser"
	"BAT-douyin/dao/dvideo"
	Res "BAT-douyin/entity/res"
	"BAT-douyin/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FavoriteVideo(c *gin.Context) {
	op := c.DefaultQuery("action_type", "0")
	if op == "0" {
		Res.SendErrMessage(c, commen.OptionError, "option wrong")
		return
	}
	strid := c.DefaultQuery("video_id", "0")
	vid, ok := utils.ParseStr2Uint(strid)
	if !ok {
		Res.SendErrMessage(c, commen.ParseError, "pares error")
		return
	}

	token := c.DefaultQuery("token", "")
	//尝试通过Token获取用户
	u := duser.GetByToken(token)
	if u == nil {
		Res.SendErrMessage(c, commen.UserNotExist, "please login")
		return
	}

	v, ok := dvideo.GetById(vid)
	if !ok {
		Res.SendErrMessage(c, commen.UserNotExist, "please login")
		return
	}
	if v.AuthorID == u.ID {
		Res.SendErrMessage(c, commen.UserCannotLikeSelfVideo, "Use can not favorite self video")
		return

	}

	if op == "1" {
		ok = dvideo.LikeVideo(u, v)
		if !ok {
			Res.SendErrMessage(c, commen.UnlikeVideoWrong, "Use favorite video wrong")
			return
		}
	} else {
		ok := dvideo.UnlikeVideo(u, v)
		if !ok {
			Res.SendErrMessage(c, commen.UserCannotLikeSelfVideo, "Use  unlike video wrong")
			return
		}
	}
	c.JSON(http.StatusOK, Res.MyResponse{
		StatusCode: commen.Success,
		StatusMsg:  "Get user message success",
	})
}
