package routes

import (
	"BAT-douyin/contorller/comment"
	"BAT-douyin/contorller/relation"
	"BAT-douyin/contorller/user"
	"BAT-douyin/contorller/video"
	"BAT-douyin/middlewares"
	"BAT-douyin/middlewares/logger"
	"BAT-douyin/setting"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Run(conf *setting.AppConfig) error {
	if conf.Mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // 如果配置文件的model设置为release则gin设置成发布模式，可以去除一些不必要的控制台输出
	}
	//使用Gin默认的中间件
	//r := gin.Default()

	//自定义中间件
	r := gin.New()
	r.Use(logger.Ginlogger(), logger.GinRecovery(true))
	//注册路由
	RegisterRouter(r)

	err := r.Run(fmt.Sprintf(":%v", conf.Port))
	if err != nil {
		return err
	}
	return nil
}

func RegisterRouter(r *gin.Engine) {

	douyin := r.Group("/douyin")
	{
		douyin.GET("/feed", video.Feed)
		u := douyin.Group("/user")
		{
			u.GET("/", user.GetMes)
			u.POST("/register/", user.Register)
			u.POST("/login/", user.Login)
		}
		public := douyin.Group("/publish")
		{
			public.GET("/list/", video.PublishList)
			public.POST("/action/", middlewares.Jwt(), video.Publish)
		}
		favorite := douyin.Group("/favorite")
		{
			favorite.POST("/action/", middlewares.Jwt(), user.FavoriteVideo)
			favorite.GET("/list/", video.FavoriteList)
		}
		comm := douyin.Group("/comment")
		{
			comm.POST("/action/", middlewares.Jwt(), comment.Post)
			comm.GET("/list/", comment.List)
		}
		rela := douyin.Group("/relation")
		{
			rela.POST("/action/", middlewares.Jwt(), relation.Action)
			rela.GET("follow/list/", relation.FollowList)
			rela.GET("follower/list/", relation.FollowerList)
		}
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404NotFound",
		})
	})

}
