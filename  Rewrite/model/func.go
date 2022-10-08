package model

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//实现注册
func Register(name string, password string, identity string) interface{} {
	user := Users{Name: name, Password: password, Identity: identity}
	if err := DB.Table("users").Create(&user).Error; err != nil {
		fmt.Println("registError" + err.Error())
		return " "
	}

	return user.UserID
}

//验证用户是否存在
func IfExistUser(ID string, Name string) bool {
	var user = make([]Users, 1)
	if err := DB.Table("users").Where("user_id=? AND user_name= ?", ID, Name).Find(&user).Error; err != nil {
		log.Println(err.Error())
		return true
	}
	if len(user) != 1 {
		return false
	}
	return true
}

//验证用户密码是否正确
func VerifyPassword(name string, password string) bool {
	var user Users
	if err := DB.Table("users").Where("user_name = ? user_password = ?", name, password).Find(&user).Error; err != nil {
		log.Println(err.Error())
		return false
	}
	if user.Name != name || user.Password != password {
		return false
	}
	return true
}

//获取userID
func GetID(name string, password string) string {
	var user Users
	if err := DB.Table("users").Where("user_name=？AND user_password = ?", name, password).Find(&user).Error; err != nil {
		log.Println(err.Error())
		return " "
	} else {
		return user.UserID
	}
}

//生成token与验证
type jwtClaims struct {
	jwt.StandardClaims
	UserID   string `json:"usr_id"`
	UserName string `json:"user_name"`
	Password string `json:"user_password"`
	Identity string `json:"identity"`
}

var (
	key        = "football" //salt
	ExpireTime = 604800     //token expire time
)

func CreateToken(name string, password string, id string, identity string) string {
	//生成Token
	claims := &jwtClaims{
		UserID:   id,
		UserName: name,
		Password: password,
		Identity: identity,
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()
	singedToken, err := genToken(*claims)
	if err != nil {
		log.Print("produceToken err:")
		fmt.Println(err)
		return ""
	}
	return singedToken
}

func genToken(claims jwtClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	signedToken, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

//var key = "sault" //加盐
//验证token
func VerifyToken(token string) (string, error) {
	TempToken, err := jwt.ParseWithClaims(token, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return "", errors.New("Token解析失败")
	}
	claims, ok := TempToken.Claims.(*jwtClaims)
	if !ok {
		return "", errors.New("发生错误")
	}
	if err := TempToken.Claims.Valid(); err != nil {
		return "", errors.New("发生错误")
	}
	return claims.UserID, nil
}

//获取用户信息
func GetUserName(UserID string) (string, string) {
	var user Users
	var Identity string
	if err := DB.Table("users").Where("user_id=?", UserID).Find(&user).Error; err != nil {
		return " ", " "
	} else {
		if user.Identity == "0" {
			Identity = "最高管理员"
		}
		if user.Identity == "1" {
			Identity = "前台人员"
		}
		if user.Identity == "2" {
			Identity = "普通用户"
		}
	}
	return user.Name, Identity
}

func GetUserPicture(UserId string) string {
	var user Users
	if err := DB.Table("users").Where("user_id", UserId).Find(&user).Error; err != nil {
		return " "
	} else {
		return user.UserPicture
	}
}

//防止出现重复用户名的用户
func IfExistUserName(UserName string) (error, int) {
	var temp Users
	if err := DB.Table("users").Where("user_name=?", UserName).Find(&temp).Error; err != nil {
		log.Println(err)
		return err, 1
	}
	return nil, 0
}
