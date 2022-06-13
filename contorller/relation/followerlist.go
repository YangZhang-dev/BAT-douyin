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
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func FollowerList(c *gin.Context) {
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

	userlist := duser.FollowerUserList(taru)
	fmt.Println(userlist)
	userList, err := convert.ConvertUserListRes(userlist, u, 2)
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
