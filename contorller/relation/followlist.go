package relation

import (
	"BAT-douyin/commen"
	"BAT-douyin/dao/duser"
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

func FollowList(c *gin.Context) {
	//查询的目标用户
	strid := c.Query("user_id")
	id, ok := utils.ParseStr2Uint(strid)
	if !ok {
		Res.SendErrMessage(c, commen.ParseError, "pares error")
		return
	}
	exists := false
	taru := &model.User{}
	err := json.Unmarshal([]byte(redis.Redis.Get(strid)), taru)
	if err != nil {
		taru, exists = duser.GetById(id)
		if !exists {
			Res.SendErrMessage(c, commen.UserNotExist, "user not exists")
			return
		}
		ok = redis.Redis.Set(strconv.Itoa(int(taru.ID)), taru, 1*time.Hour)
		if !ok {
			zap.L().Error("cache user error")
		}
	}
	//登陆用户
	claim := new(utils.UserClaim)
	token := c.Query("token")
	if token == "" {
	} else {
		claim, ok = utils.ValidateJwt(token)
		if !ok {
			Res.SendErrMessage(c, commen.TokenError, "token is not right")
			return
		}
	}
	exists = false
	u := &model.User{}
	err = json.Unmarshal([]byte(redis.Redis.Get(strconv.Itoa(int(claim.UserId)))), u)
	if err != nil {
		u, exists = duser.GetById(claim.UserId)
		if !exists {
			u = new(model.User)
		}
	}
	userlist := duser.FollowUserList(taru)
	userList, err := convert.ConvertUserListRes(userlist, u, 1)
	if err != nil {
		Res.SendErrMessage(c, commen.ParseError, "Get userRes wrong")
		return
	}
	c.JSON(http.StatusOK, Res.RelationListRes{
		MyResponse: Res.MyResponse{
			StatusCode: commen.Success,
			StatusMsg:  "Action success",
		},
		AuthorList: userList,
	})
}
