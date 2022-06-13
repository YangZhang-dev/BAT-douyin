package relation

import (
	"BAT-douyin/commen"
	"BAT-douyin/dao/duser"
	"BAT-douyin/dao/redis"
	Res "BAT-douyin/entity/res"
	"BAT-douyin/model/tuser"
	"BAT-douyin/pkg/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

func Action(c *gin.Context) {
	op := c.DefaultQuery("action_type", "0")
	if op == "0" {
		Res.SendErrMessage(c, commen.OptionError, "option wrong")
		return
	}
	//获取要关注的用户
	strid := c.Query("to_user_id")
	taruid, ok := utils.ParseStr2Uint(strid)
	if !ok {
		Res.SendErrMessage(c, commen.ParseError, "tarid parse error")
		return
	}
	exists := false
	taru := &tuser.User{}
	err := json.Unmarshal([]byte(redis.Redis.Get(strid)), taru)
	if err != nil {
		taru, exists = duser.GetById(taruid)
		if !exists {
			Res.SendErrMessage(c, commen.UserNotExist, "user not exists")
			return
		}
		ok = redis.Redis.Set(strconv.Itoa(int(taru.ID)), taru, 1*time.Hour)
		if !ok {
			zap.L().Error("cache user error")
		}
	}

	//获取登陆的用户
	claim, _ := utils.ValidateJwt(c.Query("token"))

	u := &tuser.User{}
	err = json.Unmarshal([]byte(redis.Redis.Get(strid)), u)
	if err != nil {
		u, exists = duser.GetById(claim.UserId)
		if !exists {
			Res.SendErrMessage(c, commen.UserNotExist, "user not exists")
			return
		}
		ok = redis.Redis.Set(strconv.Itoa(int(u.ID)), u, 1*time.Hour)
		if !ok {
			zap.L().Error("cache user error")
		}
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
