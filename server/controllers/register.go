/*
 * @Author: F1nley
 * @Date: 2021-10-01 09:39:11
 * @LastEditors: F1nley
 * @LastEditTime: 2021-10-03 10:39:33
 * @Description:
 */
package controllers

import (
	"Jinghong-LoveWall/server/database"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

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
