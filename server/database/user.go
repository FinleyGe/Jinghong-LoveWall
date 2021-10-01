/*
 * @Author: F1nley
 * @Date: 2021-10-01 22:39:48
 * @LastEditors: F1nley
 * @LastEditTime: 2021-10-01 23:12:40
 * @Description: 用户数据库操作
 */

package database

type UserState int8

const (
	UserNotFound      UserState = -1
	UserPasswordError UserState = -2
	UserCorrect       UserState = 0
)

// func CheckUser(user models.User) UserState {
// 	has, err := ORM.Where("e_mail = ?", user.EMail).Desc("id").Get()
// }
