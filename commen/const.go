package commen

import (
	"BAT-douyin/setting"
	"errors"
	"fmt"
)

const Success = 0

//video错误
const (
	//样例
	GetVideoError  = 20001
	SaveVideoError = 20002
	VideoNotExists = 20003
)

//user错误
const (
	UserError                = 30001
	UserAlreadyExist         = 30002
	RegisterFailed           = 30003
	UserNotExist             = 30004
	UserCreatError           = 30005
	UserPasswordMistake      = 30006
	UsernameOrPasswordIsNull = 30007
	OptionError              = 30008
)

//评论错误
const (
	CommentError       = 40001
	PostCommentWrong   = 40002
	CommentNotExists   = 40003
	DeleteCommentWrong = 40004
)

//关注错误
const (
	FollowError       = 50001
	UnFollowUserWrong = 50002
	FollowUserWrong   = 50003
	CanNotFollowSelf  = 50004
)

//点赞错误
const (
	FavoriteError           = 60001
	LikeVideoWrong          = 60002
	UnlikeVideoWrong        = 60003
	UserCannotLikeSelfVideo = 60004
)

//token和parse错误
const (
	TokenError           = 70001
	GetTokenError        = 70002
	ParseError           = 80001
	PasswordEncryptError = 80002
)

//由于没有开放个人信息编辑功能，所以直接写死，可以自定义图片
var (
	CosDNS          string
	Signature       = "我是无敌暴龙战神"
	BackgroundImage string
	Avatar          string
)

func Init(conf *setting.CosConfig) error {
	CosDNS = conf.Url
	BackgroundImage = fmt.Sprintf("%s/backgroundimage/default.jpg", CosDNS)
	Avatar = fmt.Sprintf("%s/avatar/default.jpg", CosDNS)
	if CosDNS == "" {
		return errors.New("init const wrong")
	}
	return nil
}
