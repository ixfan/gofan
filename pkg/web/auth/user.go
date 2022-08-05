package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type User struct {
	UserId   int64  `json:"userId,string"`
	UserName string `json:"userName"`
	NickName string `json:"nickName"`
	jwt.StandardClaims
}

var SecretKey = ""

var Issuer = "ixfan.cn"

// CreateTokenString 创建token
func CreateTokenString(user *User) (string, error) {
	user.StandardClaims = jwt.StandardClaims{
		ExpiresAt: time.Now().Add(3600 * 24 * 30 * time.Second).Unix(),
		Issuer:    Issuer,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, user)
	return token.SignedString([]byte(SecretKey))
}

// ParseToken 解析token
func ParseToken(token string) (*User, error) {
	user := &User{}
	tokenClaims, err := jwt.ParseWithClaims(token, user, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil || tokenClaims == nil {
		return nil, errors.New("解析token失败")
	}
	if claim, ok := tokenClaims.Claims.(*User); ok && tokenClaims.Valid {
		user.Id = claim.Id
		user.UserName = claim.UserName
		user.NickName = claim.NickName
		return user, nil
	}
	return user, nil
}
