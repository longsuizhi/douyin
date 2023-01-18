package middlewares

import (
	"douyin/api/code"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type MyClaims struct {
	UserID   uint   `json:"user_id"`
	PassWord string `json:"password"`
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 240 //设置过期时间 10天

// jwt密钥 密码自行设定
var Secret = []byte("0xDouYin")

// 生成jwt token
func GenToken(userID uint, password string) (string, error) {
	//创建一个自己的声明
	claims := &MyClaims{
		UserID:   userID,
		PassWord: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "guanliyuan01",
			Subject:   "douyin",
		},
	}
	// HS256加密 claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(Secret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	//解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, code.InvalidToken
}

// 鉴权中间件 直接拦截token 解析得到userId 鉴权并设置 user_id
func JWTMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Query("token")
		if tokenStr == "" {
			tokenStr = c.PostForm("token")
		}
		//用户不存在
		if tokenStr == "" {
			code.Response(c, code.UserNotExist, nil)
			//阻止执行
			c.Abort()
			return
		}
		//验证token
		tokenStruck, err := ParseToken(tokenStr)
		if err != nil {
			code.Response(c, code.InvalidToken, nil)
			//阻止执行
			c.Abort()
			return
		}
		//token超时
		if time.Now().Unix() > tokenStruck.ExpiresAt {
			code.Response(c, code.OverdueToken, nil)
			//阻止执行
			c.Abort()
			return
		}
		c.Set("user_id", tokenStruck.UserID)
		c.Next()
	}
}
