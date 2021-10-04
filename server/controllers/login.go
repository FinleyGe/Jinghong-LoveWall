/*
 * @Author: F1nley
 * @Date: 2021-10-03 10:34:29
 * @LastEditors: F1nley
 * @LastEditTime: 2021-10-04 10:35:17
 * @Description: 登录
 */

package controllers

import (
	"Jinghong-LoveWall/server/database"

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
