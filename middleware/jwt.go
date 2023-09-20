package middleware

import (
	"ginBlog/utils"
	"ginBlog/utils/errmsg"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

var JwtKey = []byte(utils.JwtKey)

type MyClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

var code int

// ReleaseToken 生成token
func ReleaseToken(username, password string) (string, int) {
	expireTime := time.Now().Add(10 * time.Hour)
	SetClaims := MyClaims{
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "ginBlog",
		},
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	token, err := reqClaim.SignedString(JwtKey)
	if err != nil {
		return "", errmsg.ERROR
	}

	return token, errmsg.SUCCESS
}

// CheckToken 验证token
func CheckToken(token string) (*MyClaims, int) {
	settoken, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if key, code := settoken.Claims.(*MyClaims); code && settoken.Valid {
		return key, errmsg.SUCCESS
	} else {
		return nil, errmsg.ERROR
	}
}

// JwtToken jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization")

		if tokenHeader == "" {
			code = errmsg.ERROR_TOKEN_EXIST
			c.Abort()
		}
		checkToken := strings.SplitN(tokenHeader, " ", 2)
		if len(checkToken) != 2 || checkToken[0] != "Bearer" {
			code = errmsg.ERROR_TOKEN_TYPE
			c.Abort()
		}
		key, Tcode := CheckToken(checkToken[1])
		if Tcode == errmsg.ERROR {
			code = errmsg.ERROR_TOKEN_WRONG
			c.Abort()
		}
		if time.Now().Unix() > key.ExpiresAt {
			code = errmsg.ERROR_TOKEN_RUNTIME
			c.Abort()
		}
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    errmsg.GetErrMsg(code),
		})
		c.Set("username", key.Username)
		c.Next()
	}
}
