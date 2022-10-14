package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"myblog/config"
	"myblog/util/r"
	"strings"
	"time"
)

var jwtKey = []byte(config.GinConf.JwtKey)

type MyCustomClaims struct {
	UserId int    `json:"userId"`
	Role   string `json:"role"`
	Uuid   string `json:"uuid"`
	jwt.RegisteredClaims
}

func CreateToken(userId int, role string, uuid string) (string, int) {
	//var a int
	//a = 5
	claims := MyCustomClaims{
		UserId: userId,
		Role:   role,
		Uuid:   uuid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(config.GinConf.ExpireTime) * time.Hour)),
			Issuer:    "spxzx",
		},
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // token
	token, err := reqClaim.SignedString(jwtKey)
	if err != nil {
		return "", r.TokenCreateFail
	}
	return token, r.SUCCESS
}

func CheckToken(tokenStr string) (*MyCustomClaims, int) {
	token, _ := jwt.ParseWithClaims(tokenStr, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, r.SUCCESS
	}
	return nil, r.TokenWrong
}

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			r.Send(c, r.NoLogin)
			c.Abort()
			return
		}
		tokenSlice := strings.SplitN(token, " ", 2)
		if len(tokenSlice) != 2 && tokenSlice[0] != "Bearer" {
			r.Send(c, r.TokenFormatterError)
			c.Abort()
			return
		}
		tokenStruck, code := CheckToken(tokenSlice[1])
		if code != r.SUCCESS {
			r.Send(c, code)
			c.Abort()
			return
		}
		if time.Now().Unix() > tokenStruck.ExpiresAt.Unix() {
			r.Send(c, r.TokenRuntime)
			c.Abort()
			return
		}
		c.Set("userInfoId", tokenStruck.UserId)
		c.Set("role", tokenStruck.Role)
		c.Set("uuid", tokenStruck.Uuid)
		c.Next()
	}
}
