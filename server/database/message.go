/*
 * @Author: F1nley
 * @Date: 2021-10-04 12:18:18
 * @LastEditors: F1nley
 * @LastEditTime: 2021-10-05 16:05:24
 * @Description: message
 */

package database

import (
	"Jinghong-LoveWall/server/models"
	"fmt"
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

func GetMessaegById(id int64) (models.Message, error) {
	message := make([]models.Message, 0)
	err := MessageTable.Where("id = ?", id).Find(&message)
	return message[0], err
}

func GetMessaegByUid(uid int64) ([]models.Message, error) {
	message := make([]models.Message, 0)
	err := MessageTable.Where("uid = ?", uid).Find(&message)
	return message, err
}

func RandomMessage() ([]map[string]string, error) {
	return MessageTable.QueryString("SELECT * FROM message ORDER BY RANDOM() limit 1")
}

func UpdateMessage(id int64, content string, pc string, an string) error {
	message := new(models.Message)
	message.Id = id
	message.Content = content
	if pc == "0" {
		message.Permit_comment = false
	} else {
		message.Permit_comment = true
	}
	if an == "0" {
		message.Anonymous = false
	} else {
		message.Anonymous = true
	}
	fmt.Println(message)
	a, e := MessageTable.ID(id).Cols("content", "anonymous", "permit_comment", "update_time").Update(message)
	if a == 0 {
		log.Fatalln(a)
	}
	return e
}

func DeleteMessageById(id int64) error {
	message := new(models.Message)
	message.Id = id
	a, e := MessageTable.Delete(message)
	log.Println("Delete Message:", a)
	return e
}

func LikeMessage(id, uid int64) error {
	message := new(models.Message)
	message.Like = append(message.Like, uid)
	a, e := MessageTable.ID(id).Cols("like").Update(message)
	if a == 0 {
		log.Fatalln("like", a)
	}
	return e
}
