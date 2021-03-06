/*
 * @Author: F1nley
 * @Date: 2021-10-01 09:34:18
 * @LastEditors: F1nley
 * @LastEditTime: 2021-10-05 15:55:49
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
	router.POST("/api/logout", controllers.LogoutPost)
	router.POST("/api/new_message", controllers.NewMessagePost)
	router.GET("/api/message", controllers.GetMessageGet)
	router.GET("/api/message/random", controllers.RandomMessageGet)
	router.POST("/api/message/edit", controllers.UpdateMessage)
	router.POST("/api/message/delete", controllers.DeleteMessagePost)
	router.POST("/api/message/like", controllers.LikeMessagePost)
	return router
}
