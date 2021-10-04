/*
 * @Author: F1nley
 * @Date: 2021-10-01 10:52:56
 * @LastEditors: F1nley
 * @LastEditTime: 2021-10-04 10:38:27
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
	UserTable, err = xorm.NewEngine("sqlite3", "./server/database/db/user.sqlite3")
	if err != nil {
		log.Fatalln(err)
		return
	}
	err = UserTable.Sync2(new(models.User))
	if err != nil {
		log.Fatalln(err)
		return
	}

	MessageTable, err = xorm.NewEngine("sqlite3", "./server/database/db/message.sqlite3")
	if err != nil {
		log.Fatalln(err)
		return
	}
	err = MessageTable.Sync2(new(models.Message))
	if err != nil {
		log.Fatalln(err)
		return
	}

	TokenTable, err = xorm.NewEngine("sqlite3", "./server/database/db/token.sqlite3")
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
