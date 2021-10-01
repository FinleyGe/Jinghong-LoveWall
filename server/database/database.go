/*
 * @Author: F1nley
 * @Date: 2021-10-01 10:52:56
 * @LastEditors: F1nley
 * @LastEditTime: 2021-10-01 11:14:48
 * @Description:数据库操作工具
 */

package database

import (
	"log"

	"github.com/go-xorm/xorm"
)

var ORM *xorm.Engine

func InitDB() {
	var err error
	ORM, err = xorm.NewEngine("sqlite3", "./server/database/test.sqlite3")
	if err != nil {
		log.Fatalln(err)
		return
	}
	err = ORM.Ping()
	if err != nil {
		log.Fatalln(err)
		return
	}
	ORM.ShowSQL(true)
}
