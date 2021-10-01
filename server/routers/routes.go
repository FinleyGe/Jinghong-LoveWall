/*
 * @Author: F1nley
 * @Date: 2021-10-01 09:34:18
 * @LastEditors: F1nley
 * @LastEditTime: 2021-10-01 10:46:14
 * @Description: 配置路由
 */

package routers

import (
	"Jinghong-LoveWall/server/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("./server/views/*")
	router.GET("/register", controllers.RegisterGet)
	return router
}
