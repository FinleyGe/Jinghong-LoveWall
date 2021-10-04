/*
 * @Author: F1nley
 * @Date: 2021-10-04 10:38:42
 * @LastEditors: F1nley
 * @LastEditTime: 2021-10-04 11:29:52
 * @Description: token
 */

package database

import (
	"Jinghong-LoveWall/server/models"
	"log"
)

func GetUidByToken(token string) (int64, error) {
	t := make([]models.Token, 1)
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
