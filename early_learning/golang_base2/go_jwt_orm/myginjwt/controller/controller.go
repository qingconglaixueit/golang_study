package controller

import (
	"errors"
	"fmt"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"time"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
//登录请求信息
type ReqInfo struct {
	Name   string `json:"name"`
	Passwd string `json:"passwd"`
}
// 构造用户表
type MyInfo struct {
	Id        int32  `gorm:"AUTO_INCREMENT"`
	Name      string `json:"name"`
	Passwd    string `json:"passwd"`
	CreatedAt *time.Time
	UpdateTAt *time.Time
}
//Myclaims定义载荷
type Myclaims struct {
	Name string `json:"userName"`
	// StandardClaims结构体实现了Claims接口(Valid()函数)
	jwtgo.StandardClaims
}
//密钥
type JWT struct {
	SigningKey []byte
}
var (
	DB               *gorm.DB
	secret                 = "iamsecret"
	TokenExpired     error = errors.New("Token is expired")
	TokenNotValidYet error = errors.New("Token not active yet")
	TokenMalformed   error = errors.New("That's not even a token")
	TokenInvalid     error = errors.New("Couldn't handle this token:")
)
//hello 接口
func Hello(c *gin.Context) {
	claims, _ := c.MustGet("claims").(*Myclaims)
	if claims != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":    "Hello wrold",
			"data":   claims,
		})
	}
}
//数据库连接
func InitMySQLCon() (err error) {
	// 可以在api包里设置成init函数
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "123456", "127.0.0.1", 3306, "mygorm")
	fmt.Println(connStr)
	DB, err = gorm.Open("mysql", connStr)

	if err != nil {
		return err
	}

	return DB.DB().Ping()
}
//初始化gorm对象映射
func InitModel() {
	DB.AutoMigrate(&MyInfo{})
}
func NewJWT() *JWT {
	return &JWT{
		[]byte(secret),
	}
}
// 登陆结果
type LoginResult struct {
	Token string `json:"token"`
	Name string `json:"name"`
}
// 创建Token(基于用户的基本信息claims)
// 使用HS256算法进行token生成
// 使用用户基本信息claims以及签名key(signkey)生成token
func (j *JWT) CreateToken(claims Myclaims) (string, error) {
	// 返回一个token的结构体指针
	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}
//生成token
func generateToken(c *gin.Context, info ReqInfo) {
	// 构造SignKey: 签名和解签名需要使用一个值
	j := NewJWT()

	// 构造用户claims信息(负荷)
	claims := Myclaims{
		info.Name,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 签名过期时间
			Issuer:    "pangsir",                       // 签名颁发者
		},
	}

	// 根据claims生成token对象
	token, err := j.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
			"data":   nil,
		})
	}

	log.Println(token)
	// 返回用户相关数据
	data := LoginResult{
		Name:  info.Name,
		Token: token,
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "登陆成功",
		"data":   data,
	})

	return
}
//解析token
func (j *JWT) ParserToken(tokenstr string) (*Myclaims, error) {
	// 输入token
	// 输出自定义函数来解析token字符串为jwt的Token结构体指针
	// Keyfunc是匿名函数类型: type Keyfunc func(*Token) (interface{}, error)
	// func ParseWithClaims(tokenString string, claims Claims, keyFunc Keyfunc) (*Token, error) {}
	token, err := jwtgo.ParseWithClaims(tokenstr, &Myclaims{}, func(token *jwtgo.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	fmt.Println(token, err)
	if err != nil {
		// jwt.ValidationError 是一个无效token的错误结构
		if ve, ok := err.(*jwtgo.ValidationError); ok {
			// ValidationErrorMalformed是一个uint常量，表示token不可用
			if ve.Errors&jwtgo.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
				// ValidationErrorExpired表示Token过期
			} else if ve.Errors&jwtgo.ValidationErrorExpired != 0 {
				return nil, TokenExpired
				// ValidationErrorNotValidYet表示无效token
			} else if ve.Errors&jwtgo.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}

		}
	}

	// 将token中的claims信息解析出来和用户原始数据进行校验
	// 做以下类型断言，将token.Claims转换成具体用户自定义的Claims结构体
	if claims, ok := token.Claims.(*Myclaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("token NotValid")
}
//登录
func Login(c *gin.Context) {
	var reqinfo ReqInfo
	var userInfo MyInfo

	err := c.BindJSON(&reqinfo)

	if err == nil {
		fmt.Println(reqinfo)

		if reqinfo.Name == "" || reqinfo.Passwd == ""{
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "账号密码不能为空",
				"data":   nil,
			})
			c.Abort()
			return
		}
		//校验数据库中是否有该用户
		err := DB.Where("name = ?", reqinfo.Name).Find(&userInfo)
		if err != nil {
			fmt.Println("数据库中没有该用户 ，可以进行添加用户数据")
			//添加用户到数据库中
			info := MyInfo{
				Name:   reqinfo.Name,
				Passwd: reqinfo.Passwd,
			}
			dberr := DB.Model(&MyInfo{}).Create(&info).Error
			if dberr != nil {
				c.JSON(http.StatusOK, gin.H{
					"status": -1,
					"msg":    "登录失败，数据库操作错误",
					"data":   nil,
				})
				c.Abort()
				return
			}
		}else{
			if userInfo.Name != reqinfo.Name || userInfo.Passwd != reqinfo.Passwd{
				c.JSON(http.StatusOK, gin.H{
					"status": -1,
					"msg":    "账号密码错误",
					"data":   nil,
				})
				c.Abort()
				return
			}
		}
		//创建token
		generateToken(c, reqinfo)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "登录失败，数据请求错误",
			"data":   nil,
		})
	}
}
