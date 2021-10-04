/*
 * @Author: F1nley
 * @Date: 2021-10-04 12:18:18
 * @LastEditors: F1nley
 * @LastEditTime: 2021-10-04 12:26:42
 * @Description: message
 */

package database

import (
	"Jinghong-LoveWall/server/models"
	"log"
)

func NewMessage(uid int64, content string, an bool, pc bool) error {
	message := new(models.Message)
	message.Uid = uid
	message.Content = content
	message.Anonymous = an
	message.Permit_comment = pc

	a, e := MessageTable.Insert(message)

	log.Println(a)
	return e
}
