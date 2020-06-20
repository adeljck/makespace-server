package middleware

import (
	"bytes"
	"crypto/rand"
	"errors"
	"makespace-remaster/serializer"
	"makespace-remaster/service"
	"math/big"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWTAuth 中间件，检查token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusOK, serializer.PureErrorResponse{
				Status: -1,
				Msg:    "请求未携带token，无权限访问",
			})
			c.Abort()
			return
		}
		j := NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if claims == nil || err != nil {
			c.JSON(http.StatusOK, serializer.PureErrorResponse{
				Status: -2,
				Msg:    "不合法的token",
			})
			c.Abort()
			return
		}
		if err != nil {
			if err == TokenExpired {
				c.JSON(http.StatusOK, serializer.PureErrorResponse{
					Status: -3,
					Msg:    "Token过期",
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusOK, serializer.PureErrorResponse{
				Status: -1,
				Msg:    err.Error(),
			})
			c.Abort()
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("claims", claims)
	}
}

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

// 一些常量
var (
	TokenExpired     error  = errors.New("Token is expired")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("That's not even a token")
	TokenInvalid     error  = errors.New("Couldn't handle this token:")
	signKey          string = os.Getenv("SESSIONSECRET")
)

// 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	Id string `json:"id"`
	Role int `json:role`
	jwt.StandardClaims
}

// 新建一个jwt实例
func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

// 获取signKey
func GetSignKey() string {
	return signKey
}

// CreateToken 生成一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 解析Tokne
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	var result string
	for i, _ := range tokenString {
		if i > 5 && i < len(tokenString)-6 {
			result = result + string(tokenString[i])
		}
	}
	token, err := jwt.ParseWithClaims(result, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
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
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

// 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}

// 生成令牌
func GenerateToken(c *gin.Context, info *service.Info) {
	j := &JWT{
		[]byte(signKey),
	}
	//if c.FullPath() == "/api/v1/forget"{
	//	j.SigningKey = []byte(os.Getenv("RESET_SESSIONSECRET"))
	//}
	claims := CustomClaims{
		info.ID.Hex(),
		info.Role,
		jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 1800), // 过期时间 一小时
			Issuer:    os.Getenv("SESSIONNAME"),        //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)
	token = CreateRandomString(6) + token + CreateRandomString(6)
	if err != nil {
		c.JSON(http.StatusOK, serializer.Response{
			Status: -1,
			Msg:    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,serializer.Response{
		Status: 200,
		Data:   gin.H{"_id":info.ID.Hex(),"token":token,"name":info.Name,"role":info.Role,"status":info.Status},
		Msg:    "success",
	})
	c.Header("Authorization","Bearer "+token)
	return

}
func CreateRandomString(len int) string {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}
