/*
 * @Author: F1nley
 * @Date: 2021-10-04 11:43:49
 * @LastEditors: F1nley
 * @LastEditTime: 2021-10-04 14:33:16
 * @Description:
 */

package controllers

import (
	"Jinghong-LoveWall/server/database"
	"Jinghong-LoveWall/server/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewMessage(c *gin.Context) {
	var token, content string
	var anonymous, permit_comment bool
	token = c.PostForm("token")
	content = c.PostForm("content")

	if c.PostForm("anonymous") == "0" {
		anonymous = false
	} else {
		anonymous = true
	}
	if c.PostForm("permit_comment") == "0" {
		permit_comment = false
	} else {
		permit_comment = true
	}

	if token == "" || content == "" || c.PostForm("anonymous") == "" || c.PostForm("permit_comment") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"return_value": "-1"}) // 参数错误
		return
	}
	uid, e := util.DecodeToken(token)
	if e != nil {
		// fmt.Println(e)
		c.JSON(http.StatusBadRequest, gin.H{"return_value": "-2"}) // token错误
		return
	}
	e = database.NewMessage(uid, content, anonymous, permit_comment)
	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"return_value": "-3"}) // 服务器错误
		return
	}
	c.JSON(http.StatusOK, gin.H{"return_value": "0"})
}
