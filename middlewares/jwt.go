package middlewares

import (
	"BAT-douyin/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const (
	SUCCESS              = 10001
	ERROR_TOKEN_VALIDATE = 10002
	ERROR_TOKEN_EXPIRED  = 10003
)

// Jwt JWT身份验证中间件，某些功能必须登陆才能使用就加上这个中间件
func Jwt() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		code := SUCCESS
		//通过查询字符串获取Token，如果为空则从post数据中查询
		token := ctx.Query("token")
		if token == "" {
			token = ctx.PostForm("token")
		}

		//解析Token
		claim, ok := utils.ValidateJwt(token)
		if !ok {
			code = ERROR_TOKEN_VALIDATE
			ctx.JSON(http.StatusOK, gin.H{"code": code, "message": "token is wrong"})
			ctx.Abort()
			return
		}

		//过期判断
		if time.Now().Unix() > claim.ExpiresAt {
			code = ERROR_TOKEN_EXPIRED
			ctx.JSON(http.StatusOK, gin.H{"code": code, "message": "token is out of time"})
			ctx.Abort()
			return
		}

		//设置用户信息
		ctx.Set("claim", claim)
		ctx.Next()
	}
}
