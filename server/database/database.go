/*
 * @Author: F1nley
 * @Date: 2021-10-01 10:52:56
 * @LastEditors: F1nley
 * @LastEditTime: 2021-10-02 11:15:57
 * @Description:数据库操作工具
 */

package database

import (
	"Jinghong-LoveWall/server/models"
	"log"
	"time"

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

	err = ORM.Sync2(new(models.User), new(models.Message))
	if err != nil {
		log.Fatalln(err)
		return
	}

	ORM.ShowSQL(true)
	// go pingDB()
}

func pingDB() {
	// 每5秒检查一次连接是否正常
	for {
		err := ORM.Ping()
		if err != nil {
			log.Fatalln(err)
		}
		time.Sleep(5 * time.Second)
	}
}
