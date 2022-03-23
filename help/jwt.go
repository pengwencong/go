package help

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

type JWT struct {
	SigningKey []byte //自定义密钥
}

//自定义信息结构，根据需求填写
type Claims struct {
	Id                 uint   //用户id
	NickName           string //用户名
	AuthorityId        string //用户角色id
	jwt.StandardClaims
}

//定义错误信息
var (
	TokenExpired     = errors.New("Token 已经过期")
	TokenNotValidYet = errors.New("Token 未激活")
	TokenMalformed   = errors.New("Token 错误")
	TokenInvalid     = errors.New("Token 无效")
)

var claims = &jwt.StandardClaims{
ExpiresAt: 60*60*24,
Issuer:    "illnessplaza",
}

//NewJWT 初始化
func NewJWT(key string) *JWT {
	return &JWT{SigningKey: []byte(key)}
}

//CreateToken 创建 token
func (j *JWT) CreateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

//ParseToken 解析 token
func (j *JWT) ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token == nil {
		return nil, TokenInvalid
	}
	//解析到Claims 构造中
	if c, ok := token.Claims.(*Claims); ok && token.Valid {
		return c, nil
	}
	return nil, TokenInvalid
}

//RefreshToken 更新 token
//func (j *JWT) RefreshToken(tokenString string) (string, error) {
//	jwt.TimeFunc = func() time.Time {
//		return time.Unix(0, 0)
//	}
//	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
//		return j.SigningKey, nil
//	})
//	if err != nil {
//		return "", err
//	}
//	if c, ok := token.Claims.(*Claims); ok && token.Valid {
//		jwt.TimeFunc = time.Now
//		c.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
//		return j.CreateToken(*c)
//	}
//	return "", TokenInvalid
//}
