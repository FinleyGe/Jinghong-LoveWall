/*
 * @Author: F1nley
 * @Date: 2021-10-04 11:43:49
 * @LastEditors: F1nley
 * @LastEditTime: 2021-10-04 21:38:02
 * @Description:
 */

package controllers

import (
	"Jinghong-LoveWall/server/database"
	"Jinghong-LoveWall/server/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewMessagePost(c *gin.Context) {
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

func GetMessageGet(c *gin.Context) {
	id, err := strconv.ParseInt(c.DefaultQuery("id", "0"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"return_value": "-2"})
		return
	} else if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"return_value": "-1"})
	}
	message, err := database.GetMessaegById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"return_value": "-2",
		})
		return
	}
	if message.Id != 0 {
		if message.Anonymous {
			c.JSON(http.StatusOK, gin.H{
				"return_value":   0,
				"uid":            0,
				"content":        message.Content,
				"permit_comment": message.Permit_comment,
				"create_time":    message.CreateTime,
				"update_time":    message.UpdateTime,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"return_value":   0,
				"uid":            message.Uid,
				"content":        message.Content,
				"permit_comment": message.Permit_comment,
				"create_time":    message.CreateTime,
				"update_time":    message.UpdateTime,
			})
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"return_value": -3,
		})
	}

}

func RandomMessageGet(c *gin.Context) {
	m, err := database.RandomMessage()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"return_value": "-1",
		})
		return
	}
	message := m[0]
	if message["anonymous"] == "0" {
		// 是匿名的
		c.JSON(http.StatusOK, gin.H{
			"return_value":   0,
			"uid":            0,
			"content":        message["content"],
			"permit_comment": message["permit_comment"],
			"create_time":    message["create_time"],
			"update_time":    message["update_time"],
		})
	} else if message["anonymous"] == "1" {
		c.JSON(http.StatusOK, gin.H{
			"return_value":   0,
			"uid":            message["uid"],
			"content":        message["content"],
			"permit_comment": message["permit_comment"],
			"create_time":    message["create_time"],
			"update_time":    message["update_time"],
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"return_value": "-1",
		})
	}
}
