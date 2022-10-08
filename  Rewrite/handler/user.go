package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"myproject/model"
)

func User(c *gin.Context) {
	var user model.Users
	var identity string
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"message": "输入有误，格式错误"})
		return
	}
	if _, a := model.IfExistUserName(user.Name); a != 1 {
		c.JSON(200, gin.H{
			"message": "该用户名已经存在请更换其他用户名"})
		return
	}
	user_id := model.Register(user.Name, user.Password, user.Identity)
	fmt.Println(user.Name)
	switch user.Identity {
	case "0":
		identity = "最高管理员"
	case "1":
		identity = "前台人员"
	case "2":
		identity = "普通用户"
	}
	c.JSON(200, gin.H{
		"user_id":  user_id,
		"identity": identity,
	})
}

func Login(c *gin.Context) {
	var user model.Users
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"message": "输入格式错误"})
		return
	}
	fmt.Println(user.Name, user.Password)
	//验证用户名是否存在
	if model.IfExistUser(user.UserID, user.Name) {
		c.JSON(404, gin.H{"message": "用户不存在"})
		return
	}
	if model.VerifyPassword(user.Name, user.Password) {
		c.JSON(200, gin.H{
			"message": "登录成功",
			"token":   model.CreateToken(user.Name, user.Password, user.UserID, user.Identity),
		})
		token := model.CreateToken(user.Name, user.Password, user.UserID, user.Identity)
		fmt.Println(token)
		return
	}
}
