package comment

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

func List(c *gin.Context) {
	claims := new(utils.UserClaim)
	var ok bool
	token := c.Query("token")
	if token == "" {
	} else {
		claims, ok = utils.ValidateJwt(token)
		if !ok {
			Res.SendErrMessage(c, commen.TokenError, "token is not right")
			return
		}

	}
	u := &model.User{}
	err := json.Unmarshal([]byte(redis.Redis.Get(strconv.Itoa(int(claims.UserId)))), u)
	if err != nil {
		u, ok = duser.GetById(claims.UserId)
		ok = redis.Redis.Set(strconv.Itoa(int(u.ID)), u, 1*time.Hour)
		if !ok {
			zap.L().Error("cache user error")
		}
	}

	strvid := c.Query("video_id")
	vid, ok := utils.ParseStr2Uint(strvid)
	if !ok {
		Res.SendErrMessage(c, commen.ParseError, "vid pares error")
		return
	}
	v := &model.Video{}
	err = json.Unmarshal([]byte(redis.Redis.Get(strvid)), v)
	if err != nil {
		v, ok = dvideo.GetById(vid)
		if !ok {
			Res.SendErrMessage(c, commen.UserNotExist, "please login")
			return
		}
		ok = redis.Redis.Set(strconv.Itoa(int(v.ID)), u, 1*time.Hour)
		if !ok {
			zap.L().Error("cache video error")
		}
	}

	commentlist, err := convert.GetCommentList(v, u)
	if err != nil {
		Res.SendErrMessage(c, commen.ParseError, "Get comment list error")
		return
	}
	c.JSON(http.StatusOK, Res.CommentListRes{
		MyResponse: Res.MyResponse{
			StatusCode: commen.Success,
			StatusMsg:  "success get comment list ",
		},
		Base: commentlist,
	})
}
