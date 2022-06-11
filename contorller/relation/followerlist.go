package relation

import (
	"BAT-douyin/commen"
	"BAT-douyin/dao/duser"
	Res "BAT-douyin/entity/res"
	"BAT-douyin/model"
	"BAT-douyin/pkg/utils"
	"BAT-douyin/pkg/utils/convert"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FollowerList(c *gin.Context) {
	//查询的目标用户
	strid := c.Query("user_id")
	id, ok := utils.ParseStr2Uint(strid)
	if !ok {
		Res.SendErrMessage(c, commen.ParseError, "pares error")
		return
	}
	taru, ok := duser.GetById(id)
	if !ok {
		Res.SendErrMessage(c, commen.UserNotExist, "User not exists")
		return
	}
	//登陆用户
	claims := new(utils.UserClaim)
	token := c.Query("token")
	if token == "" {
	} else {
		claims, ok = utils.ValidateJwt(token)
		if !ok {
			Res.SendErrMessage(c, commen.TokenError, "token is not right")
			return
		}
	}
	u, ok := duser.GetById(claims.UserId)
	if !ok {
		u = new(model.User)
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
