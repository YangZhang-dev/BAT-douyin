package user

import (
	"BAT-douyin/commen"
	"BAT-douyin/dao/duser"
	Res "BAT-douyin/entity/res"
	"BAT-douyin/pkg/utils"
	"BAT-douyin/redis"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"time"
)

func Register(c *gin.Context) {

	username := c.DefaultQuery("username", "")
	password := c.DefaultQuery("password", "")
	if username == "" || password == "" {
		Res.SendErrMessage(c, commen.UsernameOrPasswordIsNull, "username or password must not null")
		return
	}
	//查询用户是否存在   不存在则创建
	if _, exists := duser.GetByName(username); exists {
		Res.SendErrMessage(c, commen.UserAlreadyExist, "user already exists")
		return
	}
	fromPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		Res.SendErrMessage(c, commen.PasswordEncryptError, "Password encryption error")
		return
	}
	u, ok := duser.Create(username, string(fromPassword))
	if !ok {
		Res.SendErrMessage(c, commen.UserCreatError, "user create failed")
		return
	}

	ok = redis.Redis.Set(strconv.Itoa(int(u.ID)), u, 1*time.Hour)
	if !ok {
		zap.L().Error("cache user error")
	}

	token, err := utils.GetToken(u.ID)
	if err != nil {
		Res.SendErrMessage(c, commen.GetTokenError, "get token failed")
		return
	}
	c.JSON(http.StatusOK, Res.UserResponse{
		MyResponse: Res.MyResponse{
			StatusCode: commen.Success,
			StatusMsg:  "user create success",
		},
		UserId: u.ID,
		Token:  token,
	})

}
