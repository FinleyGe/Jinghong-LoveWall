/*
 * @Author: F1nley
 * @Date: 2021-10-01 09:39:11
 * @LastEditors: F1nley
 * @LastEditTime: 2021-10-02 13:14:36
 * @Description:
 */
package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterGet(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func RegisterPost(c *gin.Context) {
	fmt.Println("POST got")
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "AC"})
}
