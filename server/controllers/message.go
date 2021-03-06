/*
 * @Author: F1nley
 * @Date: 2021-10-04 11:43:49
 * @LastEditors: F1nley
 * @LastEditTime: 2021-10-05 16:18:59
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

// 新增一个信息 post方法
func NewMessagePost(c *gin.Context) {
	// 获取参数
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
	// 验证参数合法性
	if token == "" || content == "" || c.PostForm("anonymous") == "" || c.PostForm("permit_comment") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"return_value": "-1"}) // 参数错误
		return
	}

	// 通过token获取uid
	uid, e := util.DecodeToken(token)
	// 判断token合法性
	if e != nil && database.TokenValid(token) {
		c.JSON(http.StatusBadRequest, gin.H{"return_value": "-2"}) // token错误
		return
	}
	// 数据库操作
	e = database.NewMessage(uid, content, anonymous, permit_comment)
	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"return_value": "-3"}) // 服务器错误
		return
	}
	c.JSON(http.StatusOK, gin.H{"return_value": "0"})
}

// 获取一个Message

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
				"like":           message.Like,
				"comment_id":     message.CommentsID,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"return_value":   0,
				"uid":            message.Uid,
				"content":        message.Content,
				"permit_comment": message.Permit_comment,
				"create_time":    message.CreateTime,
				"update_time":    message.UpdateTime,
				"like":           message.Like,
				"comment_id":     message.CommentsID,
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
			"return_value": 0,
			"id":           message["id"],
			// "uid":            0,
			// "content":        message["content"],
			// "permit_comment": message["permit_comment"],
			// "create_time":    message["create_time"],
			// "update_time":    message["update_time"],
		})
	} else if message["anonymous"] == "1" {
		c.JSON(http.StatusOK, gin.H{
			"return_value": 0,
			"id":           message["id"],
			// "uid":            message["uid"],
			// "content":        message["content"],
			// "permit_comment": message["permit_comment"],
			// "create_time":    message["create_time"],
			// "update_time":    message["update_time"],
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"return_value": "-1",
		})
	}
}

// 信息修改
func UpdateMessage(c *gin.Context) {
	id, err := strconv.ParseInt(c.PostForm("id"), 10, 64)
	if err != nil || id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"return_value": "-1", // 参数错误
		})
		return
	}
	token := c.PostForm("token")
	if database.TokenValid(token) {
		content := c.PostForm("content")
		pc := c.PostForm("permit_comment")
		an := c.PostForm("anonymous")
		e := database.UpdateMessage(id, content, pc, an)
		if e != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"return_value": "-3",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"return_value": "0",
		})
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"return_value": "-2",
		})
		return
	}
}

func DeleteMessagePost(c *gin.Context) {
	id, err := strconv.ParseInt(c.PostForm("id"), 10, 64)
	token := c.PostForm("token")
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"return_value": "-1"})
		return
	}
	if database.TokenValid(token) {
		uid, _ := util.DecodeToken(token)
		mes, e := database.GetMessaegById(id)
		if e != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"return_value": "-3",
			})
			return
		}
		if mes.Uid == uid {
			database.DeleteMessageById(id)
			c.JSON(http.StatusOK, gin.H{
				"return_value": "0",
			})
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"return_value": "-2",
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"return_value": "-2"})
		return
	}
}

func LikeMessagePost(c *gin.Context) {
	id, e := strconv.ParseInt(c.PostForm("id"), 10, 64)
	token := c.PostForm("token")
	if id <= 0 || e != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"return_value": "-1",
		})
		return
	}

	if database.TokenValid(token) {
		uid, _ := util.DecodeToken(token)
		e = database.LikeMessage(id, uid)
		if e != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"return_value": "-2",
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"return_value": "0",
			})
			return
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"return_value": "-3",
		})
		return
	}
}
