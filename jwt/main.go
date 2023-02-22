package main

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"time"
)

type JWT struct {
	SigningKey []byte
}

// CustomClaims Custom claims structure
type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}

type BaseClaims struct {
	ID         uint
	UUID       uuid.UUID
	Username   string
	LoginPlace string
	LoginIp    string
}

var (
	TokenExpired     = errors.New("token is expired")           // 令牌已过期
	TokenNotValidYet = errors.New("token not active yet")       // 令牌尚未激活
	TokenMalformed   = errors.New("that's not even a token")    // 那甚至不是令牌
	TokenInvalid     = errors.New("couldn't handle this token") // 无法处理此令牌
)

func NewJWT() *JWT {
	return &JWT{
		[]byte("apt2"),
	}
}

// CreateClaims 创建Claims
func (j *JWT) CreateClaims(baseClaims BaseClaims) CustomClaims {
	//bf, _ := ParseDuration(global.GVA_CONFIG.JWT.BufferTime)
	//ep, _ := ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime)
	claims := CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: int64(1 * time.Hour), // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now()),                      // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * time.Second)), // 过期时间 7天  配置文件
			Issuer:    "linktree",                                          // 签名的发行者
		},
	}
	return claims
}

// CreateToken 创建一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析 token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid
	}
}

func main() {
	j := NewJWT()

	token, err := j.CreateToken(j.CreateClaims(BaseClaims{
		ID:         0,
		UUID:       uuid.New(),
		Username:   "dsa",
		LoginPlace: "chongqing",
		LoginIp:    "182.43.54.65",
	}))
	if err != nil {
		return
	}
	fmt.Printf("%s\n", token)

	time.Sleep(1 * time.Second)

	j2 := &JWT{
		[]byte("apsst2"),
	}

	parseToken, err := j2.ParseToken(token)
	if err != nil {
		fmt.Printf("token无效")
		return
	}

	fmt.Printf("token yes %v", parseToken)

}
