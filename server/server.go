/*
 * @Author: F1nley
 * @Date: 2021-10-01 08:36:12
 * @LastEditors: F1nley
 * @LastEditTime: 2021-10-01 09:44:58
 * @Description: 服务器
 */
package server

import (
	"Jinghong-LoveWall/server/routers"
)

func Server() {
	r := routers.InitRouter()
	r.Run()

}
