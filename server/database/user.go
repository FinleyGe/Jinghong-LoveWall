/*
 * @Author: F1nley
 * @Date: 2021-10-01 22:39:48
 * @LastEditors: F1nley
 * @LastEditTime: 2021-10-02 21:20:04
 * @Description: 用户数据库操作
 */

package database

import "Jinghong-LoveWall/server/models"

type UserState int8

const (
	UserNotFound      UserState = -1
	UserPasswordError UserState = -2
	UserCorrect       UserState = 0
)

func UserExist(email string) (bool, error) {
	return UserTable.Exist(&models.User{
		EMail: email,
	})
}

func UserRegister(email, username, password string) (int64, error) {
	user := new(models.User)
	user.EMail = email
	user.Name = username
	user.Pwd = password
	return UserTable.Insert(user)
}
