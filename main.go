/*
 * @Author: F1nley
 * @Date: 2021-09-30 16:26:19
 * @LastEditors: F1nley
 * @LastEditTime: 2021-10-01 11:08:43
 * @Description: 精弘表白墙项目 main文件
 */

package main

import (
	"Jinghong-LoveWall/server"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	server.Server()
}
