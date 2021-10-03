/*
 * @Author: F1nley
 * @Date: 2021-10-01 10:52:56
 * @LastEditors: F1nley
 * @LastEditTime: 2021-10-03 10:33:52
 * @Description:数据库操作工具
 */

package database

import (
	"Jinghong-LoveWall/server/models"
	"log"

	"github.com/go-xorm/xorm"
)

var UserTable *xorm.Engine
var MessageTable *xorm.Engine
var TokenTable *xorm.Engine

func InitDB() {
	var err error
	UserTable, err = xorm.NewEngine("sqlite3", "./server/database/user.sqlite3")
	if err != nil {
		log.Fatalln(err)
		return
	}
	err = UserTable.Sync2(new(models.User))
	if err != nil {
		log.Fatalln(err)
		return
	}

	MessageTable, err = xorm.NewEngine("sqlite3", "./server/database/message.sqlite3")
	if err != nil {
		log.Fatalln(err)
		return
	}
	err = MessageTable.Sync2(new(models.Message))
	if err != nil {
		log.Fatalln(err)
		return
	}

	TokenTable, err = xorm.NewEngine("sqlite3", "./server/database/token.sqlite3")
	if err != nil {
		log.Fatalln(err)
		return
	}
	err = TokenTable.Sync2(new(models.Token))
	if err != nil {
		log.Fatalln(err)
		return
	}

	UserTable.ShowSQL(true)
	MessageTable.ShowSQL(true)
	TokenTable.ShowSQL(true)

}

// func pingDB() {
// 	// 每5秒检查一次连接是否正常
// 	for {
// 		err := UserTable.Ping()
// 		if err != nil {
// 			log.Fatalln(err)
// 		}
// 		time.Sleep(5 * time.Second)
// 	}
// }
