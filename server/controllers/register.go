/*
 * @Author: F1nley
 * @Date: 2021-10-01 09:39:11
 * @LastEditors: F1nley
 * @LastEditTime: 2021-10-01 22:38:30
 * @Description:
 */
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterGet(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{"title": "注册页"})
}

func RegisterPost(c *gin.Context) {

}