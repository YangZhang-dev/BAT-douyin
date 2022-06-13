package video

import (
	"BAT-douyin/commen"
	"BAT-douyin/dao/duser"
	"BAT-douyin/dao/dvideo"
	Res "BAT-douyin/entity/res"
	"BAT-douyin/model"
	"BAT-douyin/pkg/utils"
	"BAT-douyin/pkg/utils/convert"
	"BAT-douyin/redis"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

func PublishList(c *gin.Context) {
	//登陆用户
	claim := new(utils.UserClaim)
	var ok bool
	token := c.Query("token")
	if token == "" {
	} else {
		claim, ok = utils.ValidateJwt(token)
		if !ok {
			Res.SendErrMessage(c, commen.TokenError, "token is not right")
			return
		}
	}
	exists := false
	u := &model.User{}
	err := json.Unmarshal([]byte(redis.Redis.Get(claim.Id)), u)
	if err != nil {
		u, exists = duser.GetById(claim.UserId)
		if !exists {
			u = new(model.User)
		}
	}

	//查询目标用户
	struid := c.Query("user_id")
	uid, ok := utils.ParseStr2Uint(struid)
	if !ok {
		Res.SendErrMessage(c, commen.ParseError, "pares error")
		return
	}
	exists = false
	taru := &model.User{}
	err = json.Unmarshal([]byte(redis.Redis.Get(struid)), taru)
	if err != nil {
		taru, exists = duser.GetById(uid)
		if !exists {
			Res.SendErrMessage(c, commen.UserNotExist, "user not exists")
			return
		}
		ok = redis.Redis.Set(strconv.Itoa(int(taru.ID)), taru, 1*time.Hour)
		if !ok {
			zap.L().Error("cache user error")
		}
	}

	videos := dvideo.GetAllVideos(taru, time.Now())
	videoList, err := convert.ConvertVideoList(videos, u)
	if err != nil {
		Res.SendErrMessage(c, commen.ParseError, "error occurred when parsing videoList")
		return
	}
	c.JSON(http.StatusOK, Res.VideoListRes{
		MyResponse: Res.MyResponse{
			StatusCode: commen.Success,
			StatusMsg:  "Get user publish video list success",
		},
		VideoList: videoList,
	})
}
