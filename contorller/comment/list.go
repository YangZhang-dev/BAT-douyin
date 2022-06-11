package comment

import (
	"BAT-douyin/commen"
	"BAT-douyin/dao/duser"
	"BAT-douyin/dao/dvideo"
	Res "BAT-douyin/entity/res"
	"BAT-douyin/pkg/utils"
	"BAT-douyin/pkg/utils/convert"
	"github.com/gin-gonic/gin"
	"net/http"
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
	u, ok := duser.GetById(claims.UserId)

	strvid := c.Query("video_id")
	vid, ok := utils.ParseStr2Uint(strvid)
	if !ok {
		Res.SendErrMessage(c, commen.ParseError, "vid pares error")
		return
	}
	v, exists := dvideo.GetById(vid)
	if !exists {
		Res.SendErrMessage(c, commen.VideoNotExists, "user not exists")
		return
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
