package user

import (
	"BAT-douyin/commen"
	"BAT-douyin/dao/duser"
	Res "BAT-douyin/entity/res"
	"BAT-douyin/pkg/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	user, exists := duser.GetByName(username)
	if !exists {
		Res.SendErrMessage(c, commen.UserNotExist, "user not exists")
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		Res.SendErrMessage(c, commen.UserPasswordMistake, "password wrong")
		return
	}
	token, err := utils.GetToken(username, user.ID)
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
