package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type UserClaim struct {
	Username string
	UserId   uint
	jwt.StandardClaims
}

var MySecret = "BAT-douyin"

// GetToken 获取Token
func GetToken(username string, id uint) (string, error) {

	//设置过期时间为一天
	expiresAt := time.Now().Add(time.Hour * 24).Unix()

	//设置载荷，username和userid还有过期时间
	claims := UserClaim{}
	claims.Username = username
	claims.UserId = id
	claims.ExpiresAt = expiresAt

	//生成令牌 采用HMAC SHA256算法加密
	//令牌签名
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(MySecret))
	if err != nil {
		return "", err
	}
	return token, nil
}

// ValidateJwt 解析token，返回claims，内含登陆用户的id和name
func ValidateJwt(tokenString string) (*UserClaim, bool) {
	//解析令牌字符串
	token, err := jwt.ParseWithClaims(tokenString, &UserClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(MySecret), nil
	})
	if err != nil {
		return nil, false
	}
	//获取载荷
	claims, ok := token.Claims.(*UserClaim)
	claims.ExpiresAt = time.Now().Add(time.Hour * 1).Unix()
	if ok && token.Valid {
		return claims, true
	}
	return nil, false
}
