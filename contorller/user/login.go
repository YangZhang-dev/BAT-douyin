package user

import (
	"BAT-douyin/commen"
	"BAT-douyin/dao/duser"
	Res "BAT-douyin/entity/res"
	"BAT-douyin/model"
	"BAT-douyin/pkg/utils"
	"BAT-douyin/redis"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"time"
)

func Login(c *gin.Context) {

	username := c.DefaultQuery("username", "")
	password := c.DefaultQuery("password", "")
	if username == "" || password == "" {
		Res.SendErrMessage(c, commen.UsernameOrPasswordIsNull, "username or password must not null")
		return
	}

	exists := false
	user := &model.User{}
	err := json.Unmarshal([]byte(redis.Redis.Get(username)), user)
	if err != nil {
		user, exists = duser.GetByName(username)
		if !exists {
			Res.SendErrMessage(c, commen.UserNotExist, "user not exists")
			return
		}
		ok := redis.Redis.Set(strconv.Itoa(int(user.ID)), user, 1*time.Hour)
		if !ok {
			zap.L().Error("cache user error")
		}
	}

	//校验密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		Res.SendErrMessage(c, commen.UserPasswordMistake, "password wrong")
		return
	}

	token, err := utils.GetToken(user.ID)
	if err != nil {
		Res.SendErrMessage(c, commen.GetTokenError, "get token failed")
		return
	}

	c.JSON(http.StatusOK, Res.UserResponse{
		MyResponse: Res.MyResponse{
			StatusCode: commen.Success,
			StatusMsg:  "user login success",
		},
		UserId: user.ID,
		Token:  token,
	})
}
