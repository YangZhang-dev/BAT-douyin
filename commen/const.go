package commen

const Success = 0

//video错误
const (
	//样例
	GetVideoError           = 20001
	SaveVideoError          = 20002
	UserCannotLikeSelfVideo = 20003
	LikeVideoWrong          = 20004
	UnlikeVideoWrong        = 20005
	VideoNotExists          = 20007
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
	CanNotFollowSelf         = 30009
	FollowUserWrong          = 30010
	UnFollowUserWrong        = 30011
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
	FollowError = 50001
)

//点赞错误
const (
	FavoriteError = 60001
)

//token和parse错误
const (
	TokenError           = 70001
	ParseError           = 80001
	PasswordEncryptError = 80002
	GetTokenError        = 70002
)

//由于没有开放个人信息编辑功能，所以直接写死，可以自定义图片
const (
	SelfIp          = "http://127.0.0.1:8080"
	Signature       = "我是无敌暴龙战神"
	BackgroundImage = SelfIp + "/static/backgroundimage/default.jpg"
	Avatar          = SelfIp + "/static/avatar/default.jpg"
)
