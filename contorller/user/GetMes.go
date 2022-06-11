package user

import (
	"BAT-douyin/commen"
	"BAT-douyin/dao/duser"
	Res "BAT-douyin/entity/res"
	"BAT-douyin/model"
	"BAT-douyin/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMes(c *gin.Context) {

	token := c.DefaultQuery("token", "")

	strid := c.DefaultQuery("user_id", "0")
	id, ok := utils.ParseStr2Uint(strid)
	if !ok {
		Res.SendErrMessage(c, commen.ParseError, "pares error")
		return
	}
	//尝试通过Token获取登陆用户
	u := duser.GetByToken(token)
	if u == nil {
		u = new(model.User)
	}
	//通过id获取要查看的目标用户
	taru, exists := duser.GetById(id)
	if !exists {
		Res.SendErrMessage(c, commen.UserNotExist, "user not exists")
		return
	}

	//查询登陆用户是否关注目标用户
	is := duser.IsFollowUser(u, taru)
	c.JSON(http.StatusOK, Res.UserMesResponse{
		MyResponse: Res.MyResponse{
			StatusCode: commen.Success,
			StatusMsg:  "Get user message success",
		},
		UserMes: Res.BaseUserRes{
			Id:              taru.ID,
			Name:            taru.UserName,
			Followcount:     taru.FollowCount,
			Followercount:   taru.FollowerCount,
			Isfollow:        is,
			Avatar:          taru.Avatar,
			FavoriteCount:   taru.FavoriteCount,
			TotalFavorited:  taru.TotalFavorited,
			BackgroundImage: taru.BackgroundImage,
			Signature:       taru.Signature,
		},
	})
}
