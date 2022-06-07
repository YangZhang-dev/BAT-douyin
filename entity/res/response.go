package Res

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//响应结构体和返回错误响应的函数

type MyResponse struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type BaseUserRes struct {
	Id              uint   `json:"id"`
	Name            string `json:"name"`
	Followcount     uint   `json:"follow_count"`
	Followercount   uint   `json:"follower_count"`
	Isfollow        bool   `json:"is_follow"`
	Avatar          string `json:"avatar"`
	Signature       string `json:"signature"`
	BackgroundImage string `json:"background_image"`
	TotalFavorited  uint   `json:"total_favorited"`
	FavoriteCount   uint   `json:"favorite_count"`
}

type BaseVideoRes struct {
	Id            uint        `json:"id"`
	Author        BaseUserRes `json:"author"`
	PlayUrl       string      `json:"play_url,omitempty"`
	Title         string      `json:"title,omitempty"`
	CoverUrl      string      `json:"cover_url,omitempty"`
	FavoriteCount uint        `json:"favorite_count,omitempty"`
	CommentCount  uint        `json:"comment_count,omitempty"`
	IsFavorite    bool        `json:"is_favorite,omitempty"`
	Time          int         `json:"time，omitempty"`
}

type UserResponse struct {
	MyResponse
	UserId uint   `json:"user_id,omitempty"`
	Token  string `json:"token,omitempty"`
}

type UserMesResponse struct {
	MyResponse
	UserMes BaseUserRes `json:"user"`
}

type VideoListRes struct {
	MyResponse
	NextTime  int64          `json:"next_time,omitempty"`
	VideoList []BaseVideoRes `json:"video_list"`
}

type BaseCommentListRes struct {
	Id         uint        `json:"id"`
	Author     BaseUserRes `json:"user"`
	Content    string      `json:"content"`
	CreateDate string      `json:"create_date"`
}

type CommentListRes struct {
	MyResponse
	Base []BaseCommentListRes `json:"comment_list"`
}

type CommentRes struct {
	MyResponse
	Id         uint        `json:"id"`
	Author     BaseUserRes `json:"user"`
	Content    string      `json:"content"`
	CreateDate string      `json:"create_date"`
}

type RelationListRes struct {
	MyResponse
	AuthorList []BaseUserRes `json:"user_list"`
}

func SendErrMessage(c *gin.Context, code int64, mes string) {
	c.JSON(http.StatusOK, MyResponse{
		StatusCode: code,
		StatusMsg:  mes,
	})
}
