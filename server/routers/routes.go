/*
 * @Author: F1nley
 * @Date: 2021-10-01 09:34:18
 * @LastEditors: F1nley
 * @LastEditTime: 2021-10-04 11:00:29
 * @Description: 配置路由
 */

package routers

import (
	"Jinghong-LoveWall/server/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	// router.LoadHTMLGlob("./server/views/*")
	// router.GET("/register", controllers.RegisterGet)
	router.POST("/api/register", controllers.RegisterPost)
	router.POST("/api/login", controllers.LoginPost)
	router.POST("api/logout", controllers.LogoutPost)
	return router
}
