/*
 * @Author: F1nley
 * @Date: 2021-10-01 11:17:55
 * @LastEditors: F1nley
 * @LastEditTime: 2021-10-04 21:09:22
 * @Description: 用户
 */
package models

import "time"

type User struct {
	Id         int64     `xorm:pk autoincr`
	EMail      string    `xorm:unique notnull`
	Name       string    `xorm:notnull`
	Pwd        string    `xorm: notnull`
	CreateTime time.Time `xorm: created`
	Logged     bool      `xorm: notnull`
	Profile    string
	Avator     string
}
