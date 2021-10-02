/*
 * @Author: F1nley
 * @Date: 2021-10-01 08:36:12
 * @LastEditors: F1nley
 * @LastEditTime: 2021-10-02 09:18:32
 * @Description: 服务器
 */
package server

import (
	"Jinghong-LoveWall/server/database"
	"Jinghong-LoveWall/server/routers"
)

func Server() {
	// 数据库初始化
	database.InitDB()
	// 路由初始化
	r := routers.InitRouter()
	// 服务器开始运行
	r.Run()
}
