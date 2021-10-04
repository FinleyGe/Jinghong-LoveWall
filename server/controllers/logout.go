/*
 * @Author: F1nley
 * @Date: 2021-10-04 10:26:18
 * @LastEditors: F1nley
 * @LastEditTime: 2021-10-04 11:01:52
 * @Description: 用户登出
 */

package controllers

import (
	"Jinghong-LoveWall/server/database"
	"Jinghong-LoveWall/server/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
