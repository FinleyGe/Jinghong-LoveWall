/*
 * @Author: F1nley
 * @Date: 2021-10-04 10:38:42
 * @LastEditors: F1nley
 * @LastEditTime: 2021-10-05 10:20:56
 * @Description: token
 */

package database

import (
	"Jinghong-LoveWall/server/models"
	"fmt"
	"log"
)

func GetUidByToken(token string) (int64, error) {
	t := make([]models.Token, 0)
	err := TokenTable.Where("token = ?", token).Find(&t)
	return t[0].Uid, err
}

func DeleteToken(uid int64) error {
	token := new(models.Token)
	token.Uid = uid
	a, e := TokenTable.Delete(token)
	log.Println("Delete Token", a)
	return e
}

func TokenValid(token string) bool {
	t := make([]models.Token, 0)
	err := TokenTable.Where("token = ?", token).Find(&t)
	fmt.Println(t)
	if t[0].Uid != 0 && err == nil {
		return true
	}
	return false
}
