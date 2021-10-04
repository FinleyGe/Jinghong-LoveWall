/*
 * @Author: F1nley
 * @Date: 2021-10-04 11:44:58
 * @LastEditors: F1nley
 * @LastEditTime: 2021-10-04 11:44:58
 * @Description:
 */
package controllers

import (
	"Jinghong-LoveWall/server/database"
	"Jinghong-LoveWall/server/util"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginPost(c *gin.Context) {
	var email string = c.PostForm("email")
	var password string = c.PostForm("password")
	if email == "" || password == "" {
		// 参数错误
		c.JSON(400, gin.H{"return_value": "-6", "token": ""})
		return
	}
	exist, err := database.UserExist(email)

	if !exist { // 用户不存在
		c.JSON(403, gin.H{"return_value": "-1", "token": ""})
		return
	}
	if err != nil { // 服务器错误
		c.JSON(500, gin.H{"return_value": "-4", "token": ""})
		return
	}
	userInfo, err := database.UserInfoByEmail(email)
	if err != nil { // 服务器错误
		c.JSON(500, gin.H{"return_value": "-4", "token": ""})
	}
	if userInfo.Logged {
		c.JSON(403, gin.H{"return_value": "-5", "token": ""})
		return
	}
	if userInfo.Pwd == password {
		token, e := database.UserLogin(userInfo.Id)
		if e != nil {
			c.JSON(500, gin.H{"return_value": "-4", "token": ""})
			return
		}
		c.JSON(200, gin.H{"return_value": "0", "token": token})
	} else {
		c.JSON(403, gin.H{"return_value": "-2", "token": ""})
		return
	}
}

func LogoutPost(c *gin.Context) {
	var token string = c.PostForm("token")
	if token == "" { // 参数错误
		c.JSON(http.StatusBadRequest, gin.H{"return_value": "-1"})
		return
	}
	uid, err := util.DecodeToken(token)
	if err != nil { // token错误
		log.Fatalln(err)
		c.JSON(http.StatusInternalServerError, gin.H{"return_value": "-3"})
		return
	}
	err = database.DeleteToken(uid)
	if err != nil { // 参数错误
		c.JSON(http.StatusBadRequest, gin.H{"return_value": "-1"})
		return
	}
	err = database.UserLogout(uid)
	if err != nil { // 参数错误
		c.JSON(http.StatusBadRequest, gin.H{"return_value": "-1"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"return_value": "0"})
}

func RegisterPost(c *gin.Context) {
	var email string = c.PostForm("email")
	var username string = c.PostForm("username")
	var password string = c.PostForm("password")

	// 参数错误
	if email == "" || username == "" || password == "" {
		c.JSON(400, gin.H{"return_value": "-3"})
	}

	//参数正确
	has, err := database.UserExist(email)
	// 数据库错误
	if err != nil {
		c.JSON(500, gin.H{"return_value": "-2"})
		return
	}
	// 用户存在
	if has {
		c.JSON(304, gin.H{"return_value": "-1"})
		return
	}

	// 没有错误
	// 用户注册
	aff, err := database.UserRegister(email, username, password)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(aff)
	// 返回返回值

	c.JSON(200, gin.H{"return_value": "0"})
}
