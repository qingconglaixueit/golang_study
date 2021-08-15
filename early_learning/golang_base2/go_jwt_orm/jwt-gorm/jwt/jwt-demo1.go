package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	ErrorServerBusy = "server is busy"
	ErrorReLogin    = "relogin"
)

// type StandardClaims struct {
// 	Audience  string `json:"aud,omitempty"` //该JWT所面向的用户
// 	ExpiresAt int64  `json:"exp,omitempty"` //token什么时候过期
// 	Id        string `json:"jti,omitempty"` //ID为web token提供唯一标识
// 	IssuedAt  int64  `json:"iat,omitempty"` //在什么时候签发的token
// 	Issuer    string `json:"iss,omitempty"` //该JWT的签发者
// 	NotBefore int64  `json:"nbf,omitempty"` //token在此时间之前不能被接收处理
// 	Subject   string `json:"sub,omitempty"` //
// }

type JWTClaims struct {
	jwt.StandardClaims
	UserID   int    `json:"user_id"`
	Password string `json:"password"`
	Username string `json:"username"`
}

var (
	Secret     = "123#111" //salt
	ExpireTime = 3600      //token expire time
)

// http://127.0.0.1:8080/login/darren/123456
func main() {
	r := gin.Default()
	r.GET("/login/:username/:password", login)
	r.GET("/verify/:token", verify)
	r.GET("/refresh/:token", refresh)
	r.GET("/sayHello/:token", sayHello)
	_ = r.Run(":8080")
}

//generate jwt token
func genToken(claims *JWTClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(Secret))
	if err != nil {
		return "", errors.New(ErrorServerBusy)
	}
	return signedToken, nil
}

//登录，获取jwt token
func login(c *gin.Context) {
	username := c.Param("username")
	password := c.Param("password")
	claims := &JWTClaims{
		UserID:   1,
		Username: username,
		Password: password,
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()
	singedToken, err := genToken(claims)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.String(http.StatusOK, singedToken)
}

//验证jwt token
func verifyAction(strToken string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(strToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if err != nil {
		return nil, errors.New(ErrorServerBusy)
	}
	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, errors.New(ErrorReLogin)
	}
	if err := token.Claims.Valid(); err != nil {
		return nil, errors.New(ErrorReLogin)
	}
	fmt.Println("verify")
	return claims, nil
}
func sayHello(c *gin.Context) {
	strToken := c.Param("token")
	claim, err := verifyAction(strToken)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
	}
	c.String(http.StatusOK, "hello, ", claim.Username)
}
func verify(c *gin.Context) {
	strToken := c.Param("token")
	claim, err := verifyAction(strToken)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.String(http.StatusOK, "verify: ", claim.Username)
}
func refresh(c *gin.Context) {
	strToken := c.Param("token")
	claims, err := verifyAction(strToken)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	claims.ExpiresAt = time.Now().Unix() + (claims.ExpiresAt - claims.IssuedAt)
	signedToken, err := genToken(claims)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.String(http.StatusOK, signedToken, ", ", claims.ExpiresAt)
}

//  指定加密密钥
var jwtSecret = []byte{0, 0, 0}

//Claim是一些实体（通常指的用户）的状态和额外的元数据
type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// 根据用户的用户名和密码产生token
func GenerateToken(username, password string) (string, error) {
	//设置token有效时间
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: expireTime.Unix(),
			// 指定token发行人
			Issuer: "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//该方法内部生成签名字符串，再用于获取完整、已签名的token
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// 根据传入的token值获取到Claims对象信息，（进而获取其中的用户名和密码）
func ParseToken(token string) (*Claims, error) {

	//用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		// 从tokenClaims中获取到Claims对象，并使用断言，将该对象转换为我们自己定义的Claims
		// 要传入指针，项目中结构体都是用指针传递，节省空间。
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err

}
