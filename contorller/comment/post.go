package comment

import (
	"BAT-douyin/commen"
	"BAT-douyin/dao/dcomment"
	"BAT-douyin/dao/duser"
	"BAT-douyin/dao/dvideo"
	Res "BAT-douyin/entity/res"
	"BAT-douyin/pkg/utils"
	"BAT-douyin/redis"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

func Post(c *gin.Context) {
	op := c.DefaultQuery("action_type", "0")

	strvid := c.DefaultQuery("video_id", "0")
	vid, ok := utils.ParseStr2Uint(strvid)
	if !ok {
		Res.SendErrMessage(c, commen.ParseError, "vid pares error")
		return
	}
	v, exists := dvideo.GetById(vid)
	if !exists {
		Res.SendErrMessage(c, commen.VideoNotExists, "Video Not exists")
		return
	}
	author, exists := duser.GetById(v.AuthorID)
	if !exists {
		Res.SendErrMessage(c, commen.UserNotExist, "User not exists")
		return
	}

	claim := new(utils.UserClaim)
	token := c.DefaultQuery("token", "")
	if token == "" {
	} else {
		claim, ok = utils.ValidateJwt(token)
		if !ok {
			Res.SendErrMessage(c, commen.TokenError, "token is wrong")
			return
		}
	}
	u, ok := duser.GetById(claim.UserId)
	if !ok {
		Res.SendErrMessage(c, commen.ParseError, "uid pares error")
		return
	}

	//发布评论时存在
	text := c.DefaultQuery("comment_text", "")

	//发布评论
	if op == "1" {
		com := dcomment.Create(u, v, text)

		if com == nil {
			Res.SendErrMessage(c, commen.PostCommentWrong, "Post comment wrong")
			return
		}
		ok := redis.Redis.Set(strconv.Itoa(int(com.ID)), com, 1*time.Hour)
		if !ok {
			zap.L().Error("cache  video error")
		}
		is := duser.IsFollowUser(u, author)
		c.JSON(http.StatusOK, Res.CommentRes{
			MyResponse: Res.MyResponse{
				StatusCode: commen.Success,
				StatusMsg:  "Post comment success",
			},
			Id:         com.ID,
			Author:     Res.BaseUserRes{Id: u.ID, Name: u.UserName, Followcount: u.FollowCount, Followercount: u.FollowerCount, Isfollow: is, Avatar: u.Avatar, Signature: u.Signature, BackgroundImage: u.BackgroundImage, TotalFavorited: u.TotalFavorited, FavoriteCount: u.FavoriteCount},
			Content:    text,
			CreateDate: com.CreatedAt.Format("01-02 15:04:05"),
		})

	} else {
		//删除评论

		//删除评论时存在
		strcid := c.DefaultQuery("comment_id", "0")
		cid, ok := utils.ParseStr2Uint(strcid)
		if !ok {
			Res.SendErrMessage(c, commen.ParseError, "vid pares error")
			return
		}
		comment, exist := dcomment.GetById(cid)
		if !exist {
			Res.SendErrMessage(c, commen.CommentNotExists, "Comment not exists")
			return
		}
		ok = redis.Redis.Del(strconv.Itoa(int(comment.ID)))
		if !ok {
			zap.L().Error("delete cache of video error")
		}
		ok = dcomment.Delete(comment, v)
		if !ok {
			Res.SendErrMessage(c, commen.DeleteCommentWrong, "Delete comment wrong")
			return
		}
		c.JSON(http.StatusOK, Res.MyResponse{
			StatusCode: commen.Success,
			StatusMsg:  "Delete comment success",
		})
	}

}
