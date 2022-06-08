package commen

const Success = 0

//video错误
const (
	//样例
	GetVideoError  = 20001
	SaveVideoError = 20002
)

//user错误
const (
	UserNotExistsError = 30001
)

//评论错误
const (
	CommentError = 40001
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
	TokenError = 70001
	ParseError = 80001
)

//由于没有开放个人信息编辑功能，所以直接写死，可以自定义图片
const (
	SelfIp          = "http://127.0.0.1:8080"
	Signature       = "我是无敌暴龙战神"
	BackgroundImage = SelfIp + "/static/backgroundimage/default.jpg"
	Avatar          = SelfIp + "/static/avatar/default.jpg"
)
