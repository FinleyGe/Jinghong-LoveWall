/*
 * @Author: F1nley
 * @Date: 2021-10-04 10:26:18
 * @LastEditors: F1nley
 * @LastEditTime: 2021-10-04 10:32:01
 * @Description: 用户登出
 */

package controllers

import "github.com/gin-gonic/gin"

func LogoutPost(c *gin.Context) {
	var token string = c.PostForm("token")

}
