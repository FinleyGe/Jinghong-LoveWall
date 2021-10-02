/*
 * @Author: F1nley
 * @Date: 2021-10-01 09:39:11
 * @LastEditors: F1nley
 * @LastEditTime: 2021-10-02 21:22:40
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
	if email == "" || username == "" || password == "" {
		c.JSON(400, gin.H{"return_value": "-3"})
	}
	has, err := database.UserExist(email)
	if err != nil {
		c.JSON(500, gin.H{"return_value": "-2"})
		return
	}
	if has {
		c.JSON(304, gin.H{"return_value": "-1"})
		return
	}

	aff, err := database.UserRegister(email, username, password)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(aff)
	c.JSON(200, gin.H{"return_value": "0"})
}
