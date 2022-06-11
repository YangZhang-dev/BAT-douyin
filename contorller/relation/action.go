package relation

import (
	"BAT-douyin/commen"
	"BAT-douyin/dao/duser"
	Res "BAT-douyin/entity/res"
	"BAT-douyin/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Action(c *gin.Context) {
	op := c.DefaultQuery("action_type", "0")
	if op == "0" {
		Res.SendErrMessage(c, commen.OptionError, "option wrong")
		return
	}
	//获取要关注的用户
	taruid, ok := utils.ParseStr2Uint(c.Query("to_user_id"))
	if !ok {
		Res.SendErrMessage(c, commen.ParseError, "tarid parse error")
		return
	}
	taru, ok := duser.GetById(taruid)
	if !ok {
		Res.SendErrMessage(c, commen.UserNotExist, "user not exists")
		return
	}

	//获取登陆的用户
	claim, _ := utils.ValidateJwt(c.Query("token"))
	u, ok := duser.GetById(claim.UserId)
	if !ok {
		Res.SendErrMessage(c, commen.UserNotExist, "user not exists")
		return
	}

	//不允许自己关注自己
	if u.ID == taruid {
		Res.SendErrMessage(c, commen.CanNotFollowSelf, "user cant not follow self")
		return
	}
	if op == "1" {
		ok := duser.FollowUser(u, taru)
		if !ok {
			Res.SendErrMessage(c, commen.FollowUserWrong, "Follow user wrong")
			return
		}
	} else {
		//取消关注
		ok = duser.UnFollowUser(u, taru)
		if !ok {
			Res.SendErrMessage(c, commen.UnFollowUserWrong, "Follow user wrong")
			return
		}
	}

	c.JSON(http.StatusOK, Res.MyResponse{
		StatusCode: commen.Success,
		StatusMsg:  "action success",
	})
}
