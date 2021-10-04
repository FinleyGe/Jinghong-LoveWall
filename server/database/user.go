/*
 * @Author: F1nley
 * @Date: 2021-10-01 22:39:48
 * @LastEditors: F1nley
 * @LastEditTime: 2021-10-04 10:55:08
 * @Description: 用户数据库操作
 */

package database

import (
	"Jinghong-LoveWall/server/models"
	"Jinghong-LoveWall/server/util"
	"log"
)

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

// 用户注册
func UserRegister(email, username, password string) (int64, error) {
	user := new(models.User)
	user.EMail = email
	user.Name = username
	user.Pwd = password
	user.Logged = false
	return UserTable.Insert(user)
}

// 查找用户 通过email
func UserInfoByEmail(email string) (models.User, error) {
	userInfo := make([]models.User, 0)
	err := UserTable.Where("e_mail = ?", email).Find(&userInfo)
	return userInfo[0], err
}

// 用户登录
func UserLogin(id int64) (string, error) {
	user := new(models.User)
	user.Id = id
	user.Logged = true

	UserTable.ID(id).Cols("logged").Update(user)

	t, err := util.GetToken(id)
	var token models.Token
	token.Uid = id
	token.Token = t
	TokenTable.Insert(token)
	return t, err
}

func UserLogout(uid int64) error {
	user := new(models.User)
	user.Id = uid
	user.Logged = false

	a, e := UserTable.ID(uid).Cols("logged").Update(user)
	log.Println("UserTable update", a)
	return e
}
